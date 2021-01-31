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
		field.String("email").Unique().MaxLen(128),
		field.String("username").Unique().MaxLen(64),
		field.String("pw").MaxLen(60),
		field.Time("joined_at").Default(time.Now),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sessions", AdminSession.Type).Ref("user").Unique(),
	}
}
