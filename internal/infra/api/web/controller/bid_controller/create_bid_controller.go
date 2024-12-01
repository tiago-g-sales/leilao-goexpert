package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/opentelemetry"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/validation"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/bid_usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)
const (		
	REQUESTNAMEOTEL = "BidController"
)

type BidController struct {
	bidUseCase  bid_usecase.BidUseCaseInterface
	TemplateData *opentelemetry.TemplateData
}


func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface, templateData *opentelemetry.TemplateData) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
		TemplateData: templateData,
	}
}


func (u *BidController)  CreateBid(c *gin.Context) {
	
	carrier := propagation.HeaderCarrier(c.Request.Header)
	ctx := c.Request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " SPAN_INICIAL")


	ctx, span := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Initial request CreateBid ")
	defer span.End()
	
	
	var bidInputDTO model.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	err := u.bidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	ctx, spanEnd := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Finish request CreateBid")
	defer spanEnd.End()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Request.Header))
	spanInicial.End()
	c.Status(http.StatusCreated)
	
}