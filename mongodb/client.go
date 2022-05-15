package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-mongodb-implementation/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewClient(ctx context.Context, configuration config.MongoDBConfig, timeout time.Duration) (MongoClient, error) {
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s", configuration.Username, configuration.Password, configuration.Host)

	clientOptions := options.Client().SetConnectTimeout(timeout).ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return MongoClient{}, err
	}

	return MongoClient{Client: client}, nil
}

func (mc *MongoClient) GetClient() *mongo.Client {
	return mc.Client
}
