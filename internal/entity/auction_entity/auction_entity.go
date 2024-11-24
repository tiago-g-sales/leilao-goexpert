package auction_entity

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)

type AuctionEntityMongo struct {
	Id 	 			string 					`bson:"_id"`
	ProductName 	string					`bson:"product_name"`
	Category 		string					`bson:"categoria"`
	Description 	string		 			`bson:"description"`
	Condition 		model.ProductCondition	`bson:"condition"`
	Status 			model.AuctionStatus  	`bson:"status"`
	Timestamp 		int64	 				`bson:"timestamp"`

}


type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction model.AuctionInputDTO) (*internal_error.InternalError)
	FindAuctions( ctx context.Context, status model.AuctionStatus, category, productName string ) ([]model.AuctionOutputDTO, *internal_error.InternalError)
	FindAuctionById(ctx context.Context,  auctionId string) (*model.AuctionOutputDTO, *internal_error.InternalError)
}