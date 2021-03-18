package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"time"
)

// PostVideo holds the schema definition for the PostVideo entity.
type PostVideo struct {
	ent.Schema
}

// Fields of the PostVideo.
func (PostVideo) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.Uint64("index"),
		field.String("url").MaxLen(256).Unique(),
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
		index.Fields("index", "post").Unique(),
	}
}
