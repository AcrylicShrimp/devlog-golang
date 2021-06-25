package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// UnsavedPostImage holds the schema definition for the UnsavedPostImage entity.
type UnsavedPostImage struct {
	ent.Schema
}

// Fields of the UnsavedPostImage.
func (UnsavedPostImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64),
		field.Enum("validity").Values("pending", "valid", "invalid").Default("pending"),
		field.Uint32("width").Optional().Nillable(),
		field.Uint32("height").Optional().Nillable(),
		field.String("hash").MaxLen(64).Optional().Nillable(),
		field.String("title").MaxLen(255).Optional().Nillable(),
		field.String("url").MaxLen(512).Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UnsavedPostImage.
func (UnsavedPostImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unsaved_post", UnsavedPost.Type).Ref("images").Unique().Required(),
	}
}

// Indices of the UnsavedPostImage.
func (UnsavedPostImage) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "unsaved_post").Unique(),
	}
}
