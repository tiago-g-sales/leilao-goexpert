package bid_usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"github.com/tiago-g-sales/leilao-goexpert/internal/infra/database/bid"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)


type BidUseCase struct {

	BidRepository  bid.BidRepository
	timer *time.Timer
	maxBatchSize int 
	batchInsertInterval time.Duration 
	bidCannel chan model.BidInputDTO

}

type BidUseCaseInterface interface {	
	CreateBid(ctx context.Context, bid model.BidInputDTO) *internal_error.InternalError
	FindWinningBidByAuctionId(ctx  context.Context, auctionId string)(*model.BidOutputDTO, *internal_error.InternalError)
	FindBidByAuctionId(ctx  context.Context, auctionId string) ([]model.BidOutputDTO, *internal_error.InternalError)
}

func NewBidUseCase(bidRepository  bid.BidRepository ) BidUseCaseInterface{
	
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()	

	bidUseCase:= &BidUseCase{
		BidRepository: bidRepository,
		maxBatchSize: maxBatchSize,
		batchInsertInterval: maxSizeInterval,
		timer: time.NewTimer(maxSizeInterval),
		bidCannel: make(chan model.BidInputDTO, maxBatchSize),
	}
	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

var bidBatch []model.BidInputDTO 

func (bu *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func ()  {
		defer close(bu.bidCannel)
		for {
			select {
			case bidEntity, ok :=  <- bu.bidCannel:
				if !ok {				
					if len(bidBatch) > 0 {
						if err :=  bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("Error trying to process bid batch list", err)
						}
					}
					return
				}			
				bidBatch = append(bidBatch, bidEntity)
				if len(bidBatch) >= bu.maxBatchSize {
					if err :=  bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("Error trying to process bid batch list", err)
					}
					bidBatch = nil
					bu.timer.Reset(bu.batchInsertInterval)
				}
			case <-bu.timer.C:
				if err :=  bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("Error trying to process bid batch list", err)
				}
				bidBatch = nil
				bu.timer.Reset(bu.batchInsertInterval)

			}

		}
	}()
}


func (bu *BidUseCase) CreateBid(ctx context.Context, bid model.BidInputDTO ) ( *internal_error.InternalError) {
	
	bidEntity, err := model.CreateBid(bid.UserID, bid.AuctionID, bid.Amount) 
	if err != nil {
		return err
	}

	bu.bidCannel <- *bidEntity

	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := viper.GetString("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}
	return duration
}

func getMaxBatchSize() int {
	value, err := strconv.Atoi(viper.GetString("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}
	return value
}