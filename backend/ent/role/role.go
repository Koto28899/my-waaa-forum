// Code generated by ent, DO NOT EDIT.

package role

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldHashedPwd holds the string denoting the hashed_pwd field in the database.
	FieldHashedPwd = "hashed_pwd"
	// FieldPasswordChangedAt holds the string denoting the password_changed_at field in the database.
	FieldPasswordChangedAt = "password_changed_at"
	// FieldRoleName holds the string denoting the role_name field in the database.
	FieldRoleName = "role_name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatement holds the string denoting the statement field in the database.
	FieldStatement = "statement"
	// FieldPostsCount holds the string denoting the posts_count field in the database.
	FieldPostsCount = "posts_count"
	// FieldCommentsCount holds the string denoting the comments_count field in the database.
	FieldCommentsCount = "comments_count"
	// FieldRepliesCount holds the string denoting the replies_count field in the database.
	FieldRepliesCount = "replies_count"
	// FieldLikesCount holds the string denoting the likes_count field in the database.
	FieldLikesCount = "likes_count"
	// FieldHelpsCount holds the string denoting the helps_count field in the database.
	FieldHelpsCount = "helps_count"
	// FieldFansCount holds the string denoting the fans_count field in the database.
	FieldFansCount = "fans_count"
	// Table holds the table name of the role in the database.
	Table = "roles"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEmail,
	FieldHashedPwd,
	FieldPasswordChangedAt,
	FieldRoleName,
	FieldType,
	FieldStatement,
	FieldPostsCount,
	FieldCommentsCount,
	FieldRepliesCount,
	FieldLikesCount,
	FieldHelpsCount,
	FieldFansCount,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// HashedPwdValidator is a validator for the "hashed_pwd" field. It is called by the builders before save.
	HashedPwdValidator func(string) error
	// DefaultPasswordChangedAt holds the default value on creation for the "password_changed_at" field.
	DefaultPasswordChangedAt func() time.Time
	// RoleNameValidator is a validator for the "role_name" field. It is called by the builders before save.
	RoleNameValidator func(string) error
	// StatementValidator is a validator for the "statement" field. It is called by the builders before save.
	StatementValidator func(string) error
	// DefaultPostsCount holds the default value on creation for the "posts_count" field.
	DefaultPostsCount int64
	// PostsCountValidator is a validator for the "posts_count" field. It is called by the builders before save.
	PostsCountValidator func(int64) error
	// DefaultCommentsCount holds the default value on creation for the "comments_count" field.
	DefaultCommentsCount int64
	// CommentsCountValidator is a validator for the "comments_count" field. It is called by the builders before save.
	CommentsCountValidator func(int64) error
	// DefaultRepliesCount holds the default value on creation for the "replies_count" field.
	DefaultRepliesCount int64
	// RepliesCountValidator is a validator for the "replies_count" field. It is called by the builders before save.
	RepliesCountValidator func(int64) error
	// DefaultLikesCount holds the default value on creation for the "likes_count" field.
	DefaultLikesCount int64
	// LikesCountValidator is a validator for the "likes_count" field. It is called by the builders before save.
	LikesCountValidator func(int64) error
	// DefaultHelpsCount holds the default value on creation for the "helps_count" field.
	DefaultHelpsCount int64
	// HelpsCountValidator is a validator for the "helps_count" field. It is called by the builders before save.
	HelpsCountValidator func(int64) error
	// DefaultFansCount holds the default value on creation for the "fans_count" field.
	DefaultFansCount int64
	// FansCountValidator is a validator for the "fans_count" field. It is called by the builders before save.
	FansCountValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// Type defines the type for the "type" enum field.
type Type string

// TypeUser is the default value of the Type enum.
const DefaultType = TypeUser

// Type values.
const (
	TypeUser  Type = "user"
	TypeAdmin Type = "admin"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeUser, TypeAdmin:
		return nil
	default:
		return fmt.Errorf("role: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Role queries.
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

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByHashedPwd orders the results by the hashed_pwd field.
func ByHashedPwd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashedPwd, opts...).ToFunc()
}

// ByPasswordChangedAt orders the results by the password_changed_at field.
func ByPasswordChangedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordChangedAt, opts...).ToFunc()
}

// ByRoleName orders the results by the role_name field.
func ByRoleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStatement orders the results by the statement field.
func ByStatement(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatement, opts...).ToFunc()
}

// ByPostsCount orders the results by the posts_count field.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostsCount, opts...).ToFunc()
}

// ByCommentsCount orders the results by the comments_count field.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentsCount, opts...).ToFunc()
}

// ByRepliesCount orders the results by the replies_count field.
func ByRepliesCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepliesCount, opts...).ToFunc()
}

// ByLikesCount orders the results by the likes_count field.
func ByLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikesCount, opts...).ToFunc()
}

// ByHelpsCount orders the results by the helps_count field.
func ByHelpsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHelpsCount, opts...).ToFunc()
}

// ByFansCount orders the results by the fans_count field.
func ByFansCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFansCount, opts...).ToFunc()
}