package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// RobotAccess holds the schema definition for the RobotAccess entity.
type RobotAccess struct {
	ent.Schema
}

// Fields of the AdminSession.
func (RobotAccess) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").MaxLen(256).Unique(),
		field.String("memo").MaxLen(1024).Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Time("expires_at").Optional().Nillable(),
		field.Time("last_access_at").Optional().Nillable(),
	}
}

// Edges of the RobotAccess.
func (RobotAccess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Admin.Type).Ref("robot_accesses").Required(),
	}
}
