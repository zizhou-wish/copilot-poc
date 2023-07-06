package graph

import (
	"context"
	"testing"

	// import graph/model
	"copilot-poc/internal/graph/model"
	"copilot-poc/internal/models"
	"copilot-poc/internal/models/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMutationResolver_CreateTodo(t *testing.T) {
	// create a new mock Todo repository
	todoRepo := &mocks.TodoRepository{}

	// create a new resolver
	resolver := &Resolver{
		TodoRepo: todoRepo,
	}

	// create a new context
	ctx := context.Background()

	// create a new input
	input := model.NewTodo{
		Text:   "test todo",
		UserID: "test user",
	}

	// set up the mock Todo repository
	todoRepo.On("InsertOne", mock.Anything, input).Return("insertedID", nil)

	// call the resolver function
	result, err := resolver.Mutation().CreateTodo(ctx, input)

	// check the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "insertedID", result.ID)
	assert.Equal(t, input.Text, result.Text)
	assert.False(t, result.Done)
	assert.Equal(t, input.UserID, result.UserID)

	// assert that the mock Todo repository was called with the correct arguments
	todoRepo.AssertCalled(t, "InsertOne", mock.Anything, input)
}

func TestMutationResolver_CreateUser(t *testing.T) {
	// create a new mock User repository
	userRepo := &mocks.UserRepository{}

	// create a new resolver
	resolver := &Resolver{
		UserRepo: userRepo,
	}

	// create a new context
	ctx := context.Background()

	// create a new input
	input := model.NewUser{
		Name: "test user",
	}

	// set up the mock User repository
	userRepo.On("InsertOne", mock.Anything, input).Return("insertedID", nil)

	// call the resolver function
	result, err := resolver.Mutation().CreateUser(ctx, input)

	// check the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "insertedID", result.ID)
	assert.Equal(t, input.Name, result.Name)

	// assert that the mock User repository was called with the correct arguments
	userRepo.AssertCalled(t, "InsertOne", mock.Anything, input)
}

func TestQueryResolver_Todos(t *testing.T) {
	// create a new mock Todo repository
	todoRepo := &mocks.TodoRepository{}

	// create a new resolver
	resolver := &Resolver{
		TodoRepo: todoRepo,
	}

	// create a new context
	ctx := context.Background()

	// create a new user ID
	userID := "test user"

	// create a new mock Todo object
	mockTodo := &models.Todo{
		ID:     "test todo",
		Text:   "test todo",
		UserID: userID,
		Done:   false,
	}

	// set up the mock Todo repository
	todoRepo.On("FindAll", mock.Anything, mock.Anything).Return([]*models.Todo{mockTodo}, nil)

	// call the resolver function
	result, err := resolver.Query().Todos(ctx, userID)

	// check the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
	assert.Equal(t, mockTodo.ID, result[0].ID)
	assert.Equal(t, mockTodo.Text, result[0].Text)
	assert.Equal(t, mockTodo.UserID, result[0].UserID)
	assert.Equal(t, mockTodo.Done, result[0].Done)

	// assert that the mock Todo repository was called with the correct arguments
	todoRepo.AssertCalled(t, "FindAll", mock.Anything, mock.Anything)
}
