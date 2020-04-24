package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-react-auth-backend/internal/user"
	"log"
	"net/http"
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
	var authed *user.User
	authed, err = user.Authenticate(context.TODO(), u.db, usr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAuthenticated User: %+v\n", authed)
	json.NewEncoder(w).Encode(authed)
}
