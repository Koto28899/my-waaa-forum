// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/vote"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Vote is the model entity for the Vote schema.
type Vote struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// FromPostID holds the value of the "from_post_id" field.
	FromPostID int64 `json:"from_post_id,omitempty"`
	// Register holds the value of the "register" field.
	Register     bool `json:"register,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vote.FieldRegister:
			values[i] = new(sql.NullBool)
		case vote.FieldID, vote.FieldFromPostID:
			values[i] = new(sql.NullInt64)
		case vote.FieldCreatedAt, vote.FieldUpdatedAt, vote.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vote fields.
func (v *Vote) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int64(value.Int64)
		case vote.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case vote.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case vote.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				v.DeletedAt = new(time.Time)
				*v.DeletedAt = value.Time
			}
		case vote.FieldFromPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_post_id", values[i])
			} else if value.Valid {
				v.FromPostID = value.Int64
			}
		case vote.FieldRegister:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field register", values[i])
			} else if value.Valid {
				v.Register = value.Bool
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vote.
// This includes values selected through modifiers, order, etc.
func (v *Vote) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// Update returns a builder for updating this Vote.
// Note that you need to call Vote.Unwrap() before calling this method if this Vote
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vote) Update() *VoteUpdateOne {
	return NewVoteClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vote) Unwrap() *Vote {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vote is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vote) String() string {
	var builder strings.Builder
	builder.WriteString("Vote(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := v.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("from_post_id=")
	builder.WriteString(fmt.Sprintf("%v", v.FromPostID))
	builder.WriteString(", ")
	builder.WriteString("register=")
	builder.WriteString(fmt.Sprintf("%v", v.Register))
	builder.WriteByte(')')
	return builder.String()
}

// Votes is a parsable slice of Vote.
type Votes []*Vote
