package main

import (
	"net/http"
	"os"
	"time"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

// server is our graphql server.
type server struct {
	posts []Post
}

// schema builds the graphql schema.
func (s *server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.RegisterPost(builder)

	s.RegisterQuery(builder)
	s.RegisterMutation(builder)
	return builder.MustBuild()
}

func serveGraphql() {
	// Instantiate a server, build a server, and serve the schema on port 5000.
	server := &server{
		posts: []Post{
			{Title: "first post!", Body: "I was here first!", CreatedAt: time.Now()},
			{Title: "graphql", Body: "did you hear about Thunder?", CreatedAt: time.Now()},
		},
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
