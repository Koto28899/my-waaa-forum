// Code generated by ent, DO NOT EDIT.

package section

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldDeletedAt, v))
}

// SectionName applies equality check predicate on the "section_name" field. It's identical to SectionNameEQ.
func SectionName(v string) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldSectionName, v))
}

// Statement applies equality check predicate on the "statement" field. It's identical to StatementEQ.
func Statement(v string) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldStatement, v))
}

// ManagerID applies equality check predicate on the "manager_id" field. It's identical to ManagerIDEQ.
func ManagerID(v int64) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldManagerID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Section {
	return predicate.Section(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Section {
	return predicate.Section(sql.FieldNotNull(FieldDeletedAt))
}

// SectionNameEQ applies the EQ predicate on the "section_name" field.
func SectionNameEQ(v string) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldSectionName, v))
}

// SectionNameNEQ applies the NEQ predicate on the "section_name" field.
func SectionNameNEQ(v string) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldSectionName, v))
}

// SectionNameIn applies the In predicate on the "section_name" field.
func SectionNameIn(vs ...string) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldSectionName, vs...))
}

// SectionNameNotIn applies the NotIn predicate on the "section_name" field.
func SectionNameNotIn(vs ...string) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldSectionName, vs...))
}

// SectionNameGT applies the GT predicate on the "section_name" field.
func SectionNameGT(v string) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldSectionName, v))
}

// SectionNameGTE applies the GTE predicate on the "section_name" field.
func SectionNameGTE(v string) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldSectionName, v))
}

// SectionNameLT applies the LT predicate on the "section_name" field.
func SectionNameLT(v string) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldSectionName, v))
}

// SectionNameLTE applies the LTE predicate on the "section_name" field.
func SectionNameLTE(v string) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldSectionName, v))
}

// SectionNameContains applies the Contains predicate on the "section_name" field.
func SectionNameContains(v string) predicate.Section {
	return predicate.Section(sql.FieldContains(FieldSectionName, v))
}

// SectionNameHasPrefix applies the HasPrefix predicate on the "section_name" field.
func SectionNameHasPrefix(v string) predicate.Section {
	return predicate.Section(sql.FieldHasPrefix(FieldSectionName, v))
}

// SectionNameHasSuffix applies the HasSuffix predicate on the "section_name" field.
func SectionNameHasSuffix(v string) predicate.Section {
	return predicate.Section(sql.FieldHasSuffix(FieldSectionName, v))
}

// SectionNameEqualFold applies the EqualFold predicate on the "section_name" field.
func SectionNameEqualFold(v string) predicate.Section {
	return predicate.Section(sql.FieldEqualFold(FieldSectionName, v))
}

// SectionNameContainsFold applies the ContainsFold predicate on the "section_name" field.
func SectionNameContainsFold(v string) predicate.Section {
	return predicate.Section(sql.FieldContainsFold(FieldSectionName, v))
}

// StatementEQ applies the EQ predicate on the "statement" field.
func StatementEQ(v string) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldStatement, v))
}

// StatementNEQ applies the NEQ predicate on the "statement" field.
func StatementNEQ(v string) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldStatement, v))
}

// StatementIn applies the In predicate on the "statement" field.
func StatementIn(vs ...string) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldStatement, vs...))
}

// StatementNotIn applies the NotIn predicate on the "statement" field.
func StatementNotIn(vs ...string) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldStatement, vs...))
}

// StatementGT applies the GT predicate on the "statement" field.
func StatementGT(v string) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldStatement, v))
}

// StatementGTE applies the GTE predicate on the "statement" field.
func StatementGTE(v string) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldStatement, v))
}

// StatementLT applies the LT predicate on the "statement" field.
func StatementLT(v string) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldStatement, v))
}

// StatementLTE applies the LTE predicate on the "statement" field.
func StatementLTE(v string) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldStatement, v))
}

// StatementContains applies the Contains predicate on the "statement" field.
func StatementContains(v string) predicate.Section {
	return predicate.Section(sql.FieldContains(FieldStatement, v))
}

// StatementHasPrefix applies the HasPrefix predicate on the "statement" field.
func StatementHasPrefix(v string) predicate.Section {
	return predicate.Section(sql.FieldHasPrefix(FieldStatement, v))
}

// StatementHasSuffix applies the HasSuffix predicate on the "statement" field.
func StatementHasSuffix(v string) predicate.Section {
	return predicate.Section(sql.FieldHasSuffix(FieldStatement, v))
}

// StatementIsNil applies the IsNil predicate on the "statement" field.
func StatementIsNil() predicate.Section {
	return predicate.Section(sql.FieldIsNull(FieldStatement))
}

// StatementNotNil applies the NotNil predicate on the "statement" field.
func StatementNotNil() predicate.Section {
	return predicate.Section(sql.FieldNotNull(FieldStatement))
}

// StatementEqualFold applies the EqualFold predicate on the "statement" field.
func StatementEqualFold(v string) predicate.Section {
	return predicate.Section(sql.FieldEqualFold(FieldStatement, v))
}

// StatementContainsFold applies the ContainsFold predicate on the "statement" field.
func StatementContainsFold(v string) predicate.Section {
	return predicate.Section(sql.FieldContainsFold(FieldStatement, v))
}

// ManagerIDEQ applies the EQ predicate on the "manager_id" field.
func ManagerIDEQ(v int64) predicate.Section {
	return predicate.Section(sql.FieldEQ(FieldManagerID, v))
}

// ManagerIDNEQ applies the NEQ predicate on the "manager_id" field.
func ManagerIDNEQ(v int64) predicate.Section {
	return predicate.Section(sql.FieldNEQ(FieldManagerID, v))
}

// ManagerIDIn applies the In predicate on the "manager_id" field.
func ManagerIDIn(vs ...int64) predicate.Section {
	return predicate.Section(sql.FieldIn(FieldManagerID, vs...))
}

// ManagerIDNotIn applies the NotIn predicate on the "manager_id" field.
func ManagerIDNotIn(vs ...int64) predicate.Section {
	return predicate.Section(sql.FieldNotIn(FieldManagerID, vs...))
}

// ManagerIDGT applies the GT predicate on the "manager_id" field.
func ManagerIDGT(v int64) predicate.Section {
	return predicate.Section(sql.FieldGT(FieldManagerID, v))
}

// ManagerIDGTE applies the GTE predicate on the "manager_id" field.
func ManagerIDGTE(v int64) predicate.Section {
	return predicate.Section(sql.FieldGTE(FieldManagerID, v))
}

// ManagerIDLT applies the LT predicate on the "manager_id" field.
func ManagerIDLT(v int64) predicate.Section {
	return predicate.Section(sql.FieldLT(FieldManagerID, v))
}

// ManagerIDLTE applies the LTE predicate on the "manager_id" field.
func ManagerIDLTE(v int64) predicate.Section {
	return predicate.Section(sql.FieldLTE(FieldManagerID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Section) predicate.Section {
	return predicate.Section(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Section) predicate.Section {
	return predicate.Section(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Section) predicate.Section {
	return predicate.Section(sql.NotPredicates(p))
}
