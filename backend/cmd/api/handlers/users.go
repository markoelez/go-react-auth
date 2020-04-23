package handlers

import (
	"log"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("handlers : started : Called UsersHandler")

}
