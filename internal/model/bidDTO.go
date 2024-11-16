package model

import "time"


type Bid struct {
	Id 			string 		
	UserID 		string		
	AuctionID 	string		
	Amount 		float64		
	Timestamp 	time.Time   
}
