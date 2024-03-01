// Code generated by ent, DO NOT EDIT.

package help

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the help type in the database.
	Label = "help"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldFromPostID holds the string denoting the from_post_id field in the database.
	FieldFromPostID = "from_post_id"
	// FieldAdoptCommentID holds the string denoting the adopt_comment_id field in the database.
	FieldAdoptCommentID = "adopt_comment_id"
	// Table holds the table name of the help in the database.
	Table = "helps"
)

// Columns holds all SQL columns for help fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldFromPostID,
	FieldAdoptCommentID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UpdateDefaultDeletedAt holds the default value on update for the "deleted_at" field.
	UpdateDefaultDeletedAt func() time.Time
	// FromPostIDValidator is a validator for the "from_post_id" field. It is called by the builders before save.
	FromPostIDValidator func(int64) error
	// AdoptCommentIDValidator is a validator for the "adopt_comment_id" field. It is called by the builders before save.
	AdoptCommentIDValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the Help queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByFromPostID orders the results by the from_post_id field.
func ByFromPostID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromPostID, opts...).ToFunc()
}

// ByAdoptCommentID orders the results by the adopt_comment_id field.
func ByAdoptCommentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdoptCommentID, opts...).ToFunc()
}
