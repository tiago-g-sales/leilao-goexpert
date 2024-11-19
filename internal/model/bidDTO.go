package model

import "time"


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
