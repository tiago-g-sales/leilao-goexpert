package auction

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/auction_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)


type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auction model.AuctionInputDTO) (*internal_error.InternalError) {
	auctionEntityMongo := &auction_entity.AuctionEntityMongo{ 
		Id: auction.Id,
		ProductName: auction.ProductName,
		Category: auction.Category,
		Description: auction.Description,
		Condition: auction.Condition,
		Status: auction.Status,
		Timestamp: auction.Timestamp.Unix(),

	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	return nil
}




