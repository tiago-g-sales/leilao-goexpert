package bid_usecase

import (
	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/bid_entity"

)


type BidUseCase struct {

	BidRepositoryInterface  bid_entity.BidRepositoryInterface 
}