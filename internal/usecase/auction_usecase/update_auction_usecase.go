package auction_usecase

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)



func (au *AuctionUseCase) UpdateAuctionById(ctx context.Context, auction model.AuctionInputDTO) *internal_error.InternalError {

	err := au.auctionRepositoryInterface.UpdateAuctionById(ctx, auction)
	if err != nil {
		return err
	}

	return nil

}
