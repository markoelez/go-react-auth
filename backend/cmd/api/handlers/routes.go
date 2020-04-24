package handlers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func API(db *mongo.Database) *mux.Router {
	r := mux.NewRouter()

	// user routes
	s := r.PathPrefix("/users").Subrouter()
	u := User{db: db}

	s.HandleFunc("/register", u.Create).Methods("POST")
	s.HandleFunc("/authenticate", u.Authenticate).Methods("POST")

	return r

}
