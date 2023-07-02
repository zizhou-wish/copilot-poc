package graph

import "copilot-poc/graph/model"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// add Todo model as dependency
type Resolver struct {
	todos []*model.Todo
}
