package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	PostsCol *mongo.Collection
}

func GetDb() *Db {
	mongo_url, exists := os.LookupEnv("DB_URL")
	if !exists {
		mongo_url = "mongodb://localhost:1000/"
	}

	db_url, exists := os.LookupEnv("DB_NAME")
	if !exists {
		db_url = "evilmoney"
	}

	// Connect to database
	clientOptions := options.Client().ApplyURI(mongo_url)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	return &Db{
		PostsCol: client.Database(db_url).Collection("posts"),
	}
}
