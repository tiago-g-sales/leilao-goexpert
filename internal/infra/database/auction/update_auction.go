package auction

import (
	"context"
	"fmt"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)



func (ar *AuctionRepository) UpdateAuctionById(ctx context.Context,  auction model.AuctionInputDTO) *internal_error.InternalError{
		filter := bson.M{"_id": auction.Id} 

	update := bson.M{"$set": bson.M{"status": auction.Status  }}


	_, err := ar.Collection.UpdateOne(ctx, filter, update) 
	if err != nil {
		fmt.Println(err)
		logger.Error(fmt.Sprintf("Error trying to update auction by auctionID = %s",auction.Id ), err)
		return internal_error.NewInternalServerError("Error trying to update auction by auctionID")
	}
	
	return nil
	
}
