package mongodb

import (
	"context"

	"github.com/spf13/viper"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const (

	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
	
)

func NewMongoDBConnection(ctx context.Context)(*mongo.Database, error){
	mongoURL := viper.GetString(MONGODB_URL)
	mongoDatabase := viper.GetString(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Error trying to connect to mongodb database", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("Error trying to ping to mongodb database", err)
		return nil, err
	}
	return client.Database(mongoDatabase), nil

} 