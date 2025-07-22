package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"first_name" validate:"required,min=3,max=50"`
	LastName     string             `json:"last_name" validate:"required,min=3,max=50"`
	Username     string             `json:"username" validate:"required,min=5,max=25"`
	Token        string             `json:"token"`
	Password     string             `json:"password"`
	UserType     string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken string             `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}
