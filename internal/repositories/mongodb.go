package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
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

func NewMongoRepositotry(mongoURL, mongoDB string, mongoTimeout time.Duration) (ports.TodoRepository, error) {
	repo := &mongoRepository{
		database: mongoDB,
		timeout:  mongoTimeout,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (r *mongoRepository) CreateTodo(todo *domain.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	fmt.Print(todo)

	collection := r.client.Database(r.database).Collection("todos")
	_, err := collection.InsertOne(ctx, bson.M{
		"Content":   todo.Content,
		"CreatedAt": todo.CreatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) FindById(id string) (*domain.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	todo := &domain.Todo{}
	collection := r.client.Database(r.database).Collection("users")
	filter := bson.M{"ID": id}
	err := collection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
