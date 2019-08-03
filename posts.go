package main

import (
	"context"
	"time"

	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/samsarahq/thunder/reactive"
)

type Post struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

var posts = []Post{
	{Title: "first post!", Body: "I was here first!", CreatedAt: time.Now()},
	{Title: "graphql test", Body: "did you hear about Thunder?", CreatedAt: time.Now()},
}

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
		return posts
	})
}

// registerMutation registers the root mutation type.
func (s *Server) RegisterMutation(schema *schemabuilder.Schema) {
	obj := schema.Mutation()

	obj.FieldFunc("echo", func(args struct{ Message string }) string {
		return args.Message
	})
}
