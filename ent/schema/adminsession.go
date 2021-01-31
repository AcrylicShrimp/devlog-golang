package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// AdminSession holds the schema definition for the AdminSession entity.
type AdminSession struct {
	ent.Schema
}

// Fields of the AdminSession.
func (AdminSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique().MaxLen(256),
		field.Time("used_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the AdminSession.
func (AdminSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", Admin.Type).Required(),
	}
}
