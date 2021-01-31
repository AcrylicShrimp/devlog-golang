package schema

import "github.com/facebook/ent"

// PostImage holds the schema definition for the PostImage entity.
type PostImage struct {
	ent.Schema
}

// Fields of the PostImage.
func (PostImage) Fields() []ent.Field {
	return nil
}

// Edges of the PostImage.
func (PostImage) Edges() []ent.Edge {
	return nil
}
