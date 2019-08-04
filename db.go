package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	PostsCol *mongo.Collection
}

func GetDb() *Db {
	mongo_url, exists := os.LookupEnv("MONGODB_URI")
	var db_url string

	if !exists {
		mongo_url = "mongodb://localhost:27017/"
		db_url = "evilmoney"
	} else {
		_url := strings.Split(mongo_url, "/")
		db_url = _url[len(_url)-1]
	}

	// Connect to database
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(mongo_url)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	return &Db{
		PostsCol: client.Database(db_url).Collection("posts"),
	}
}
