package models

import "go.mongodb.org/mongo-driver/mongo"

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
func NewUserRepository(userColl *mongo.Collection) *userRepository {
	return &userRepository{
		UserColl: userColl,
	}
}

// Define InsertOne method
func (r *userRepository) InsertOne(input *User) (insertedID string, err error) {
	// insert input into MongoDB
	doc, err := r.UserColl.InsertOne(nil, input)
	if err != nil {
		return "", err
	}
	return doc.InsertedID.(string), nil
}
