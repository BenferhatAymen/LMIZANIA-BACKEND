package models

import (
	"time"
)

type Goal struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserID        string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title         string    `json:"title" bson:"title"`
	Description   string    `json:"description" bson:"description"`
	TotalAmount   float64   `json:"total_amount" bson:"total_amount"`
	CurrentAmount float64   `json:"current_amount" bson:"current_amount"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
