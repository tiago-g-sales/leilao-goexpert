package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)


func (u *BidController) FindBidByAuctionId(c *gin.Context) {

	carrier := propagation.HeaderCarrier(c.Request.Header)
	ctx := c.Request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " SPAN_INICIAL")


	ctx, span := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Initial request FindBidByAuctionId")
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

	bidOutputList, err :=  u.bidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ctx, spanEnd := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Finish request FindBidByAuctionId")
	defer spanEnd.End()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Request.Header))
	spanInicial.End()

	c.JSON(http.StatusOK, bidOutputList)
}