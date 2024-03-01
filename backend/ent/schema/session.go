package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Int64("role_id").
			Immutable().
			Positive(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("expires_at").
			Default(time.Now),
		field.String("user_agent").
			NotEmpty().
			Immutable(),
		field.String("client_ip").
			NotEmpty().
			Immutable(),
		field.Bool("is_blocked").
			Default(false),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
