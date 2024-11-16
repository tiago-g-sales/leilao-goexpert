package bid_entity


type BidEntityMongo struct {
	Id 			string 		`bson:"_id"`
	UserID 		string		`bson:"user_id"`
	AuctionID 	string		`bson:"auction_id"`
	Amount 		float64		`bson:"amount"`
	Timestamp   int64   	`bson:"timestamp"`
}


