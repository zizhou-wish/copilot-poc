package graph

import (
	"copilot-poc/internal/graph/model"

	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// add MongoDB client as a dependency
type Resolver struct {
	TodoCollection *mongo.Collection
	UserCollection *mongo.Collection
	todos          []*model.Todo
}
