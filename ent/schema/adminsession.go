package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// AdminSession holds the schema definition for the AdminSession entity.
type AdminSession struct {
	ent.Schema
}

// Fields of the AdminSession.
func (AdminSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").MaxLen(256).Unique(),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the AdminSession.
func (AdminSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Admin.Type).Ref("sessions").Unique().Required(),
	}
}

// Indices of the AdminSession.
func (AdminSession) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("expires_at"),
	}
}
