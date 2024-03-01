package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
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
		field.Int64("from_post_id").
			Positive(),
		field.Bool("register").
			Default(false),
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return nil
}
