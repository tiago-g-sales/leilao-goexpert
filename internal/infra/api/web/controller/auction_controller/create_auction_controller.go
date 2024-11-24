package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/api/web/validation"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/auction_usecase"
)


type AuctionController struct {
	auctionUseCase  auction_usecase.AuctionUseCaseInterface
}


func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		auctionUseCase: auctionUseCase,
	}
}


func (u *AuctionController)  CreateAuction(c *gin.Context) {
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
	c.Status(http.StatusCreated)
	
}