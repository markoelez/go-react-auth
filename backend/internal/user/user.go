package user

import (
	"context"
	"errors"
	"go-react-auth-backend/internal/platform/auth"
	"go-react-auth-backend/internal/platform/database"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const userCollection = "users"

// query errors
var (
	// used if user does not exist
	ErrNotFound = errors.New("User not found")
	// used if id is incorrect
	ErrInvalidID = errors.New("ID is not in the proper form")
	// used if auth failed for any reason
	ErrAuthFailure = errors.New("Authentication failed")
	// used if user performs forbidden action
	ErrForbidden = errors.New("This action is forbidden")
)

// creates a new user and inserts into DB
func Create(ctx context.Context, db *mongo.Database, n NewUser, now time.Time) (*User, error) {

	// first hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(n.Password), 5)
	if err != nil {
		return nil, err
	}

	// create user
	u := User{
		Name:         n.Name,
		Email:        n.Email,
		PasswordHash: string(hash),
		DateCreated:  now.UTC(),
	}
	// upload to db
	uc := database.GetCollection(db, userCollection)
	// check if email exists first
	var tmp interface{}
	err = uc.FindOne(ctx, bson.D{{"email", u.Email}}).Decode(&tmp)
	if err == nil {
		// email exists
		return nil, err
	}

	// create user
	_, err = uc.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func Authenticate(ctx context.Context, db *mongo.Database, u *NewUser) (*auth.TokenPair, error) {

	log.Println("user : started : Checking if account exists")

	// check that acocunt exists
	uc := database.GetCollection(db, userCollection)
	var dbu *User
	err := uc.FindOne(ctx, bson.D{{"email", u.Email}}).Decode(&dbu)
	if err != nil {
		return nil, err
	}

	log.Println("user : started : Validating password hash")

	// compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(dbu.PasswordHash), []byte(u.Password))
	if err != nil {
		return nil, err
	}

	log.Println("user : started : Generating auth tokens")

	// generate access & refresh tokens
	t, err := auth.GetTokenPair(u.Email, "SECRET")

	// add refresh tokens to user for more convenient revocation
	if dbu.RefreshTokens == nil || len(dbu.RefreshTokens) == 0 {
		dbu.RefreshTokens = make([]string, 0, 20)
	}
	dbu.RefreshTokens = append(dbu.RefreshTokens, t.RefreshToken)

	// add token to user in DB
	_, err = uc.UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": dbu.ID}}, bson.M{"$set": bson.M{"access_token": t.AccessToken, "refresh_tokens": dbu.RefreshTokens}}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}

	log.Println("user : started : Authentication successful, returning user with auth tokens")

	return t, nil
}

// validates an access token
func ValidateToken(ctx context.Context, db *mongo.Database, token string) error {
	_, err := auth.ValidateToken(token, "SECRET")
	if err != nil {
		return err
	}

	return nil
}

// retrieves a new access token
//func RetrieveToken() {}

// logs a user out by clearing their local & remote access tokens and refresh tokens
//func Logout() {}
