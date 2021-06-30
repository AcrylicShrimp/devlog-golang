package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// UnsavedPostVideo holds the schema definition for the UnsavedPostVideo entity.
type UnsavedPostVideo struct {
	ent.Schema
}

// Fields of the UnsavedPostVideo.
func (UnsavedPostVideo) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64),
		field.Enum("validity").Values("pending", "valid", "invalid").Default("pending"),
		field.String("title").MaxLen(255).Optional().Nillable(),
		field.String("url").MaxLen(512).Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UnsavedPostVideo.
func (UnsavedPostVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unsaved_post", UnsavedPost.Type).Ref("videos").Unique().Required(),
	}
}

// Indices of the UnsavedPostVideo.
func (UnsavedPostVideo) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "unsaved_post").Unique(),
		index.Fields("created_at", "uuid"),
	}
}
