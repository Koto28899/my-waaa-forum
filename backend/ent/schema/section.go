package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

func (Section) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "section"},
	}
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
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
		field.String("section_name").
			MaxLen(21).
			NotEmpty().
			Unique(),
		field.String("statement").
			Optional().
			Nillable().
			MaxLen(127),
		field.Int64("manager_id").
			Positive(),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return nil
}
