package auction_usecase

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, auctionId string) (*model.AuctionOutputDTO, *internal_error.InternalError) {

	auctionEntity, err := au.auctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	return &model.AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp,
	}, nil

}

func (au *AuctionUseCase) FindAuctions(ctx context.Context,	status model.AuctionStatus,	category, productName string) ([]model.AuctionOutputDTO, *internal_error.InternalError) {

	auctionOutputDTOs , err := au.auctionRepositoryInterface.FindAuctions(ctx, status, category, productName)
	if err != nil {
		return nil, err
	}

	return auctionOutputDTOs, nil
}


func (au *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string)(*model.WinningInfoOutputDTO, *internal_error.InternalError){

	auction, err := au.auctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	bidWinning, err := au.bidRepositoryInterface.FindWinningBidByAuctionId(ctx, auction.Id)
	if err != nil {
		return &model.WinningInfoOutputDTO{
			Auction: *auction,
			Bid: nil,
		}, nil
	} else {
		return &model.WinningInfoOutputDTO{
			Auction: *auction,
			Bid: bidWinning,
		}, nil
		}
	}



