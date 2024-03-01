package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
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
		field.Int64("to_post_id").
			Positive(),
		field.Int64("likes_count").
			NonNegative().
			Default(0),
		field.Bool("is_top").
			Default(false),
		field.String("body").
			NotEmpty().
			MaxLen(1023),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
