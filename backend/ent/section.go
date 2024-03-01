// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/section"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Section is the model entity for the Section schema.
type Section struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// SectionName holds the value of the "section_name" field.
	SectionName string `json:"section_name,omitempty"`
	// Statement holds the value of the "statement" field.
	Statement *string `json:"statement,omitempty"`
	// ManagerID holds the value of the "manager_id" field.
	ManagerID    int64 `json:"manager_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Section) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case section.FieldID, section.FieldManagerID:
			values[i] = new(sql.NullInt64)
		case section.FieldSectionName, section.FieldStatement:
			values[i] = new(sql.NullString)
		case section.FieldCreatedAt, section.FieldUpdatedAt, section.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Section fields.
func (s *Section) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case section.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case section.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case section.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case section.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = new(time.Time)
				*s.DeletedAt = value.Time
			}
		case section.FieldSectionName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field section_name", values[i])
			} else if value.Valid {
				s.SectionName = value.String
			}
		case section.FieldStatement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field statement", values[i])
			} else if value.Valid {
				s.Statement = new(string)
				*s.Statement = value.String
			}
		case section.FieldManagerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field manager_id", values[i])
			} else if value.Valid {
				s.ManagerID = value.Int64
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Section.
// This includes values selected through modifiers, order, etc.
func (s *Section) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Section.
// Note that you need to call Section.Unwrap() before calling this method if this Section
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Section) Update() *SectionUpdateOne {
	return NewSectionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Section entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Section) Unwrap() *Section {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Section is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Section) String() string {
	var builder strings.Builder
	builder.WriteString("Section(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := s.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("section_name=")
	builder.WriteString(s.SectionName)
	builder.WriteString(", ")
	if v := s.Statement; v != nil {
		builder.WriteString("statement=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("manager_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ManagerID))
	builder.WriteByte(')')
	return builder.String()
}

// Sections is a parsable slice of Section.
type Sections []*Section