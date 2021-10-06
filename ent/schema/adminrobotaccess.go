package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// AdminRobotAccess holds the schema definition for the AdminRobotAccess entity.
type AdminRobotAccess struct {
	ent.Schema
}

// Fields of the AdminSession.
func (AdminRobotAccess) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").MaxLen(256).Unique(),
		field.Time("expires_at").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("last_access_at").Optional(),
	}
}

// Edges of the AdminRobotAccess.
func (AdminRobotAccess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Admin.Type).Ref("robot_accesses").Required(),
	}
}

// Indices of the AdminRobotAccess.
func (AdminRobotAccess) Indices() []ent.Index {
	return []ent.Index{
		index.Fields("expires_at"),
	}
}
