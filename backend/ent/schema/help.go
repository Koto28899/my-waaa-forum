package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Help holds the schema definition for the Help entity.
type Help struct {
	ent.Schema
}

// Fields of the Help.
func (Help) Fields() []ent.Field {
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
		field.Int64("adopt_comment_id").
			Positive().
			Optional().
			Nillable(),
	}
}

// Edges of the Help.
func (Help) Edges() []ent.Edge {
	return nil
}
