package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	PostsCol *mongo.Collection
}

func GetDb() *Db {
	mongo_url, exists := os.LookupEnv("DB_URL")
	var db_url string

	if !exists {
		mongo_url = "mongodb://localhost:1000/"
		db_url = "evilmoney"
	} else {
		_url := strings.Split(mongo_url, "/")

		mongo_url = strings.Join(_url[0:len(_url)-1], "/")
		db_url = _url[len(_url)-1]
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
