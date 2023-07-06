package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Todo is the todo model
type Todo struct {
	ID     string `bson:"_id,omitempty"`
	Text   string `bson:"text"`
	UserID string `bson:"userId"`
	Done   bool   `bson:"done"`
}

// Define todoRepository struct, should implement TodoRepository interface
type todoRepository struct {
	// add mongoDB collection as dependency
	TodoColl *mongo.Collection
}

// Define NewTodoRepository function
func NewTodoRepository(todoColl *mongo.Collection) *todoRepository {
	return &todoRepository{
		TodoColl: todoColl,
	}
}

// Define FindAll method
func (r *todoRepository) FindAll(ctx context.Context, query primitive.M) ([]*Todo, error) {
	// fetch todos from MongoDB
	todos := []*Todo{}
	// query mongoDB for todos with userID
	cur, err := r.TodoColl.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	// decode todos from mongoDB cursor
	if err := cur.All(ctx, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

// Define InsertOne method
func (r *todoRepository) InsertOne(ctx context.Context, input *Todo) (insertedID string, err error) {
	// insert input into MongoDB
	doc, err := r.TodoColl.InsertOne(ctx, input)
	if err != nil {
		return "", err
	}
	return doc.InsertedID.(primitive.ObjectID).Hex(), nil
}
