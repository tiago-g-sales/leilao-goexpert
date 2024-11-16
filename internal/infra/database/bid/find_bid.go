package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/bid_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string)([]model.Bid, *internal_error.InternalError){

	filter := bson.M{"auctionID": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionID %s", auctionId ), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auctionID %s", auctionId))
	}
	var entities_bid []bid_entity.BidEntityMongo
	err = cursor.All(ctx, &entities_bid)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionID %s", auctionId ), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auctionID %s", auctionId))
	}

	var bidList []model.Bid
	for _, bid := range entities_bid {
		bidList = append(bidList, model.Bid{
			Id: bid.Id,
			UserID: bid.UserID,
			AuctionID: bid.AuctionID,
			Amount: bid.Amount,
			Timestamp: time.Unix(bid.Timestamp, 0),
		})

	}

	return bidList, nil

}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string)(*model.Bid, *internal_error.InternalError){

	filter := bson.M{"auctionID": auctionId}
	var bid bid_entity.BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})
	err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bid)
	if err != nil {
		logger.Error("Error trying to find the auction winner", err)
		return nil, internal_error.NewInternalServerError("Error trying to find the auction winner")	
	}
	return &model.Bid{
		Id: bid.Id,
		UserID: bid.UserID,
		AuctionID: bid.AuctionID,
		Amount: bid.Amount,
		Timestamp: time.Unix(bid.Timestamp, 0),
	}, nil
	


}