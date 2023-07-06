package graph

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// import Todo model from internal/models
	"copilot-poc/internal/models"
)

//go:generate go run github.com/99designs/gqlgen generate

// UserRepository interface
type UserRepository interface {
	// define InsertOne
	InsertOne(input *models.User) (insertedID string, err error)
}

type TodoRepository interface {
	// define FindAll and InsertOne
	FindAll(ctx context.Context, query primitive.M) ([]*models.Todo, error)
	InsertOne(ctx context.Context, input *models.Todo) (insertedID string, err error)
}

type Resolver struct {
	TodoRepo TodoRepository
	UserRepo UserRepository
}
