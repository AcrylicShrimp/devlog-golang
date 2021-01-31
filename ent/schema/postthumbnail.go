package schema

import "github.com/facebook/ent"

// PostThumbnail holds the schema definition for the PostThumbnail entity.
type PostThumbnail struct {
	ent.Schema
}

// Fields of the PostThumbnail.
func (PostThumbnail) Fields() []ent.Field {
	return nil
}

// Edges of the PostThumbnail.
func (PostThumbnail) Edges() []ent.Edge {
	return nil
}
