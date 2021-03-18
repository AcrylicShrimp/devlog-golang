package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
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
		field.Time("joined_at").Default(time.Now),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", AdminSession.Type),
	}
}
