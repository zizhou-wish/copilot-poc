package graph

import "copilot-poc/internal/models"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	TodoRepo models.TodoRepository
	UserRepo models.UserRepository
}
