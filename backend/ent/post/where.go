// Code generated by ent, DO NOT EDIT.

package post

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDeletedAt, v))
}

// FromRoleID applies equality check predicate on the "from_role_id" field. It's identical to FromRoleIDEQ.
func FromRoleID(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldFromRoleID, v))
}

// ToSectionID applies equality check predicate on the "to_section_id" field. It's identical to ToSectionIDEQ.
func ToSectionID(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldToSectionID, v))
}

// LikesCount applies equality check predicate on the "likes_count" field. It's identical to LikesCountEQ.
func LikesCount(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLikesCount, v))
}

// IsTop applies equality check predicate on the "is_top" field. It's identical to IsTopEQ.
func IsTop(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsTop, v))
}

// IsHighlight applies equality check predicate on the "is_highlight" field. It's identical to IsHighlightEQ.
func IsHighlight(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsHighlight, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldDeletedAt))
}

// FromRoleIDEQ applies the EQ predicate on the "from_role_id" field.
func FromRoleIDEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldFromRoleID, v))
}

// FromRoleIDNEQ applies the NEQ predicate on the "from_role_id" field.
func FromRoleIDNEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldFromRoleID, v))
}

// FromRoleIDIn applies the In predicate on the "from_role_id" field.
func FromRoleIDIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldFromRoleID, vs...))
}

// FromRoleIDNotIn applies the NotIn predicate on the "from_role_id" field.
func FromRoleIDNotIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldFromRoleID, vs...))
}

// FromRoleIDGT applies the GT predicate on the "from_role_id" field.
func FromRoleIDGT(v int64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldFromRoleID, v))
}

// FromRoleIDGTE applies the GTE predicate on the "from_role_id" field.
func FromRoleIDGTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldFromRoleID, v))
}

// FromRoleIDLT applies the LT predicate on the "from_role_id" field.
func FromRoleIDLT(v int64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldFromRoleID, v))
}

// FromRoleIDLTE applies the LTE predicate on the "from_role_id" field.
func FromRoleIDLTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldFromRoleID, v))
}

// ToSectionIDEQ applies the EQ predicate on the "to_section_id" field.
func ToSectionIDEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldToSectionID, v))
}

// ToSectionIDNEQ applies the NEQ predicate on the "to_section_id" field.
func ToSectionIDNEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldToSectionID, v))
}

// ToSectionIDIn applies the In predicate on the "to_section_id" field.
func ToSectionIDIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldToSectionID, vs...))
}

// ToSectionIDNotIn applies the NotIn predicate on the "to_section_id" field.
func ToSectionIDNotIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldToSectionID, vs...))
}

// ToSectionIDGT applies the GT predicate on the "to_section_id" field.
func ToSectionIDGT(v int64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldToSectionID, v))
}

// ToSectionIDGTE applies the GTE predicate on the "to_section_id" field.
func ToSectionIDGTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldToSectionID, v))
}

// ToSectionIDLT applies the LT predicate on the "to_section_id" field.
func ToSectionIDLT(v int64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldToSectionID, v))
}

// ToSectionIDLTE applies the LTE predicate on the "to_section_id" field.
func ToSectionIDLTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldToSectionID, v))
}

// LikesCountEQ applies the EQ predicate on the "likes_count" field.
func LikesCountEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLikesCount, v))
}

// LikesCountNEQ applies the NEQ predicate on the "likes_count" field.
func LikesCountNEQ(v int64) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldLikesCount, v))
}

// LikesCountIn applies the In predicate on the "likes_count" field.
func LikesCountIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldLikesCount, vs...))
}

// LikesCountNotIn applies the NotIn predicate on the "likes_count" field.
func LikesCountNotIn(vs ...int64) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldLikesCount, vs...))
}

// LikesCountGT applies the GT predicate on the "likes_count" field.
func LikesCountGT(v int64) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldLikesCount, v))
}

// LikesCountGTE applies the GTE predicate on the "likes_count" field.
func LikesCountGTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldLikesCount, v))
}

// LikesCountLT applies the LT predicate on the "likes_count" field.
func LikesCountLT(v int64) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldLikesCount, v))
}

// LikesCountLTE applies the LTE predicate on the "likes_count" field.
func LikesCountLTE(v int64) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldLikesCount, v))
}

// IsTopEQ applies the EQ predicate on the "is_top" field.
func IsTopEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsTop, v))
}

// IsTopNEQ applies the NEQ predicate on the "is_top" field.
func IsTopNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsTop, v))
}

// IsHighlightEQ applies the EQ predicate on the "is_highlight" field.
func IsHighlightEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsHighlight, v))
}

// IsHighlightNEQ applies the NEQ predicate on the "is_highlight" field.
func IsHighlightNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsHighlight, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTitle, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldBody, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}
