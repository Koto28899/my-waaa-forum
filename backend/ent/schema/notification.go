package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
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
		field.Enum("notification_type").
			Values("adopt", "reply", "like", "at", "from_admin"),
		field.Int64("from_role_id").
			Positive(),
		field.Enum("secne_type").
			Values("post", "comment", "reply"),
		field.Int64("sence_id").
			Positive(),
		field.Int64("to_role_id").
			Positive(),
		field.Bool("is_checked").
			Default(false),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}
