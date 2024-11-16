package auction_entity

import (
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)

type AuctionEntityMongo struct {
	Id 	 			string 					`bson:"_id"`
	ProductName 	string					`bson:"product_name"`
	Category 		string					`bson:"categoria"`
	Description 	string		 			`bson:"description"`
	Condition 		model.ProductCondition	`bson:"condition"`
	Status 			model.AuctionStatus  	`bson:"status"`
	Timestamp 		int64	 				`bson:"timestamp"`

}


