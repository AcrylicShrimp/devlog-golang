package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(32).Unique(),
		field.String("description").MaxLen(255).Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.SetNull,
		}),
		edge.To("unsaved_posts", UnsavedPost.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.SetNull,
		}),
	}
}
