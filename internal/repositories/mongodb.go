package repositories

import (
	"context"
	"fmt"
	"time"

	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo := &mongoRepository{
		database: mongoDB,
		timeout:  mongoTimeout,
	}
	repo.client = client

	return repo, nil
}

func (r *mongoRepository) Create(ctx context.Context, todo *model.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	fmt.Print(todo)

	collection := r.client.Database(r.database).Collection("todos")
	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) FindById(ctx context.Context, id string) (*model.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	todo := &model.Todo{}
	collection := r.client.Database(r.database).Collection("todos")
	filter := bson.M{"ID": id}
	err := collection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]*model.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	filter := bson.D{{}}
	todo := []*model.Todo{}
	collection := r.client.Database(r.database).Collection("todos")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *mongoRepository) Update(ctx context.Context, id string, todo *model.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	todoID, _ := primitive.ObjectIDFromHex(id)

	query := bson.D{{Key: "_id", Value: todoID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "content", Value: todo.Content},
		{Key: "updatedAt", Value: time.Now()},
	}}}

	err := r.client.Database(r.database).Collection("todos").FindOneAndUpdate(ctx, query, update).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	todoID, _ := primitive.ObjectIDFromHex(id)

	query := bson.D{{Key: "_id", Value: todoID}}
	err := r.client.Database(r.database).Collection("todos").FindOneAndDelete(ctx, query).Err()

	if err != nil {
		return err
	}

	return nil
}
