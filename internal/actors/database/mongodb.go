package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/famesensor/playground-go-fiber-todonotes/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConn struct {
	TodoMongo *mongo.Client
	TodoDB    *mongo.Database
}

var mongoConn = &MongoConn{}

func NewMongoTodo(cfg *config.Config) (*MongoConn, error) {
	ctx := context.Background()

	var mongoURL string
	switch cfg.Env {
	case "local":
		mongoURL = fmt.Sprint("mongodb://", cfg.MongoHost, ":", cfg.MongoPort, "/", cfg.MongoDatabase)
		break
	case "dev":
		mongoURL = fmt.Sprint("mongodb://", cfg.MongoUser, ":", cfg.MongoPassword, "@", cfg.MongoHost, "/", cfg.MongoDatabase)
		break
	default:
		return nil, errors.New("unexpected run mode from config.NewMongoTodo")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	mongoConn.TodoMongo = client
	mongoConn.TodoDB = client.Database(cfg.MongoDatabase)
	return mongoConn, nil
}
