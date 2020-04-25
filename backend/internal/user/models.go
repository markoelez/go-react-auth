package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user with access to api
type User struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string              `json:"name"`
	Email         string              `json:"email"`
	PasswordHash  string              `json:"password_hash" bson:"password_hash"`
	DateCreated   time.Time           `json:"date_created" bson:"date_created"`
	RefreshTokens []string            `json:"refresh_tokens" bson:"refresh_tokens"`
}

// for creating new users
type NewUser struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}
