package service

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/auction_usecase"
)

const (
	DATE_TIME_FORMAT = "2006-01-02 15:04:05"
)

type AuctionService struct {
	auctionUseCase  auction_usecase.AuctionUseCaseInterface
}


func NewAuctionService(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionService {
	return &AuctionService{
		auctionUseCase: auctionUseCase,
	}
}




func (as *AuctionService) UpdateAuctionsToEnd(){
	status := 0
	auctionEndTime := viper.GetString("AUCTION_END_TIME")
	duration, errConv := time.ParseDuration(auctionEndTime)
	 
	if errConv != nil {
		return 
	}

	for {

		timeNow :=  time.Now()
		auctions, err :=  as.auctionUseCase.FindAuctions(context.Background(), model.AuctionStatus(status), "", "")
		if err != nil {
			return
		}
		totalAuctions := 0 
		for _, auction := range auctions {
			diff := timeNow.Sub(auction.Timestamp)
			endTimeEstimate := auction.Timestamp.Add(duration)
			if   diff >= duration{

				auctionIn := model.AuctionInputDTO{		
					Id: auction.Id,
					Status: model.Completed,
				}
				as.auctionUseCase.UpdateAuctionById(context.Background(), auctionIn )				
				fmt.Print("=====================================================================================\n")
				fmt.Printf("Leilão encerrado com sucesso ID: %s \n" , auction.Id )
				fmt.Printf("Data e hora de criação do Leilão: %s \n" , auction.Timestamp.Format(DATE_TIME_FORMAT) )
				fmt.Printf("Data e Hora atual               : %s \n", timeNow.Format(DATE_TIME_FORMAT))
				fmt.Printf("Paramentro definido para a duração do Leilão: %s \n" , auctionEndTime )
				fmt.Print("=====================================================================================\n")
				totalAuctions++	
			}else {
				fmt.Print("=====================================================================================\n")
				fmt.Printf("Leilão ativo ID: %s \n" , auction.Id )				
				fmt.Printf("Paramentro definido para a duração do Leilão: %s \n" , auctionEndTime )
				fmt.Printf("Data e Hora atual       : %s \n", timeNow.Format(DATE_TIME_FORMAT))
				fmt.Printf("Data e hora de criação  : %s \n", auction.Timestamp.Format(DATE_TIME_FORMAT) )
				fmt.Printf("Data e Hora encerramento: %s \n", endTimeEstimate.Format(DATE_TIME_FORMAT))
				fmt.Print("=====================================================================================\n")
			}
			
			
		}
		
		if totalAuctions > 0 {
			fmt.Println("Todos os Leiloes foram encerrados")
			fmt.Print("=====================================================================================\n")
		}
	
		time.Sleep(10 * time.Second)
	}
}