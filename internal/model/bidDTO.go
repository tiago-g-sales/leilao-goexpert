package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
)


type BidInputDTO struct {
	Id 			string 			`json:"id"` 		
	UserID 		string			`json:"userID"`
	AuctionID 	string			`json:"auctionID"`
	Amount 		float64			`json:"amount"`
	Timestamp 	time.Time   	`json:"timestamp"`
}


type BidOutputDTO struct {
	Id 			string 			`json:"id"`
	UserID 		string			`json:"userID"`
	AuctionID 	string			`json:"auctionID"`
	Amount 		float64			`json:"amount"`
	Timestamp 	time.Time   	`json:"timestamp" time_format:"2006-01-02T15:04:05Z07:00"`
}


func CreateBid( userId, auctionId string, amount float64) (*BidInputDTO, *internal_error.InternalError) {
	bid := &BidInputDTO{
		Id: uuid.New().String(),	
		UserID: userId,
		AuctionID: auctionId,
		Amount: amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validadte(); err != nil {
		return nil, err
	}

	return bid, nil


}

func (b *BidInputDTO) Validadte() *internal_error.InternalError {
	if err := uuid.Validate(b.UserID); err != nil {
		return internal_error.NewBadRequestError("UserId is not valid ID")
	}
	if err := uuid.Validate(b.AuctionID); err != nil {
		return internal_error.NewBadRequestError("AuctionID is not valid ID")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Amount must be greater than zero")
	}

	return nil
}