package models

// User is the user model
type User struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}
