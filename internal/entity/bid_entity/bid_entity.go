package bid_entity

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)


type BidEntityMongo struct {
	Id 			string 		`bson:"_id"`
	UserID 		string		`bson:"user_id"`
	AuctionID 	string		`bson:"auction_id"`
	Amount 		float64		`bson:"amount"`
	Timestamp   int64   	`bson:"timestamp"`
}






type BidRepositoryInterface interface {
	CreateBid( ctx context.Context, bidList []model.BidInputDTO)(*internal_error.InternalError) 
	FindBidByAuctionId(ctx context.Context, auctionId string)([]model.BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string)(*model.BidOutputDTO, *internal_error.InternalError)

} 