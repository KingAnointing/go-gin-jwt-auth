package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string             `json:"first_name"`
	LastName     *string             `json:"last_name"`
	Email        *string             `json:"email"`
	Password     *string             `json:"password"`
	Phone        *string             `json:"phone"`
	UserType     *string             `json:"user_type"`
	Token        *string             `json:"token"`
	RefreshToken *string             `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string          `json:"user_id"`
}
