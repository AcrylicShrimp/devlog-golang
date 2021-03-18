package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// PostThumbnail holds the schema definition for the PostThumbnail entity.
type PostThumbnail struct {
	ent.Schema
}

// Fields of the PostThumbnail.
func (PostThumbnail) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("width"),
		field.Uint32("height"),
		field.String("hash").MaxLen(64),
		field.String("url").MaxLen(256).Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the PostThumbnail.
func (PostThumbnail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("thumbnail").Unique().Required(),
	}
}
