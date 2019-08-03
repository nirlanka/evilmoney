package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/samsarahq/thunder/reactive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

// var posts = []Post{
// 	{Title: "first post!", Body: "I was here first!", CreatedAt: time.Now()},
// 	{Title: "graphql test", Body: "did you hear about Thunder?", CreatedAt: time.Now()},
// }

// registerPost registers the post type.
func (s *Server) RegisterPost(schema *schemabuilder.Schema) {
	obj := schema.Object("Post", Post{})

	obj.FieldFunc("age", func(ctx context.Context, p *Post) string {
		reactive.InvalidateAfter(ctx, 5*time.Second)
		return time.Since(p.CreatedAt).String()
	})
}

// registerQuery registers the root query type.
func (s *Server) RegisterQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()

	obj.FieldFunc("posts", func() []Post {
		// Pass these options to the Find method
		findOptions := options.Find()
		findOptions.SetLimit(2)

		// Here's an array in which you can store the decoded documents
		var results []Post

		// Passing bson.D{{}} as the filter matches all documents in the collection
		cur, err := s.Db.PostsCol.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Finding multiple documents returns a cursor
		// Iterating through the cursor allows us to decode documents one at a time
		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var elem Post
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
				fmt.Println(err)
			}

			results = append(results, elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		// Close the cursor once finished
		cur.Close(context.TODO())

		return results
	})
}

// registerMutation registers the root mutation type.
func (s *Server) RegisterMutation(schema *schemabuilder.Schema) {
	obj := schema.Mutation()

	obj.FieldFunc("echo", func(args struct{ Message string }) string {
		return args.Message
	})
}
