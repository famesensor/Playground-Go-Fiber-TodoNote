package repositories

import (
	"context"
	"time"

	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	interfaces "github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/famesensor/playground-go-fiber-todonotes/pkg/errs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	TodoClient *mongo.Client
	TodoDB     *mongo.Database
}

func NewMongoRepositotry(TodoClient *mongo.Client, TodoDB *mongo.Database) interfaces.TodoRepository {
	return &mongoRepository{TodoClient, TodoDB}
}

func (r *mongoRepository) Create(ctx context.Context, todo *model.Todo) error {
	collection := r.TodoDB.Collection("todos")
	_, err := collection.InsertOne(ctx, bson.M{
		"title":       todo.Title,
		"description": todo.Description,
		"status":      todo.Status,
		"priority":    todo.Priority,
		"createdAt":   time.Now(),
		"updatedAt":   time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) FindById(ctx context.Context, id string) (*model.Todo, error) {
	todo := &model.Todo{}

	todoID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": todoID}

	collection := r.TodoDB.Collection("todos")
	err := collection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.DocumentNotFound
		}
		return nil, err
	}
	return todo, nil
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]*model.Todo, error) {
	todo := []*model.Todo{}
	filter := bson.D{{}}

	collection := r.TodoDB.Collection("todos")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *mongoRepository) UpdateTodo(ctx context.Context, id string, todo *model.Todo) error {
	todoID, _ := primitive.ObjectIDFromHex(id)

	query := bson.D{{Key: "_id", Value: todoID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: todo.Title},
		{Key: "desceiption", Value: todo.Description},
		{Key: "priority", Value: todo.Priority},
	}}}

	err := r.TodoDB.Collection("todos").FindOneAndUpdate(ctx, query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errs.DocumentNotFound
		}
		return err
	}

	return nil
}

func (r *mongoRepository) UpdateStatus(ctx context.Context, id string) error {
	todoID, _ := primitive.ObjectIDFromHex(id)

	query := bson.D{{Key: "_id", Value: todoID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "status", Value: "success"},
	}}}

	err := r.TodoDB.Collection("todos").FindOneAndUpdate(ctx, query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errs.DocumentNotFound
		}
		return err
	}

	return nil
}
