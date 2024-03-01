// Code generated by ent, DO NOT EDIT.

package follow

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Follow {
	return predicate.Follow(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Follow {
	return predicate.Follow(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Follow {
	return predicate.Follow(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldDeletedAt, v))
}

// FromRoleID applies equality check predicate on the "from_role_id" field. It's identical to FromRoleIDEQ.
func FromRoleID(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldFromRoleID, v))
}

// ToRoleID applies equality check predicate on the "to_role_id" field. It's identical to ToRoleIDEQ.
func ToRoleID(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldToRoleID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Follow {
	return predicate.Follow(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Follow {
	return predicate.Follow(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Follow {
	return predicate.Follow(sql.FieldNotNull(FieldDeletedAt))
}

// FromRoleIDEQ applies the EQ predicate on the "from_role_id" field.
func FromRoleIDEQ(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldFromRoleID, v))
}

// FromRoleIDNEQ applies the NEQ predicate on the "from_role_id" field.
func FromRoleIDNEQ(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldNEQ(FieldFromRoleID, v))
}

// FromRoleIDIn applies the In predicate on the "from_role_id" field.
func FromRoleIDIn(vs ...int64) predicate.Follow {
	return predicate.Follow(sql.FieldIn(FieldFromRoleID, vs...))
}

// FromRoleIDNotIn applies the NotIn predicate on the "from_role_id" field.
func FromRoleIDNotIn(vs ...int64) predicate.Follow {
	return predicate.Follow(sql.FieldNotIn(FieldFromRoleID, vs...))
}

// FromRoleIDGT applies the GT predicate on the "from_role_id" field.
func FromRoleIDGT(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldGT(FieldFromRoleID, v))
}

// FromRoleIDGTE applies the GTE predicate on the "from_role_id" field.
func FromRoleIDGTE(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldGTE(FieldFromRoleID, v))
}

// FromRoleIDLT applies the LT predicate on the "from_role_id" field.
func FromRoleIDLT(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldLT(FieldFromRoleID, v))
}

// FromRoleIDLTE applies the LTE predicate on the "from_role_id" field.
func FromRoleIDLTE(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldLTE(FieldFromRoleID, v))
}

// ToRoleIDEQ applies the EQ predicate on the "to_role_id" field.
func ToRoleIDEQ(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldEQ(FieldToRoleID, v))
}

// ToRoleIDNEQ applies the NEQ predicate on the "to_role_id" field.
func ToRoleIDNEQ(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldNEQ(FieldToRoleID, v))
}

// ToRoleIDIn applies the In predicate on the "to_role_id" field.
func ToRoleIDIn(vs ...int64) predicate.Follow {
	return predicate.Follow(sql.FieldIn(FieldToRoleID, vs...))
}

// ToRoleIDNotIn applies the NotIn predicate on the "to_role_id" field.
func ToRoleIDNotIn(vs ...int64) predicate.Follow {
	return predicate.Follow(sql.FieldNotIn(FieldToRoleID, vs...))
}

// ToRoleIDGT applies the GT predicate on the "to_role_id" field.
func ToRoleIDGT(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldGT(FieldToRoleID, v))
}

// ToRoleIDGTE applies the GTE predicate on the "to_role_id" field.
func ToRoleIDGTE(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldGTE(FieldToRoleID, v))
}

// ToRoleIDLT applies the LT predicate on the "to_role_id" field.
func ToRoleIDLT(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldLT(FieldToRoleID, v))
}

// ToRoleIDLTE applies the LTE predicate on the "to_role_id" field.
func ToRoleIDLTE(v int64) predicate.Follow {
	return predicate.Follow(sql.FieldLTE(FieldToRoleID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Follow) predicate.Follow {
	return predicate.Follow(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Follow) predicate.Follow {
	return predicate.Follow(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Follow) predicate.Follow {
	return predicate.Follow(sql.NotPredicates(p))
}