package handlers

import (
	"context"
	"encoding/json"
	"go-react-auth-backend/internal/platform/auth"
	"go-react-auth-backend/internal/user"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	db *mongo.Database
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// create NewUser var
	var nu user.NewUser
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		log.Fatal(err)
	}

	// push to DB
	var usr *user.User
	usr, err = user.Create(context.TODO(), u.db, nu, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	// send back user
	json.NewEncoder(w).Encode(usr)
	return
}

// authenticates users by validating email & password
func (u *User) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decode into NewUser struct
	var usr *user.NewUser
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Fatal(err)
	}

	// handle auth
	var authTokens *auth.TokenPair
	authTokens, err = user.Authenticate(context.TODO(), u.db, usr)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(authTokens)
}

// authenticates users by validating email & password
func (u *User) ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decode into User struct
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		// Error: Bearer token not in proper format
	}

	reqToken = strings.TrimSpace(splitToken[1])

	err := user.ValidateToken(context.TODO(), u.db, reqToken)
	if err != nil {
		log.Fatal(err)
	}

}
