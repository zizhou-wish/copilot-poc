package graph

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// add user and todo mongoDB collection as dependencies
type Resolver struct {
	TodoColl *mongo.Collection
	UserColl *mongo.Collection
}
