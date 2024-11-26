package model

import (
	"time"

	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
)

type WinningInfoOutputDTO struct {
	Auction 		AuctionOutputDTO `json:"auction"`
	Bid				*BidOutputDTO 	 `json:"bid,omitempty"`  
}


type AuctionInputDTO struct {
	Id 	 			string 				
	ProductName 	string				`json:"productName" binding:"required,min=1"`					
	Category 		string				`json:"category" binding:"required,min=2"`
	Description 	string				`json:"description" binding:"required,min=10,max=200"`
	Condition 		ProductCondition	`json:"condition"`
	Status 			AuctionStatus		
	Timestamp 		time.Time		

}


type AuctionOutputDTO struct {
	Id 	 			string 				`json:"id"`		
	ProductName 	string				`json:"productName"`	
	Category 		string				`json:"category"`
	Description 	string				`json:"description"`
	Condition 		ProductCondition	`json:"condition"`
	Status 			AuctionStatus		`json:"status"`
	Timestamp 		time.Time			`json:"timestamp" time_format:"2006-01-02T15:04:05Z07:00"`

}

type ProductCondition int 
type AuctionStatus int 

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductCondition = iota
	Used
	Refurbished
)

func CreateAuctionInputDTO(productName, category, description string, condition ProductCondition) (*AuctionInputDTO, *internal_error.InternalError) {
	auction :=  &AuctionInputDTO{
		ProductName: productName,
		Category: category,
		Description: description,
		Condition: condition,
		Status: Active,
		Timestamp: time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}
	return auction, nil
}

func (au *AuctionInputDTO) Validate() *internal_error.InternalError {
	
	if  len(au.ProductName) <= 1 || 
		len(au.Category) <= 2 || 
		len(au.Description) < 10 &&
		   (au.Condition != New && 
			au.Condition != Used &&	
			au.Condition != Refurbished){ 
		return internal_error.NewBadRequestError("Invalid Auction object" )
	}
	
	return nil
}