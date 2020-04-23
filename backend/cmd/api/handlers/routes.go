package handlers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func API(db *mongo.Client) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/api/")

	// users routes
	s := r.PathPrefix("/users").Subrouter()

	s.HandleFunc("/getAll", UsersHandler).Methods("GET")

	return r

}
