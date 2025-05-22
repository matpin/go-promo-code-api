package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	uri := os.Getenv("DB_URI")

	if uri == "" {
		return nil, fmt.Errorf("mongo is not connected")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("mongo is connected")
	return client, nil
}