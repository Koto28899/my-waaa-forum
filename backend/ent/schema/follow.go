package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Follow holds the schema definition for the Follow entity.
type Follow struct {
	ent.Schema
}

// Fields of the Follow.
func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			UpdateDefault(time.Now).
			Nillable(),
		field.Int64("from_role_id").
			Positive(),
		field.Int64("to_role_id").
			Positive(),
	}
}

// Edges of the Follow.
func (Follow) Edges() []ent.Edge {
	return nil
}
