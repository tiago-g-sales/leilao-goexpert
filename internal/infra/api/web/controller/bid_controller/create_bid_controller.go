package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/validation"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/bid_usecase"
)

type BidController struct {
	bidUseCase  bid_usecase.BidUseCaseInterface
}


func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
	}
}


func (u *BidController)  CreateBid(c *gin.Context) {
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
	c.Status(http.StatusCreated)
	
}