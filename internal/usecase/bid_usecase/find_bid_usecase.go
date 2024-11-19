package bid_usecase

import (
	"context"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)



func (bu *BidUseCase) FindBidByAuctionId(ctx  context.Context, auctionId string) ([]model.BidOutputDTO, *internal_error.InternalError) {

	bidList, err := bu.BidRepositoryInterface.FindBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}
	
	return bidList, nil

}

func (bu *BidUseCase) FindWinningBidByAuctionId(ctx  context.Context, auctionId string)(*model.BidOutputDTO, *internal_error.InternalError) {
	
	bidOutputDTO, err := bu.BidRepositoryInterface.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}
	
	return bidOutputDTO, nil
}