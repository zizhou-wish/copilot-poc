package graph

import (
	"context"
	"testing"

	// import graph/model
	"copilot-poc/internal/graph/model"
	mongoModels "copilot-poc/internal/models"
	"copilot-poc/internal/models/mocks"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMutationResolver_CreateTodo(t *testing.T) {
	// create a new mock Todo repository
	todoRepo := mocks.NewTodoRepository(t)

	// create a new resolver
	resolver := &Resolver{
		TodoRepo: todoRepo,
	}

	// create a new context
	ctx := context.Background()

	userID := primitive.NewObjectID()

	// create graphQL mutation input
	input := model.NewTodo{
		Text:   "test todo",
		UserID: userID.Hex(),
	}

	// create todo repo input
	repoInput := &mongoModels.Todo{
		Text:   input.Text,
		UserID: userID,
	}

	// todoRepo.On("InsertOne", mock.Anything, repoInput).Return("insertedID", nil)
	// use mockery mock to return a mock insertedID
	todoRepo.EXPECT().InsertOne(ctx, repoInput).Return("insertedID", nil)

	// call the resolver function
	result, err := resolver.Mutation().CreateTodo(ctx, input)

	// check the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "insertedID", result.ID)
	assert.Equal(t, input.Text, result.Text)
	assert.False(t, result.Done)
	assert.Equal(t, input.UserID, result.UserID)
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

	// create user repo input
	repoInput := &mongoModels.User{
		Name: input.Name,
	}

	// use mockery mock to return a mock insertedID
	userRepo.EXPECT().InsertOne(ctx, repoInput).Return("insertedID", nil)

	// call the resolver function
	result, err := resolver.Mutation().CreateUser(ctx, input)

	// check the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "insertedID", result.ID)
	assert.Equal(t, input.Name, result.Name)
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
	userID := primitive.NewObjectID()

	// create a new mock Todo object
	mockTodo := &mongoModels.Todo{
		ID:     "test todo",
		Text:   "test todo",
		UserID: userID,
		Done:   false,
	}

	// use mockery mock to return a mock array of Todos
	todoRepo.EXPECT().FindAll(ctx, primitive.M{"userId": userID}).Return([]*mongoModels.Todo{mockTodo}, nil)

	// call the resolver function
	result, err := resolver.Query().Todos(ctx, userID.Hex())

	// check the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
	assert.Equal(t, mockTodo.ID, result[0].ID)
	assert.Equal(t, mockTodo.Text, result[0].Text)
	assert.Equal(t, mockTodo.UserID.Hex(), result[0].UserID)
	assert.Equal(t, mockTodo.Done, result[0].Done)
}
