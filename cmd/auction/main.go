package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/database/mongodb"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/opentelemetry"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/service"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/controller/auction_controller"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/controller/bid_controller"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/controller/user_controller"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/database/auction"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/database/bid"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/database/user"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/auction_usecase"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/bid_usecase"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)


type TemplateData struct {
	Title              string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}



func init() {
	viper.AutomaticEnv()
}

func main() {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)


	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	otel_active := viper.GetString("OTEL")

	if otel_active != "false" {
		shutdown, err := opentelemetry.InitProvider(viper.GetString("OTEL_SERVICE_NAME"), viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"))
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
		}()

		if err != nil {
			log.Fatal(err)
		}
	}

	tracer := otel.Tracer("microservice-tracer")

	templateData := &TemplateData{
		Title:              viper.GetString("TITLE"),
		ResponseTime:       time.Duration(viper.GetInt("RESPONSE_TIME")),
		ExternalCallURL:    viper.GetString("EXTERNAL_CALL_URL"),
		ExternalCallMethod: viper.GetString("EXTERNAL_CALL_METHOD"),
		RequestNameOTEL:    viper.GetString("REQUEST_NAME_OTEL"),
		OTELTracer:         tracer,
	}

	fmt.Println(templateData)

	databaseConnection , err :=  mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionsController, auctionService := initDependencies(databaseConnection)

	go auctionService.UpdateAuctionsToEnd()

	router.GET("/auctions", auctionsController.FindAuctions) 
	router.GET("/auctions/:auctionId", auctionsController.FindBidByAuctionId) 
	router.POST("/auctions", auctionsController.CreateAuction) 
	router.GET("/auctions/winner/:auctionId", auctionsController.FindWinningBidByAuctionId) 
	router.POST("/bid", bidController.CreateBid) 
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId) 
	router.GET("/user/:userId", userController.FindUserById) 

	router.Run(viper.GetString("HTTP_PORT"))


}

func initDependencies(database *mongo.Database) (
	
	userConstroller *user_controller.UserController,
	bidConstroller *bid_controller.BidController,
	auctionConstroller *auction_controller.AuctionController,
	auctionService *service.AuctionService ) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userConstroller = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionConstroller = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidConstroller = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))
	auctionService = service.NewAuctionService(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	
	return
}

