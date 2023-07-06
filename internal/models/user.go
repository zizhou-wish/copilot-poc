package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockery --name UserRepository --output ./mocks --filename mock_UserRepository.go

// UserRepository interface
type UserRepository interface {
	// define InsertOne
	InsertOne(ctx context.Context, input *User) (insertedID string, err error)
}

// User is the user model
type User struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

// Define userRepository struct, should implement UserRepository interface
type userRepository struct {
	// add mongoDB collection as dependency
	UserColl *mongo.Collection
}

// Define NewUserRepository function
func NewUserRepository(userColl *mongo.Collection) UserRepository {
	return &userRepository{
		UserColl: userColl,
	}
}

// Define InsertOne method
func (r *userRepository) InsertOne(context context.Context, input *User) (insertedID string, err error) {
	// insert input into MongoDB
	doc, err := r.UserColl.InsertOne(context, input)
	if err != nil {
		return "", err
	}
	return doc.InsertedID.(string), nil
}
