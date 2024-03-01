package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
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
		field.String("email").
			Unique().
			NotEmpty().
			MaxLen(321).
			Match(regexp.MustCompile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`)),
		field.String("hashed_pwd").
			NotEmpty(),
		field.Time("password_changed_at").
			Default(time.Now),
		field.String("role_name").
			MaxLen(21).
			NotEmpty(),
		field.Enum("type").
			Values("user", "admin").
			Default("user"),
		field.String("statement").
			Optional().
			Nillable().
			MaxLen(127),
		field.Int64("posts_count").
			NonNegative().
			Default(0),
		field.Int64("comments_count").
			NonNegative().
			Default(0),
		field.Int64("replies_count").
			NonNegative().
			Default(0),
		field.Int64("likes_count").
			NonNegative().
			Default(0),
		field.Int64("helps_count").
			NonNegative().
			Default(0),
		field.Int64("fans_count").
			NonNegative().
			Default(0),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
