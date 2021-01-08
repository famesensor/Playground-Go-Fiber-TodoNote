package repositories

import (
	"context"

	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	interfaces "github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	TodoClient *mongo.Client
	TodoDB     *mongo.Database
}

func NewMongoRepositotry(TodoClient *mongo.Client, TodoDB *mongo.Database) (interfaces.TodoRepository, error) {
	return &mongoRepository{TodoClient, TodoDB}, nil
}

func (r *mongoRepository) Create(ctx context.Context, todo *model.Todo) error {

	collection := r.TodoDB.Collection("todos")
	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) FindById(ctx context.Context, id string) (*model.Todo, error) {
	todo := &model.Todo{}
	collection := r.TodoDB.Collection("todos")
	filter := bson.M{"ID": id}
	err := collection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]*model.Todo, error) {
	filter := bson.D{{}}
	todo := []*model.Todo{}
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

func (r *mongoRepository) Update(ctx context.Context, id string, todo *model.Todo) error {
	// todoID, _ := primitive.ObjectIDFromHex(id)

	// query := bson.D{{Key: "_id", Value: todoID}}
	// update := bson.D{{Key: "$set", Value: bson.D{
	// 	{Key: "content", Value: todo.Content},
	// 	{Key: "updatedAt", Value: time.Now()},
	// }}}

	// err := r.TodoDB.Collection("todos").FindOneAndUpdate(ctx, query, update).Err()

	// if err != nil {
	// 	return err
	// }

	return nil
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	// todoID, _ := primitive.ObjectIDFromHex(id)

	// query := bson.D{{Key: "_id", Value: todoID}}
	// err := r.TodoDB.Collection("todos").FindOneAndDelete(ctx, query).Err()

	// if err != nil {
	// 	return err
	// }

	return nil
}
