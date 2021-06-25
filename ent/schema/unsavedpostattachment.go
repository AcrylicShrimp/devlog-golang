package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// UnsavedPostAttachment holds the schema definition for the UnsavedPostAttachment entity.
type UnsavedPostAttachment struct {
	ent.Schema
}

// Fields of the UnsavedPostAttachment.
func (UnsavedPostAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64),
		field.Enum("validity").Values("pending", "valid", "invalid").Default("pending"),
		field.Uint64("size").Optional().Nillable(),
		field.String("name").MaxLen(255).Optional().Nillable(),
		field.String("mime").MaxLen(64).Optional().Nillable(),
		field.String("url").MaxLen(512).Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UnsavedPostAttachment.
func (UnsavedPostAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unsaved_post", UnsavedPost.Type).Ref("attachments").Unique().Required(),
	}
}

// Indices of the UnsavedPostAttachment.
func (UnsavedPostAttachment) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "unsaved_post").Unique(),
	}
}
