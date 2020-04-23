package main

import (
	"context"
	"go-react-auth-backend/cmd/api/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	err := godotenv.Load()
	if err != nil {
		return err
	}
	c := cfg{Port: os.Getenv("PORT")} // get port from docker env (for now use .env file)

	// =========================================================================
	// setup database

	log.Println("main : Started : Initializing database support")

	// connect to mongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://def_user:TESTER@algoprepdb-vopr0.mongodb.net/t    est?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}

	// test connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("main : Started : Connected to database...")

	// =========================================================================
	// setup api

	log.Println("main : Started : Initializing API support")
	router := handlers.API(client)

	log.Printf("main : API listening on port %s", c.Port)

	return http.ListenAndServe(":"+c.Port, router)
}
