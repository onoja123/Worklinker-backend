package models

type Wallet struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	UserID    string  `json:"user_id" bson:"user_id"`
	Balance   float64 `json:"balance" bson:"balance"`
	Currency  string  `json:"currency" bson:"currency"`
	CreatedAt int64   `json:"created_at" bson:"created_at"`
	UpdatedAt int64   `json:"updated_at" bson:"updated_at"`
}
