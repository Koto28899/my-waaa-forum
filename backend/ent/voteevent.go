// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/voteevent"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VoteEvent is the model entity for the VoteEvent schema.
type VoteEvent struct {
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
	// ToVoteID holds the value of the "to_vote_id" field.
	ToVoteID int64 `json:"to_vote_id,omitempty"`
	// ToVoteOption holds the value of the "to_vote_option" field.
	ToVoteOption int64 `json:"to_vote_option,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VoteEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case voteevent.FieldID, voteevent.FieldFromRoleID, voteevent.FieldToVoteID, voteevent.FieldToVoteOption:
			values[i] = new(sql.NullInt64)
		case voteevent.FieldCreatedAt, voteevent.FieldUpdatedAt, voteevent.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VoteEvent fields.
func (ve *VoteEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case voteevent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ve.ID = int64(value.Int64)
		case voteevent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ve.CreatedAt = value.Time
			}
		case voteevent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ve.UpdatedAt = value.Time
			}
		case voteevent.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ve.DeletedAt = new(time.Time)
				*ve.DeletedAt = value.Time
			}
		case voteevent.FieldFromRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_role_id", values[i])
			} else if value.Valid {
				ve.FromRoleID = value.Int64
			}
		case voteevent.FieldToVoteID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_vote_id", values[i])
			} else if value.Valid {
				ve.ToVoteID = value.Int64
			}
		case voteevent.FieldToVoteOption:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_vote_option", values[i])
			} else if value.Valid {
				ve.ToVoteOption = value.Int64
			}
		default:
			ve.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VoteEvent.
// This includes values selected through modifiers, order, etc.
func (ve *VoteEvent) Value(name string) (ent.Value, error) {
	return ve.selectValues.Get(name)
}

// Update returns a builder for updating this VoteEvent.
// Note that you need to call VoteEvent.Unwrap() before calling this method if this VoteEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (ve *VoteEvent) Update() *VoteEventUpdateOne {
	return NewVoteEventClient(ve.config).UpdateOne(ve)
}

// Unwrap unwraps the VoteEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ve *VoteEvent) Unwrap() *VoteEvent {
	_tx, ok := ve.config.driver.(*txDriver)
	if !ok {
		panic("ent: VoteEvent is not a transactional entity")
	}
	ve.config.driver = _tx.drv
	return ve
}

// String implements the fmt.Stringer.
func (ve *VoteEvent) String() string {
	var builder strings.Builder
	builder.WriteString("VoteEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ve.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ve.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ve.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ve.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("from_role_id=")
	builder.WriteString(fmt.Sprintf("%v", ve.FromRoleID))
	builder.WriteString(", ")
	builder.WriteString("to_vote_id=")
	builder.WriteString(fmt.Sprintf("%v", ve.ToVoteID))
	builder.WriteString(", ")
	builder.WriteString("to_vote_option=")
	builder.WriteString(fmt.Sprintf("%v", ve.ToVoteOption))
	builder.WriteByte(')')
	return builder.String()
}

// VoteEvents is a parsable slice of VoteEvent.
type VoteEvents []*VoteEvent
