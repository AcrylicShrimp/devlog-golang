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
		field.Uint64("size"),
		field.String("name").MaxLen(255),
		field.String("mime").MaxLen(64),
		field.String("url").MaxLen(512).Unique(),
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
