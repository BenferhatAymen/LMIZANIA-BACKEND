package models

import (
	"time"

)

type Transaction struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Type        string             `json:"type" bson:"type"` // "expense" or "income"
	Amount      float64            `json:"amount" bson:"amount"`
	Date        time.Time          `json:"date" bson:"date"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
