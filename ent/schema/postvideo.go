package schema

import "github.com/facebook/ent"

// PostVideo holds the schema definition for the PostVideo entity.
type PostVideo struct {
	ent.Schema
}

// Fields of the PostVideo.
func (PostVideo) Fields() []ent.Field {
	return nil
}

// Edges of the PostVideo.
func (PostVideo) Edges() []ent.Edge {
	return nil
}
