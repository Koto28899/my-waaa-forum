package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
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
		field.Int64("to_section_id").
			Positive(),
		field.Int64("likes_count").
			NonNegative().
			Default(0),
		field.Bool("is_top").
			Default(false),
		field.Bool("is_highlight").
			Default(false),
		field.String("title").
			MaxLen(127).
			NotEmpty(),
		field.String("body").
			MaxLen(1023).
			NotEmpty(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
