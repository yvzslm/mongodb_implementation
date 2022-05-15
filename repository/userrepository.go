package repository

import (
	"context"

	"github.com/go-mongodb-implementation/model"
	"github.com/go-mongodb-implementation/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTION_NAME = "user"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Insert(ctx context.Context, user *model.User) error
}

type userRepository struct {
	mongoClient  mongodb.MongoClient
	databaseName string
}

func (ur *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	var client = ur.mongoClient.GetClient()

	var records []*model.User

	cur, err := client.Database(ur.databaseName).Collection(COLLECTION_NAME).Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (ur *userRepository) Insert(ctx context.Context, user *model.User) error {
	var client = ur.mongoClient.GetClient()

	_, err := client.Database(ur.databaseName).Collection(COLLECTION_NAME).InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(client mongodb.MongoClient, databaseName string) UserRepository {
	return &userRepository{mongoClient: client, databaseName: databaseName}
}
