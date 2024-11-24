package auction_usecase

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/auction_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/bid_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)

type AuctionUseCase struct {

	auctionRepositoryInterface   auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface       bid_entity.BidRepositoryInterface

}

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auction model.AuctionInputDTO) (*internal_error.InternalError) 
	FindAuctionById(ctx context.Context, auctionId string) (*model.AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context,	status model.AuctionStatus,	category, productName string) ([]model.AuctionOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string)(*model.WinningInfoOutputDTO, *internal_error.InternalError)
}


func NewAuctionUseCase(auctionRepositoryInterface auction_entity.AuctionRepositoryInterface, bidRepositoryInterface bid_entity.BidRepositoryInterface) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepositoryInterface: auctionRepositoryInterface,
		bidRepositoryInterface:     bidRepositoryInterface,
	}
}


func (au *AuctionUseCase) CreateAuction(ctx context.Context, auction model.AuctionInputDTO) (*internal_error.InternalError) {

	err := au.auctionRepositoryInterface.CreateAuction(ctx, auction)
	if err != nil {
		return err
	}

	return nil
}


