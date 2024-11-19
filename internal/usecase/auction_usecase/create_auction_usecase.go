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




func (au *AuctionUseCase) CreateAuction(ctx context.Context, auction *model.AuctionInputDTO) (*internal_error.InternalError) {

	err := au.auctionRepositoryInterface.CreateAuction(ctx, auction)
	if err != nil {
		return err
	}

	return nil
}


