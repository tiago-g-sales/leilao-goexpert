package bid

import (
	"context"
	"sync"

	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"github.com/tiago-g-sales/leilao-goexpert/internal/entity/bid_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/database/auction"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)


type BidRepository struct {

	Collection 			*mongo.Collection
	AuctionRepository	*auction.AuctionRepository
	
}

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		Collection: database.Collection("bids"),
		AuctionRepository: auctionRepository,
	}
}



func (bd *BidRepository) CreateBid( ctx context.Context, bidList []model.BidInputDTO)(*internal_error.InternalError) {
	var wg sync.WaitGroup

	for _, bid := range bidList {
		wg.Add(1)

		go func (bidValue model.BidInputDTO) {
			defer wg.Done()
			auction, err := bd.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionID)
			if err != nil {
				logger.Error("Error trying to find auction by id ", err)
				return 
			}
			if auction.Status != model.Active {
				return
			}
			
			bid_entity := &bid_entity.BidEntityMongo{
				Id: bidValue.Id,
				UserID: bidValue.UserID,
				AuctionID: bidValue.AuctionID,
				Amount: bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}
			if _ , err := bd.Collection.InsertOne(ctx, bid_entity); err != nil {
				logger.Error("Error trying to insert bid", err)
				return
			}
		}(bid)
	
	}


	wg.Wait()
	return nil

	
}