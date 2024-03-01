// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/help"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HelpUpdate is the builder for updating Help entities.
type HelpUpdate struct {
	config
	hooks    []Hook
	mutation *HelpMutation
}

// Where appends a list predicates to the HelpUpdate builder.
func (hu *HelpUpdate) Where(ps ...predicate.Help) *HelpUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HelpUpdate) SetUpdatedAt(t time.Time) *HelpUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// SetDeletedAt sets the "deleted_at" field.
func (hu *HelpUpdate) SetDeletedAt(t time.Time) *HelpUpdate {
	hu.mutation.SetDeletedAt(t)
	return hu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (hu *HelpUpdate) ClearDeletedAt() *HelpUpdate {
	hu.mutation.ClearDeletedAt()
	return hu
}

// SetFromPostID sets the "from_post_id" field.
func (hu *HelpUpdate) SetFromPostID(i int64) *HelpUpdate {
	hu.mutation.ResetFromPostID()
	hu.mutation.SetFromPostID(i)
	return hu
}

// SetNillableFromPostID sets the "from_post_id" field if the given value is not nil.
func (hu *HelpUpdate) SetNillableFromPostID(i *int64) *HelpUpdate {
	if i != nil {
		hu.SetFromPostID(*i)
	}
	return hu
}

// AddFromPostID adds i to the "from_post_id" field.
func (hu *HelpUpdate) AddFromPostID(i int64) *HelpUpdate {
	hu.mutation.AddFromPostID(i)
	return hu
}

// SetAdoptCommentID sets the "adopt_comment_id" field.
func (hu *HelpUpdate) SetAdoptCommentID(i int64) *HelpUpdate {
	hu.mutation.ResetAdoptCommentID()
	hu.mutation.SetAdoptCommentID(i)
	return hu
}

// SetNillableAdoptCommentID sets the "adopt_comment_id" field if the given value is not nil.
func (hu *HelpUpdate) SetNillableAdoptCommentID(i *int64) *HelpUpdate {
	if i != nil {
		hu.SetAdoptCommentID(*i)
	}
	return hu
}

// AddAdoptCommentID adds i to the "adopt_comment_id" field.
func (hu *HelpUpdate) AddAdoptCommentID(i int64) *HelpUpdate {
	hu.mutation.AddAdoptCommentID(i)
	return hu
}

// ClearAdoptCommentID clears the value of the "adopt_comment_id" field.
func (hu *HelpUpdate) ClearAdoptCommentID() *HelpUpdate {
	hu.mutation.ClearAdoptCommentID()
	return hu
}

// Mutation returns the HelpMutation object of the builder.
func (hu *HelpUpdate) Mutation() *HelpMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HelpUpdate) Save(ctx context.Context) (int, error) {
	hu.defaults()
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HelpUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HelpUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HelpUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HelpUpdate) defaults() {
	if _, ok := hu.mutation.UpdatedAt(); !ok {
		v := help.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
	if _, ok := hu.mutation.DeletedAt(); !ok && !hu.mutation.DeletedAtCleared() {
		v := help.UpdateDefaultDeletedAt()
		hu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HelpUpdate) check() error {
	if v, ok := hu.mutation.FromPostID(); ok {
		if err := help.FromPostIDValidator(v); err != nil {
			return &ValidationError{Name: "from_post_id", err: fmt.Errorf(`ent: validator failed for field "Help.from_post_id": %w`, err)}
		}
	}
	if v, ok := hu.mutation.AdoptCommentID(); ok {
		if err := help.AdoptCommentIDValidator(v); err != nil {
			return &ValidationError{Name: "adopt_comment_id", err: fmt.Errorf(`ent: validator failed for field "Help.adopt_comment_id": %w`, err)}
		}
	}
	return nil
}

func (hu *HelpUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(help.Table, help.Columns, sqlgraph.NewFieldSpec(help.FieldID, field.TypeInt64))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.SetField(help.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hu.mutation.DeletedAt(); ok {
		_spec.SetField(help.FieldDeletedAt, field.TypeTime, value)
	}
	if hu.mutation.DeletedAtCleared() {
		_spec.ClearField(help.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := hu.mutation.FromPostID(); ok {
		_spec.SetField(help.FieldFromPostID, field.TypeInt64, value)
	}
	if value, ok := hu.mutation.AddedFromPostID(); ok {
		_spec.AddField(help.FieldFromPostID, field.TypeInt64, value)
	}
	if value, ok := hu.mutation.AdoptCommentID(); ok {
		_spec.SetField(help.FieldAdoptCommentID, field.TypeInt64, value)
	}
	if value, ok := hu.mutation.AddedAdoptCommentID(); ok {
		_spec.AddField(help.FieldAdoptCommentID, field.TypeInt64, value)
	}
	if hu.mutation.AdoptCommentIDCleared() {
		_spec.ClearField(help.FieldAdoptCommentID, field.TypeInt64)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{help.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HelpUpdateOne is the builder for updating a single Help entity.
type HelpUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HelpMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HelpUpdateOne) SetUpdatedAt(t time.Time) *HelpUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// SetDeletedAt sets the "deleted_at" field.
func (huo *HelpUpdateOne) SetDeletedAt(t time.Time) *HelpUpdateOne {
	huo.mutation.SetDeletedAt(t)
	return huo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (huo *HelpUpdateOne) ClearDeletedAt() *HelpUpdateOne {
	huo.mutation.ClearDeletedAt()
	return huo
}

// SetFromPostID sets the "from_post_id" field.
func (huo *HelpUpdateOne) SetFromPostID(i int64) *HelpUpdateOne {
	huo.mutation.ResetFromPostID()
	huo.mutation.SetFromPostID(i)
	return huo
}

// SetNillableFromPostID sets the "from_post_id" field if the given value is not nil.
func (huo *HelpUpdateOne) SetNillableFromPostID(i *int64) *HelpUpdateOne {
	if i != nil {
		huo.SetFromPostID(*i)
	}
	return huo
}

// AddFromPostID adds i to the "from_post_id" field.
func (huo *HelpUpdateOne) AddFromPostID(i int64) *HelpUpdateOne {
	huo.mutation.AddFromPostID(i)
	return huo
}

// SetAdoptCommentID sets the "adopt_comment_id" field.
func (huo *HelpUpdateOne) SetAdoptCommentID(i int64) *HelpUpdateOne {
	huo.mutation.ResetAdoptCommentID()
	huo.mutation.SetAdoptCommentID(i)
	return huo
}

// SetNillableAdoptCommentID sets the "adopt_comment_id" field if the given value is not nil.
func (huo *HelpUpdateOne) SetNillableAdoptCommentID(i *int64) *HelpUpdateOne {
	if i != nil {
		huo.SetAdoptCommentID(*i)
	}
	return huo
}

// AddAdoptCommentID adds i to the "adopt_comment_id" field.
func (huo *HelpUpdateOne) AddAdoptCommentID(i int64) *HelpUpdateOne {
	huo.mutation.AddAdoptCommentID(i)
	return huo
}

// ClearAdoptCommentID clears the value of the "adopt_comment_id" field.
func (huo *HelpUpdateOne) ClearAdoptCommentID() *HelpUpdateOne {
	huo.mutation.ClearAdoptCommentID()
	return huo
}

// Mutation returns the HelpMutation object of the builder.
func (huo *HelpUpdateOne) Mutation() *HelpMutation {
	return huo.mutation
}

// Where appends a list predicates to the HelpUpdate builder.
func (huo *HelpUpdateOne) Where(ps ...predicate.Help) *HelpUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HelpUpdateOne) Select(field string, fields ...string) *HelpUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Help entity.
func (huo *HelpUpdateOne) Save(ctx context.Context) (*Help, error) {
	huo.defaults()
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HelpUpdateOne) SaveX(ctx context.Context) *Help {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HelpUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HelpUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HelpUpdateOne) defaults() {
	if _, ok := huo.mutation.UpdatedAt(); !ok {
		v := help.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
	if _, ok := huo.mutation.DeletedAt(); !ok && !huo.mutation.DeletedAtCleared() {
		v := help.UpdateDefaultDeletedAt()
		huo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HelpUpdateOne) check() error {
	if v, ok := huo.mutation.FromPostID(); ok {
		if err := help.FromPostIDValidator(v); err != nil {
			return &ValidationError{Name: "from_post_id", err: fmt.Errorf(`ent: validator failed for field "Help.from_post_id": %w`, err)}
		}
	}
	if v, ok := huo.mutation.AdoptCommentID(); ok {
		if err := help.AdoptCommentIDValidator(v); err != nil {
			return &ValidationError{Name: "adopt_comment_id", err: fmt.Errorf(`ent: validator failed for field "Help.adopt_comment_id": %w`, err)}
		}
	}
	return nil
}

func (huo *HelpUpdateOne) sqlSave(ctx context.Context) (_node *Help, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(help.Table, help.Columns, sqlgraph.NewFieldSpec(help.FieldID, field.TypeInt64))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Help.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, help.FieldID)
		for _, f := range fields {
			if !help.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != help.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.SetField(help.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := huo.mutation.DeletedAt(); ok {
		_spec.SetField(help.FieldDeletedAt, field.TypeTime, value)
	}
	if huo.mutation.DeletedAtCleared() {
		_spec.ClearField(help.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := huo.mutation.FromPostID(); ok {
		_spec.SetField(help.FieldFromPostID, field.TypeInt64, value)
	}
	if value, ok := huo.mutation.AddedFromPostID(); ok {
		_spec.AddField(help.FieldFromPostID, field.TypeInt64, value)
	}
	if value, ok := huo.mutation.AdoptCommentID(); ok {
		_spec.SetField(help.FieldAdoptCommentID, field.TypeInt64, value)
	}
	if value, ok := huo.mutation.AddedAdoptCommentID(); ok {
		_spec.AddField(help.FieldAdoptCommentID, field.TypeInt64, value)
	}
	if huo.mutation.AdoptCommentIDCleared() {
		_spec.ClearField(help.FieldAdoptCommentID, field.TypeInt64)
	}
	_node = &Help{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{help.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
