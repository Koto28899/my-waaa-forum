// Code generated by ent, DO NOT EDIT.

package voteoption

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldDeletedAt, v))
}

// VoteID applies equality check predicate on the "vote_id" field. It's identical to VoteIDEQ.
func VoteID(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldVoteID, v))
}

// Info applies equality check predicate on the "info" field. It's identical to InfoEQ.
func Info(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldInfo, v))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotNull(FieldDeletedAt))
}

// VoteIDEQ applies the EQ predicate on the "vote_id" field.
func VoteIDEQ(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldVoteID, v))
}

// VoteIDNEQ applies the NEQ predicate on the "vote_id" field.
func VoteIDNEQ(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldVoteID, v))
}

// VoteIDIn applies the In predicate on the "vote_id" field.
func VoteIDIn(vs ...int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldVoteID, vs...))
}

// VoteIDNotIn applies the NotIn predicate on the "vote_id" field.
func VoteIDNotIn(vs ...int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldVoteID, vs...))
}

// VoteIDGT applies the GT predicate on the "vote_id" field.
func VoteIDGT(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldVoteID, v))
}

// VoteIDGTE applies the GTE predicate on the "vote_id" field.
func VoteIDGTE(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldVoteID, v))
}

// VoteIDLT applies the LT predicate on the "vote_id" field.
func VoteIDLT(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldVoteID, v))
}

// VoteIDLTE applies the LTE predicate on the "vote_id" field.
func VoteIDLTE(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldVoteID, v))
}

// InfoEQ applies the EQ predicate on the "info" field.
func InfoEQ(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldInfo, v))
}

// InfoNEQ applies the NEQ predicate on the "info" field.
func InfoNEQ(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldInfo, v))
}

// InfoIn applies the In predicate on the "info" field.
func InfoIn(vs ...string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldInfo, vs...))
}

// InfoNotIn applies the NotIn predicate on the "info" field.
func InfoNotIn(vs ...string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldInfo, vs...))
}

// InfoGT applies the GT predicate on the "info" field.
func InfoGT(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldInfo, v))
}

// InfoGTE applies the GTE predicate on the "info" field.
func InfoGTE(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldInfo, v))
}

// InfoLT applies the LT predicate on the "info" field.
func InfoLT(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldInfo, v))
}

// InfoLTE applies the LTE predicate on the "info" field.
func InfoLTE(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldInfo, v))
}

// InfoContains applies the Contains predicate on the "info" field.
func InfoContains(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldContains(FieldInfo, v))
}

// InfoHasPrefix applies the HasPrefix predicate on the "info" field.
func InfoHasPrefix(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldHasPrefix(FieldInfo, v))
}

// InfoHasSuffix applies the HasSuffix predicate on the "info" field.
func InfoHasSuffix(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldHasSuffix(FieldInfo, v))
}

// InfoEqualFold applies the EqualFold predicate on the "info" field.
func InfoEqualFold(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEqualFold(FieldInfo, v))
}

// InfoContainsFold applies the ContainsFold predicate on the "info" field.
func InfoContainsFold(v string) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldContainsFold(FieldInfo, v))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int64) predicate.VoteOption {
	return predicate.VoteOption(sql.FieldLTE(FieldCount, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VoteOption) predicate.VoteOption {
	return predicate.VoteOption(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VoteOption) predicate.VoteOption {
	return predicate.VoteOption(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VoteOption) predicate.VoteOption {
	return predicate.VoteOption(sql.NotPredicates(p))
}
