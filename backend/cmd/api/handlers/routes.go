package handlers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func API(db *mongo.Client) *mux.Router {
	r := mux.NewRouter()

	// users routes
	s := r.PathPrefix("/users").Subrouter()

	s.HandleFunc("/getAll", UsersHandler).Methods("GET")

	return r

}
