package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").MaxLen(128).Unique(),
		field.String("username").MaxLen(64).Unique(),
		field.String("password").MaxLen(60),
		field.String("key").MaxLen(64),
		field.Time("joined_at").Default(time.Now),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", AdminSession.Type),
		edge.To("robot_accesses", AdminRobotAccess.Type),
		edge.To("posts", Post.Type),
		edge.To("unsaved_posts", UnsavedPost.Type),
	}
}

// Indices of the Admin.
func (Admin) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("username", "key").Unique(),
	}
}
