package auction_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)

func (au *AuctionController) FindBidByAuctionId(c *gin.Context) {

		auctionId := c.Param("auctionId")
	if err :=  uuid.Validate(auctionId); err != nil {
		 errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field: "auctionId",	
			Message: "Invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err :=  au.auctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

func (au *AuctionController) FindAuctions(c *gin.Context){
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("Error trying to validade auction status param")
		c.JSON(errRest.Code, errRest)
		return
		
	}

	auctions, err :=  au.auctionUseCase.FindAuctions(context.Background(), model.AuctionStatus(statusNumber), category, productName)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	if len(auctions) == 0 {
		errRest := rest_err.NewNotFoundError("Auctions not found") 
		c.JSON(errRest.Code, errRest)
		return	
	}

	c.JSON(http.StatusOK, auctions)

}


func (au *AuctionController) FindWinningBidByAuctionId(c *gin.Context){

	auctionId := c.Param("auctionId")
	if err :=  uuid.Validate(auctionId); err != nil {
		 errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field: "auctionId",	
			Message: "Invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err :=  au.auctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	c.JSON(http.StatusOK, auctionData)

}

func (au *AuctionController) FindAuctionsEnd(){
	status := 0
	auctionEndTime := viper.GetString("AUCTION_END_TIME")
	duration, errConv := time.ParseDuration(auctionEndTime)
	if errConv != nil {
		return 
	}

	for {

		timeNow :=  time.Now()
		auctions, err :=  au.auctionUseCase.FindAuctions(context.Background(), model.AuctionStatus(status), "", "")
		if err != nil {
			return
		}

		for _, auction := range auctions {
			diff := timeNow.Sub(auction.Timestamp)
			
			if   diff >= duration{

				auctionIn := model.AuctionInputDTO{		
					Id: auction.Id,
					Status: model.Completed,
				}
				au.auctionUseCase.UpdateAuctionById(context.Background(), auctionIn )				
				fmt.Printf("Auction timestamp is after current time auction.Timestamp: %s, timeNow: %s", auction.Timestamp, timeNow)

			}else {
				fmt.Printf("AVALIACAO DE TIMESTAMP auction.Timestamp: %s, timeNow: %s", auction.Timestamp, timeNow)
			}
			

		}
		fmt.Println("Todos os Leiloes foram encerrados")
		time.Sleep(10 * time.Second)
	}
}