// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/star"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Star is the model entity for the Star schema.
type Star struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// FromRoleID holds the value of the "from_role_id" field.
	FromRoleID int64 `json:"from_role_id,omitempty"`
	// SecneType holds the value of the "secne_type" field.
	SecneType star.SecneType `json:"secne_type,omitempty"`
	// SenceID holds the value of the "sence_id" field.
	SenceID      int64 `json:"sence_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Star) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case star.FieldID, star.FieldFromRoleID, star.FieldSenceID:
			values[i] = new(sql.NullInt64)
		case star.FieldSecneType:
			values[i] = new(sql.NullString)
		case star.FieldCreatedAt, star.FieldUpdatedAt, star.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Star fields.
func (s *Star) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case star.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case star.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case star.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case star.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = new(time.Time)
				*s.DeletedAt = value.Time
			}
		case star.FieldFromRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_role_id", values[i])
			} else if value.Valid {
				s.FromRoleID = value.Int64
			}
		case star.FieldSecneType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secne_type", values[i])
			} else if value.Valid {
				s.SecneType = star.SecneType(value.String)
			}
		case star.FieldSenceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sence_id", values[i])
			} else if value.Valid {
				s.SenceID = value.Int64
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Star.
// This includes values selected through modifiers, order, etc.
func (s *Star) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Star.
// Note that you need to call Star.Unwrap() before calling this method if this Star
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Star) Update() *StarUpdateOne {
	return NewStarClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Star entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Star) Unwrap() *Star {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Star is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Star) String() string {
	var builder strings.Builder
	builder.WriteString("Star(")
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
	builder.WriteString("from_role_id=")
	builder.WriteString(fmt.Sprintf("%v", s.FromRoleID))
	builder.WriteString(", ")
	builder.WriteString("secne_type=")
	builder.WriteString(fmt.Sprintf("%v", s.SecneType))
	builder.WriteString(", ")
	builder.WriteString("sence_id=")
	builder.WriteString(fmt.Sprintf("%v", s.SenceID))
	builder.WriteByte(')')
	return builder.String()
}

// Stars is a parsable slice of Star.
type Stars []*Star
