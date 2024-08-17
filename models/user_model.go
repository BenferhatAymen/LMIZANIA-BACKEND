package models

type User struct {
	ID        string `json:"id,omitempty" bson:"_id"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	FamilyName string `json:"family_name,omitempty" bson:"family_name"`
	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
	IsVerified bool `json:"is_verified,omitempty" bson:"is_verified"`
}