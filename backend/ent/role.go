// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/role"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// HashedPwd holds the value of the "hashed_pwd" field.
	HashedPwd string `json:"hashed_pwd,omitempty"`
	// PasswordChangedAt holds the value of the "password_changed_at" field.
	PasswordChangedAt time.Time `json:"password_changed_at,omitempty"`
	// RoleName holds the value of the "role_name" field.
	RoleName string `json:"role_name,omitempty"`
	// Type holds the value of the "type" field.
	Type role.Type `json:"type,omitempty"`
	// Statement holds the value of the "statement" field.
	Statement *string `json:"statement,omitempty"`
	// PostsCount holds the value of the "posts_count" field.
	PostsCount int64 `json:"posts_count,omitempty"`
	// CommentsCount holds the value of the "comments_count" field.
	CommentsCount int64 `json:"comments_count,omitempty"`
	// RepliesCount holds the value of the "replies_count" field.
	RepliesCount int64 `json:"replies_count,omitempty"`
	// LikesCount holds the value of the "likes_count" field.
	LikesCount int64 `json:"likes_count,omitempty"`
	// HelpsCount holds the value of the "helps_count" field.
	HelpsCount int64 `json:"helps_count,omitempty"`
	// FansCount holds the value of the "fans_count" field.
	FansCount    int64 `json:"fans_count,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID, role.FieldPostsCount, role.FieldCommentsCount, role.FieldRepliesCount, role.FieldLikesCount, role.FieldHelpsCount, role.FieldFansCount:
			values[i] = new(sql.NullInt64)
		case role.FieldEmail, role.FieldHashedPwd, role.FieldRoleName, role.FieldType, role.FieldStatement:
			values[i] = new(sql.NullString)
		case role.FieldCreatedAt, role.FieldUpdatedAt, role.FieldDeletedAt, role.FieldPasswordChangedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case role.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case role.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case role.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = new(time.Time)
				*r.DeletedAt = value.Time
			}
		case role.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				r.Email = value.String
			}
		case role.FieldHashedPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashed_pwd", values[i])
			} else if value.Valid {
				r.HashedPwd = value.String
			}
		case role.FieldPasswordChangedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field password_changed_at", values[i])
			} else if value.Valid {
				r.PasswordChangedAt = value.Time
			}
		case role.FieldRoleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_name", values[i])
			} else if value.Valid {
				r.RoleName = value.String
			}
		case role.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = role.Type(value.String)
			}
		case role.FieldStatement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field statement", values[i])
			} else if value.Valid {
				r.Statement = new(string)
				*r.Statement = value.String
			}
		case role.FieldPostsCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field posts_count", values[i])
			} else if value.Valid {
				r.PostsCount = value.Int64
			}
		case role.FieldCommentsCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comments_count", values[i])
			} else if value.Valid {
				r.CommentsCount = value.Int64
			}
		case role.FieldRepliesCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field replies_count", values[i])
			} else if value.Valid {
				r.RepliesCount = value.Int64
			}
		case role.FieldLikesCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field likes_count", values[i])
			} else if value.Valid {
				r.LikesCount = value.Int64
			}
		case role.FieldHelpsCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field helps_count", values[i])
			} else if value.Valid {
				r.HelpsCount = value.Int64
			}
		case role.FieldFansCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fans_count", values[i])
			} else if value.Valid {
				r.FansCount = value.Int64
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role.
// This includes values selected through modifiers, order, etc.
func (r *Role) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := r.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(r.Email)
	builder.WriteString(", ")
	builder.WriteString("hashed_pwd=")
	builder.WriteString(r.HashedPwd)
	builder.WriteString(", ")
	builder.WriteString("password_changed_at=")
	builder.WriteString(r.PasswordChangedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("role_name=")
	builder.WriteString(r.RoleName)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", r.Type))
	builder.WriteString(", ")
	if v := r.Statement; v != nil {
		builder.WriteString("statement=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("posts_count=")
	builder.WriteString(fmt.Sprintf("%v", r.PostsCount))
	builder.WriteString(", ")
	builder.WriteString("comments_count=")
	builder.WriteString(fmt.Sprintf("%v", r.CommentsCount))
	builder.WriteString(", ")
	builder.WriteString("replies_count=")
	builder.WriteString(fmt.Sprintf("%v", r.RepliesCount))
	builder.WriteString(", ")
	builder.WriteString("likes_count=")
	builder.WriteString(fmt.Sprintf("%v", r.LikesCount))
	builder.WriteString(", ")
	builder.WriteString("helps_count=")
	builder.WriteString(fmt.Sprintf("%v", r.HelpsCount))
	builder.WriteString(", ")
	builder.WriteString("fans_count=")
	builder.WriteString(fmt.Sprintf("%v", r.FansCount))
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role