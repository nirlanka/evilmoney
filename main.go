package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// server is our graphql server.
type Server struct {
	Client *mongo.Client
}

// schema builds the graphql schema.
func (s *Server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.RegisterPost(builder)

	s.RegisterQuery(builder)
	s.RegisterMutation(builder)
	return builder.MustBuild()
}

func serveGraphql() {
	// Connect to database
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	// Instantiate a server, build a server, and serve the schema on port 5000.
	server := &Server{
		Client: client,
	}

	schema := server.schema()
	introspection.AddIntrospectionToSchema(schema)

	// Expose GraphQL POST endpoint.
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "5000"
	}
	http.Handle("/graphql", graphql.HTTPHandler(schema))
	http.ListenAndServe(":"+port, nil)
}

func main() {
	serveGraphql()
}
