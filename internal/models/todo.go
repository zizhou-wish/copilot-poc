package models

// Todo is the todo model
type Todo struct {
	ID     string `bson:"_id,omitempty"`
	Text   string `bson:"text"`
	UserID string `bson:"userId"`
}
