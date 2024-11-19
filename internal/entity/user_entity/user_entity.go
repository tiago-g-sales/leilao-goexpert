package entity

import (
	"context"

	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)


type UserEntityMongo struct {
	Id 	 	string 		`bson:"_id"`
	Name 	string		`bson:"name"`

}

type UserRepositoryInterface interface {

	FindUserById(ctx context.Context,  userId string) (*model.UserOutputDTO, *internal_error.InternalError)
}