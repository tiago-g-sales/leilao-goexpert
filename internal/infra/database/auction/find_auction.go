package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/auction_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context,  auctionId string) (*model.AuctionOutputDTO, *internal_error.InternalError){
	filter := bson.M{"_id": auctionId} 
	var auctionEntityMongo auction_entity.AuctionEntityMongo
	
	err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by auctionID = %s",auctionId ), err)
		return nil, internal_error.NewInternalServerError("Error trying to find auction by auctionID")
	}
	
	return &model.AuctionOutputDTO{
		Id: auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category: auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition: auctionEntityMongo.Condition,
		Status: auctionEntityMongo.Status,
		Timestamp: time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
	
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context, 
	status model.AuctionStatus, 
	category, productName string ) ([]model.AuctionOutputDTO, *internal_error.InternalError){
	
		filter := bson.M{} 
		if status != 0{
			filter["status"] = status
		}

		if category != ""{
			filter["category"] = category
		}

		if productName != ""{
			filter["productName"] =  primitive.Regex{Pattern: productName, Options: "i"}
		}

		cursor , err := ar.Collection.Find(ctx, filter)

		if err != nil {
			logger.Error("Error trying to find auctions", err)
			return nil, internal_error.NewInternalServerError("Error trying to find auctions")
		}

		defer cursor.Close(ctx)

		var auctions []auction_entity.AuctionEntityMongo
		err = cursor.All(ctx, &auctions)
		if err != nil {
			logger.Error("Error trying to find auctions", err)
			return nil, internal_error.NewInternalServerError("Error trying to find auctions")
		}

		var auctionsDTO []model.AuctionOutputDTO
		for _, auction := range auctions {
			auctionsDTO = append(auctionsDTO, model.AuctionOutputDTO{
				Id: auction.Id,
				ProductName: auction.ProductName,
				Category: auction.Category,
				Description: auction.Description,
				Condition: auction.Condition,
				Status: auction.Status,
				Timestamp: time.Unix(auction.Timestamp, 0),
			})
		}	
		return auctionsDTO, nil
}