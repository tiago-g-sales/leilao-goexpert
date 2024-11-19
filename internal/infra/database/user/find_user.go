package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	entity "github.com/tiago-g-sales/leilao-goexpert/internal/entity/user_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserRepository struct {
	Collection   *mongo.Collection
}

func NewUserRepository( database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserById(ctx context.Context,  userId string) (*model.UserOutputDTO, *internal_error.InternalError) {
	
	filter := bson.M{"_id": userId}
	var userEntityMongo entity.UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {	
			logger.Error(fmt.Sprintf("User not found with this id = %s", userId), err)
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("User not found with this id = %s", userId))	
		}
		
		logger.Error("Error trying to find user by userID", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by userID")	
	}

	modelUser := &model.UserOutputDTO{
		Id: userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return modelUser, nil
}