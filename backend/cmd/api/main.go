package main

import (
	"errors"
	"go-react-auth-backend/cmd/api/handlers"
	"go-react-auth-backend/internal/platform/database"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
}

type cfg struct {
	Port string
}

func run() error {

	// =========================================================================
	// setup server config
	c := cfg{Port: os.Getenv("PORT")} //  get port from docker
	if len(c.Port) == 0 {
		c.Port = "8090" // default to 8090 if not found
	}

	// =========================================================================
	// setup database

	log.Println("main : Started : Initializing database support")

	// get db vars
	//godotenv.Load() // for now load from .env
	du := os.Getenv("DB_USER")
	dp := os.Getenv("DB_PASSWORD")
	dh := os.Getenv("DB_HOST")
	dn := os.Getenv("DB_NAME")
	if len(du) == 0 || len(dp) == 0 || len(dn) == 0 {
		return errors.New("No database information provided!")
	}

	// connect to mongoDB
	db_cfg := database.Config{
		User:     du,
		Password: dp,
		Name:     dn,
		Host:     dh,
	}
	db, err := database.Open(db_cfg)
	if err != nil {
		return err
	}

	log.Println("main : Started : Successfully connected to database!")

	// =========================================================================
	// setup api

	log.Println("main : Started : Initializing API support...")
	router := handlers.API(db)

	log.Printf("main : API listening on port %s", c.Port)

	return http.ListenAndServe(":"+c.Port, router)
}
