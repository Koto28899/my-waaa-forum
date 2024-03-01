package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Star holds the schema definition for the Star entity.
type Star struct {
	ent.Schema
}

// Fields of the Star.
func (Star) Fields() []ent.Field {
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
		field.Enum("secne_type").
			Values("post", "comment", "reply"),
		field.Int64("sence_id").
			Positive(),
	}
}

// Edges of the Star.
func (Star) Edges() []ent.Edge {
	return nil
}
