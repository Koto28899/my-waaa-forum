// Code generated by ent, DO NOT EDIT.

package role

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDeletedAt, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldEmail, v))
}

// HashedPwd applies equality check predicate on the "hashed_pwd" field. It's identical to HashedPwdEQ.
func HashedPwd(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldHashedPwd, v))
}

// PasswordChangedAt applies equality check predicate on the "password_changed_at" field. It's identical to PasswordChangedAtEQ.
func PasswordChangedAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldPasswordChangedAt, v))
}

// RoleName applies equality check predicate on the "role_name" field. It's identical to RoleNameEQ.
func RoleName(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRoleName, v))
}

// Statement applies equality check predicate on the "statement" field. It's identical to StatementEQ.
func Statement(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldStatement, v))
}

// PostsCount applies equality check predicate on the "posts_count" field. It's identical to PostsCountEQ.
func PostsCount(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldPostsCount, v))
}

// CommentsCount applies equality check predicate on the "comments_count" field. It's identical to CommentsCountEQ.
func CommentsCount(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCommentsCount, v))
}

// RepliesCount applies equality check predicate on the "replies_count" field. It's identical to RepliesCountEQ.
func RepliesCount(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRepliesCount, v))
}

// LikesCount applies equality check predicate on the "likes_count" field. It's identical to LikesCountEQ.
func LikesCount(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldLikesCount, v))
}

// HelpsCount applies equality check predicate on the "helps_count" field. It's identical to HelpsCountEQ.
func HelpsCount(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldHelpsCount, v))
}

// FansCount applies equality check predicate on the "fans_count" field. It's identical to FansCountEQ.
func FansCount(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldFansCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Role {
	return predicate.Role(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Role {
	return predicate.Role(sql.FieldNotNull(FieldDeletedAt))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldEmail, v))
}

// HashedPwdEQ applies the EQ predicate on the "hashed_pwd" field.
func HashedPwdEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldHashedPwd, v))
}

// HashedPwdNEQ applies the NEQ predicate on the "hashed_pwd" field.
func HashedPwdNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldHashedPwd, v))
}

// HashedPwdIn applies the In predicate on the "hashed_pwd" field.
func HashedPwdIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldHashedPwd, vs...))
}

// HashedPwdNotIn applies the NotIn predicate on the "hashed_pwd" field.
func HashedPwdNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldHashedPwd, vs...))
}

// HashedPwdGT applies the GT predicate on the "hashed_pwd" field.
func HashedPwdGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldHashedPwd, v))
}

// HashedPwdGTE applies the GTE predicate on the "hashed_pwd" field.
func HashedPwdGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldHashedPwd, v))
}

// HashedPwdLT applies the LT predicate on the "hashed_pwd" field.
func HashedPwdLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldHashedPwd, v))
}

// HashedPwdLTE applies the LTE predicate on the "hashed_pwd" field.
func HashedPwdLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldHashedPwd, v))
}

// HashedPwdContains applies the Contains predicate on the "hashed_pwd" field.
func HashedPwdContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldHashedPwd, v))
}

// HashedPwdHasPrefix applies the HasPrefix predicate on the "hashed_pwd" field.
func HashedPwdHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldHashedPwd, v))
}

// HashedPwdHasSuffix applies the HasSuffix predicate on the "hashed_pwd" field.
func HashedPwdHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldHashedPwd, v))
}

// HashedPwdEqualFold applies the EqualFold predicate on the "hashed_pwd" field.
func HashedPwdEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldHashedPwd, v))
}

// HashedPwdContainsFold applies the ContainsFold predicate on the "hashed_pwd" field.
func HashedPwdContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldHashedPwd, v))
}

// PasswordChangedAtEQ applies the EQ predicate on the "password_changed_at" field.
func PasswordChangedAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldPasswordChangedAt, v))
}

// PasswordChangedAtNEQ applies the NEQ predicate on the "password_changed_at" field.
func PasswordChangedAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldPasswordChangedAt, v))
}

// PasswordChangedAtIn applies the In predicate on the "password_changed_at" field.
func PasswordChangedAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldPasswordChangedAt, vs...))
}

// PasswordChangedAtNotIn applies the NotIn predicate on the "password_changed_at" field.
func PasswordChangedAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldPasswordChangedAt, vs...))
}

// PasswordChangedAtGT applies the GT predicate on the "password_changed_at" field.
func PasswordChangedAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldPasswordChangedAt, v))
}

// PasswordChangedAtGTE applies the GTE predicate on the "password_changed_at" field.
func PasswordChangedAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldPasswordChangedAt, v))
}

// PasswordChangedAtLT applies the LT predicate on the "password_changed_at" field.
func PasswordChangedAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldPasswordChangedAt, v))
}

// PasswordChangedAtLTE applies the LTE predicate on the "password_changed_at" field.
func PasswordChangedAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldPasswordChangedAt, v))
}

// RoleNameEQ applies the EQ predicate on the "role_name" field.
func RoleNameEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRoleName, v))
}

// RoleNameNEQ applies the NEQ predicate on the "role_name" field.
func RoleNameNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldRoleName, v))
}

// RoleNameIn applies the In predicate on the "role_name" field.
func RoleNameIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldRoleName, vs...))
}

// RoleNameNotIn applies the NotIn predicate on the "role_name" field.
func RoleNameNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldRoleName, vs...))
}

// RoleNameGT applies the GT predicate on the "role_name" field.
func RoleNameGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldRoleName, v))
}

// RoleNameGTE applies the GTE predicate on the "role_name" field.
func RoleNameGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldRoleName, v))
}

// RoleNameLT applies the LT predicate on the "role_name" field.
func RoleNameLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldRoleName, v))
}

// RoleNameLTE applies the LTE predicate on the "role_name" field.
func RoleNameLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldRoleName, v))
}

// RoleNameContains applies the Contains predicate on the "role_name" field.
func RoleNameContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldRoleName, v))
}

// RoleNameHasPrefix applies the HasPrefix predicate on the "role_name" field.
func RoleNameHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldRoleName, v))
}

// RoleNameHasSuffix applies the HasSuffix predicate on the "role_name" field.
func RoleNameHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldRoleName, v))
}

// RoleNameEqualFold applies the EqualFold predicate on the "role_name" field.
func RoleNameEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldRoleName, v))
}

// RoleNameContainsFold applies the ContainsFold predicate on the "role_name" field.
func RoleNameContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldRoleName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldType, vs...))
}

// StatementEQ applies the EQ predicate on the "statement" field.
func StatementEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldStatement, v))
}

// StatementNEQ applies the NEQ predicate on the "statement" field.
func StatementNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldStatement, v))
}

// StatementIn applies the In predicate on the "statement" field.
func StatementIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldStatement, vs...))
}

// StatementNotIn applies the NotIn predicate on the "statement" field.
func StatementNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldStatement, vs...))
}

// StatementGT applies the GT predicate on the "statement" field.
func StatementGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldStatement, v))
}

// StatementGTE applies the GTE predicate on the "statement" field.
func StatementGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldStatement, v))
}

// StatementLT applies the LT predicate on the "statement" field.
func StatementLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldStatement, v))
}

// StatementLTE applies the LTE predicate on the "statement" field.
func StatementLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldStatement, v))
}

// StatementContains applies the Contains predicate on the "statement" field.
func StatementContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldStatement, v))
}

// StatementHasPrefix applies the HasPrefix predicate on the "statement" field.
func StatementHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldStatement, v))
}

// StatementHasSuffix applies the HasSuffix predicate on the "statement" field.
func StatementHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldStatement, v))
}

// StatementIsNil applies the IsNil predicate on the "statement" field.
func StatementIsNil() predicate.Role {
	return predicate.Role(sql.FieldIsNull(FieldStatement))
}

// StatementNotNil applies the NotNil predicate on the "statement" field.
func StatementNotNil() predicate.Role {
	return predicate.Role(sql.FieldNotNull(FieldStatement))
}

// StatementEqualFold applies the EqualFold predicate on the "statement" field.
func StatementEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldStatement, v))
}

// StatementContainsFold applies the ContainsFold predicate on the "statement" field.
func StatementContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldStatement, v))
}

// PostsCountEQ applies the EQ predicate on the "posts_count" field.
func PostsCountEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldPostsCount, v))
}

// PostsCountNEQ applies the NEQ predicate on the "posts_count" field.
func PostsCountNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldPostsCount, v))
}

// PostsCountIn applies the In predicate on the "posts_count" field.
func PostsCountIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldPostsCount, vs...))
}

// PostsCountNotIn applies the NotIn predicate on the "posts_count" field.
func PostsCountNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldPostsCount, vs...))
}

// PostsCountGT applies the GT predicate on the "posts_count" field.
func PostsCountGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldPostsCount, v))
}

// PostsCountGTE applies the GTE predicate on the "posts_count" field.
func PostsCountGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldPostsCount, v))
}

// PostsCountLT applies the LT predicate on the "posts_count" field.
func PostsCountLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldPostsCount, v))
}

// PostsCountLTE applies the LTE predicate on the "posts_count" field.
func PostsCountLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldPostsCount, v))
}

// CommentsCountEQ applies the EQ predicate on the "comments_count" field.
func CommentsCountEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCommentsCount, v))
}

// CommentsCountNEQ applies the NEQ predicate on the "comments_count" field.
func CommentsCountNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldCommentsCount, v))
}

// CommentsCountIn applies the In predicate on the "comments_count" field.
func CommentsCountIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldCommentsCount, vs...))
}

// CommentsCountNotIn applies the NotIn predicate on the "comments_count" field.
func CommentsCountNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldCommentsCount, vs...))
}

// CommentsCountGT applies the GT predicate on the "comments_count" field.
func CommentsCountGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldCommentsCount, v))
}

// CommentsCountGTE applies the GTE predicate on the "comments_count" field.
func CommentsCountGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldCommentsCount, v))
}

// CommentsCountLT applies the LT predicate on the "comments_count" field.
func CommentsCountLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldCommentsCount, v))
}

// CommentsCountLTE applies the LTE predicate on the "comments_count" field.
func CommentsCountLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldCommentsCount, v))
}

// RepliesCountEQ applies the EQ predicate on the "replies_count" field.
func RepliesCountEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRepliesCount, v))
}

// RepliesCountNEQ applies the NEQ predicate on the "replies_count" field.
func RepliesCountNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldRepliesCount, v))
}

// RepliesCountIn applies the In predicate on the "replies_count" field.
func RepliesCountIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldRepliesCount, vs...))
}

// RepliesCountNotIn applies the NotIn predicate on the "replies_count" field.
func RepliesCountNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldRepliesCount, vs...))
}

// RepliesCountGT applies the GT predicate on the "replies_count" field.
func RepliesCountGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldRepliesCount, v))
}

// RepliesCountGTE applies the GTE predicate on the "replies_count" field.
func RepliesCountGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldRepliesCount, v))
}

// RepliesCountLT applies the LT predicate on the "replies_count" field.
func RepliesCountLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldRepliesCount, v))
}

// RepliesCountLTE applies the LTE predicate on the "replies_count" field.
func RepliesCountLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldRepliesCount, v))
}

// LikesCountEQ applies the EQ predicate on the "likes_count" field.
func LikesCountEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldLikesCount, v))
}

// LikesCountNEQ applies the NEQ predicate on the "likes_count" field.
func LikesCountNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldLikesCount, v))
}

// LikesCountIn applies the In predicate on the "likes_count" field.
func LikesCountIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldLikesCount, vs...))
}

// LikesCountNotIn applies the NotIn predicate on the "likes_count" field.
func LikesCountNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldLikesCount, vs...))
}

// LikesCountGT applies the GT predicate on the "likes_count" field.
func LikesCountGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldLikesCount, v))
}

// LikesCountGTE applies the GTE predicate on the "likes_count" field.
func LikesCountGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldLikesCount, v))
}

// LikesCountLT applies the LT predicate on the "likes_count" field.
func LikesCountLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldLikesCount, v))
}

// LikesCountLTE applies the LTE predicate on the "likes_count" field.
func LikesCountLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldLikesCount, v))
}

// HelpsCountEQ applies the EQ predicate on the "helps_count" field.
func HelpsCountEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldHelpsCount, v))
}

// HelpsCountNEQ applies the NEQ predicate on the "helps_count" field.
func HelpsCountNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldHelpsCount, v))
}

// HelpsCountIn applies the In predicate on the "helps_count" field.
func HelpsCountIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldHelpsCount, vs...))
}

// HelpsCountNotIn applies the NotIn predicate on the "helps_count" field.
func HelpsCountNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldHelpsCount, vs...))
}

// HelpsCountGT applies the GT predicate on the "helps_count" field.
func HelpsCountGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldHelpsCount, v))
}

// HelpsCountGTE applies the GTE predicate on the "helps_count" field.
func HelpsCountGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldHelpsCount, v))
}

// HelpsCountLT applies the LT predicate on the "helps_count" field.
func HelpsCountLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldHelpsCount, v))
}

// HelpsCountLTE applies the LTE predicate on the "helps_count" field.
func HelpsCountLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldHelpsCount, v))
}

// FansCountEQ applies the EQ predicate on the "fans_count" field.
func FansCountEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldFansCount, v))
}

// FansCountNEQ applies the NEQ predicate on the "fans_count" field.
func FansCountNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldFansCount, v))
}

// FansCountIn applies the In predicate on the "fans_count" field.
func FansCountIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldFansCount, vs...))
}

// FansCountNotIn applies the NotIn predicate on the "fans_count" field.
func FansCountNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldFansCount, vs...))
}

// FansCountGT applies the GT predicate on the "fans_count" field.
func FansCountGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldFansCount, v))
}

// FansCountGTE applies the GTE predicate on the "fans_count" field.
func FansCountGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldFansCount, v))
}

// FansCountLT applies the LT predicate on the "fans_count" field.
func FansCountLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldFansCount, v))
}

// FansCountLTE applies the LTE predicate on the "fans_count" field.
func FansCountLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldFansCount, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(sql.NotPredicates(p))
}
