package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// PostImage holds the schema definition for the PostImage entity.
type PostImage struct {
	ent.Schema
}

// Fields of the PostImage.
func (PostImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64),
		field.Uint32("width"),
		field.Uint32("height"),
		field.String("hash").MaxLen(64),
		field.String("title").MaxLen(255),
		field.String("url").MaxLen(512).Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the PostImage.
func (PostImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("images").Unique().Required(),
	}
}

// Indices of the PostImage.
func (PostImage) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "post").Unique(),
		index.Fields("created_at", "uuid"),
	}
}
