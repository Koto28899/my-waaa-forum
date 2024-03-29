// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldUserAgent holds the string denoting the user_agent field in the database.
	FieldUserAgent = "user_agent"
	// FieldClientIP holds the string denoting the client_ip field in the database.
	FieldClientIP = "client_ip"
	// FieldIsBlocked holds the string denoting the is_blocked field in the database.
	FieldIsBlocked = "is_blocked"
	// Table holds the table name of the session in the database.
	Table = "sessions"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldRoleID,
	FieldCreatedAt,
	FieldExpiresAt,
	FieldUserAgent,
	FieldClientIP,
	FieldIsBlocked,
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
	// RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	RoleIDValidator func(int64) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultCreatedAt holds the default value on update for the "created_at" field.
	UpdateDefaultCreatedAt func() time.Time
	// DefaultExpiresAt holds the default value on creation for the "expires_at" field.
	DefaultExpiresAt func() time.Time
	// UserAgentValidator is a validator for the "user_agent" field. It is called by the builders before save.
	UserAgentValidator func(string) error
	// ClientIPValidator is a validator for the "client_ip" field. It is called by the builders before save.
	ClientIPValidator func(string) error
	// DefaultIsBlocked holds the default value on creation for the "is_blocked" field.
	DefaultIsBlocked bool
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByUserAgent orders the results by the user_agent field.
func ByUserAgent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserAgent, opts...).ToFunc()
}

// ByClientIP orders the results by the client_ip field.
func ByClientIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientIP, opts...).ToFunc()
}

// ByIsBlocked orders the results by the is_blocked field.
func ByIsBlocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsBlocked, opts...).ToFunc()
}
