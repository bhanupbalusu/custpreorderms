package preorder_metadata

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

type MongoRepository struct {
	Client  *mongo.Client
	DB      string
	Timeout time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (repo.PreOrderRepoInterface, error) {
	repo := &MongoRepository{
		DB:      mongoDB,
		Timeout: time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.Client = client
	return repo, nil
}
