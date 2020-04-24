package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	User     string
	Password string
	Name     string
	Host     string
}

// connect to mongoDB
func Open(cfg Config) (*mongo.Database, error) {
	db_url := fmt.Sprintf("mongodb+srv://%s:%s@%s-vopr0.mongodb.net/test?retryWrites=true&w=majority", cfg.User, cfg.Password, cfg.Host)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))
	if err != nil {
		return nil, err
	}

	// test connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.Name), nil
}

// return a database collection
func GetCollection(db *mongo.Database, collection string) *mongo.Collection {
	// just return collection
	return db.Collection(collection)
}
