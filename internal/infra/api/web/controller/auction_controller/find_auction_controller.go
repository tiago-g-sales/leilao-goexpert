package auction_controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (au *AuctionController) FindBidByAuctionId(c *gin.Context) {

	carrier := propagation.HeaderCarrier(c.Request.Header)
	ctx := c.Request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := au.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " SPAN_INICIAL ")


	ctx, span := au.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Initial request FindBidByAuctionId")
	defer span.End()



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

	ctx, spanEnd := au.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Finish request FindBidByAuctionId")
	defer spanEnd.End()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Request.Header))
	spanInicial.End()
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
