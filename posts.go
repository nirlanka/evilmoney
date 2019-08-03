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

// registerPost registers the post type.
func (s *server) RegisterPost(schema *schemabuilder.Schema) {
	obj := schema.Object("Post", Post{})

	obj.FieldFunc("age", func(ctx context.Context, p *Post) string {
		reactive.InvalidateAfter(ctx, 5*time.Second)
		return time.Since(p.CreatedAt).String()
	})
}

// registerQuery registers the root query type.
func (s *server) RegisterQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()

	obj.FieldFunc("posts", func() []Post {
		return s.posts
	})
}

// registerMutation registers the root mutation type.
func (s *server) RegisterMutation(schema *schemabuilder.Schema) {
	obj := schema.Mutation()

	obj.FieldFunc("echo", func(args struct{ Message string }) string {
		return args.Message
	})
}
