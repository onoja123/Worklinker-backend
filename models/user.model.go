package models

type User struct {
	ID        string   `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string   `json:"email" bson:"email"`
	Password  string   `json:"password" bson:"password"`
	UserType  string   `json:"user_type" bson:"user_type"`
	Interests []string `json:"interests" bson:"interests"`
	Language  string   `json:"language" bson:"language"`
}
