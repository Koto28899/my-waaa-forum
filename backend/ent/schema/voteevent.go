package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// VoteEvent holds the schema definition for the VoteEvent entity.
type VoteEvent struct {
	ent.Schema
}

func (VoteEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "vote_events"},
	}
}

// Fields of the VoteEvent.
func (VoteEvent) Fields() []ent.Field {
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
		field.Int64("to_vote_id").
			Positive(),
		field.Int64("to_vote_option").
			Positive(),
	}
}

// Edges of the VoteEvent.
func (VoteEvent) Edges() []ent.Edge {
	return nil
}
