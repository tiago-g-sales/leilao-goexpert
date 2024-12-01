package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/opentelemetry"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/validation"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/auction_usecase"
)

const (		
	REQUESTNAMEOTEL = "AuctionController"
)


type AuctionController struct {
	auctionUseCase  auction_usecase.AuctionUseCaseInterface
	TemplateData *opentelemetry.TemplateData
}


func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface, templateData *opentelemetry.TemplateData) *AuctionController {
	return &AuctionController{
		auctionUseCase: auctionUseCase,
		TemplateData: templateData,
	}
}


func (u *AuctionController)  CreateAuction(c *gin.Context) {
	
	carrier := propagation.HeaderCarrier(c.Request.Header)
	ctx := c.Request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " SPAN_INICIAL")
	spanInicial.End()

	ctx, span := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Initial request CreateAuction" )
	defer span.End()
	
	
	var auctionInputDTO model.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	err := u.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ctx, spanEnd := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Finish request CreateAuction")
	defer spanEnd.End()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Request.Header))

	c.Status(http.StatusCreated)
	
}