package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// VoteOption holds the schema definition for the VoteOption entity.
type VoteOption struct {
	ent.Schema
}

func (VoteOption) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "vote_options"},
	}
}

// Fields of the VoteOption.
func (VoteOption) Fields() []ent.Field {
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
		field.Int64("vote_id").
			Positive(),
		field.String("info").
			MaxLen(127).
			NotEmpty(),
		field.Int64("count").
			NonNegative().
			Default(0),
	}
}

// Edges of the VoteOption.
func (VoteOption) Edges() []ent.Edge {
	return nil
}
