package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// PostVideo holds the schema definition for the PostVideo entity.
type PostVideo struct {
	ent.Schema
}

// Fields of the PostVideo.
func (PostVideo) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").MaxLen(64),
		field.String("title").MaxLen(255),
		field.String("url").MaxLen(512).Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the PostVideo.
func (PostVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("videos").Unique().Required(),
	}
}

// Indices of the PostVideo.
func (PostVideo) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "post").Unique(),
	}
}
