package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"time"
)

// PostAttachment holds the schema definition for the PostAttachment entity.
type PostAttachment struct {
	ent.Schema
}

// Fields of the PostAttachment.
func (PostAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.Uint64("size"),
		field.String("name").MaxLen(255),
		field.String("mime").MaxLen(64),
		field.String("url").MaxLen(512).Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the PostAttachment.
func (PostAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("attachments").Unique().Required(),
	}
}

// Indices of the PostAttachment.
func (PostAttachment) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "post").Unique(),
	}
}
