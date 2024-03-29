package main

import (
	"net/http"
	"os"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

// server is our graphql server.
type Server struct {
	Db *Db
}

// schema builds the graphql schema.
func (s *Server) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.RegisterPost(builder)

	s.RegisterQuery(builder)
	s.RegisterMutation(builder)
	return builder.MustBuild()
}

func serve() {
	db := GetDb()

	// Instantiate a server, build a server, and serve the schema on port 5000.
	server := &Server{
		Db: db,
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
	serve()
}
