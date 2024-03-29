// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/reply"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReplyUpdate is the builder for updating Reply entities.
type ReplyUpdate struct {
	config
	hooks    []Hook
	mutation *ReplyMutation
}

// Where appends a list predicates to the ReplyUpdate builder.
func (ru *ReplyUpdate) Where(ps ...predicate.Reply) *ReplyUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReplyUpdate) SetUpdatedAt(t time.Time) *ReplyUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *ReplyUpdate) SetDeletedAt(t time.Time) *ReplyUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ru *ReplyUpdate) ClearDeletedAt() *ReplyUpdate {
	ru.mutation.ClearDeletedAt()
	return ru
}

// SetFromRoleID sets the "from_role_id" field.
func (ru *ReplyUpdate) SetFromRoleID(i int64) *ReplyUpdate {
	ru.mutation.ResetFromRoleID()
	ru.mutation.SetFromRoleID(i)
	return ru
}

// SetNillableFromRoleID sets the "from_role_id" field if the given value is not nil.
func (ru *ReplyUpdate) SetNillableFromRoleID(i *int64) *ReplyUpdate {
	if i != nil {
		ru.SetFromRoleID(*i)
	}
	return ru
}

// AddFromRoleID adds i to the "from_role_id" field.
func (ru *ReplyUpdate) AddFromRoleID(i int64) *ReplyUpdate {
	ru.mutation.AddFromRoleID(i)
	return ru
}

// SetToPostID sets the "to_post_id" field.
func (ru *ReplyUpdate) SetToPostID(i int64) *ReplyUpdate {
	ru.mutation.ResetToPostID()
	ru.mutation.SetToPostID(i)
	return ru
}

// SetNillableToPostID sets the "to_post_id" field if the given value is not nil.
func (ru *ReplyUpdate) SetNillableToPostID(i *int64) *ReplyUpdate {
	if i != nil {
		ru.SetToPostID(*i)
	}
	return ru
}

// AddToPostID adds i to the "to_post_id" field.
func (ru *ReplyUpdate) AddToPostID(i int64) *ReplyUpdate {
	ru.mutation.AddToPostID(i)
	return ru
}

// SetLikesCount sets the "likes_count" field.
func (ru *ReplyUpdate) SetLikesCount(i int64) *ReplyUpdate {
	ru.mutation.ResetLikesCount()
	ru.mutation.SetLikesCount(i)
	return ru
}

// SetNillableLikesCount sets the "likes_count" field if the given value is not nil.
func (ru *ReplyUpdate) SetNillableLikesCount(i *int64) *ReplyUpdate {
	if i != nil {
		ru.SetLikesCount(*i)
	}
	return ru
}

// AddLikesCount adds i to the "likes_count" field.
func (ru *ReplyUpdate) AddLikesCount(i int64) *ReplyUpdate {
	ru.mutation.AddLikesCount(i)
	return ru
}

// SetIsTop sets the "is_top" field.
func (ru *ReplyUpdate) SetIsTop(b bool) *ReplyUpdate {
	ru.mutation.SetIsTop(b)
	return ru
}

// SetNillableIsTop sets the "is_top" field if the given value is not nil.
func (ru *ReplyUpdate) SetNillableIsTop(b *bool) *ReplyUpdate {
	if b != nil {
		ru.SetIsTop(*b)
	}
	return ru
}

// SetBody sets the "body" field.
func (ru *ReplyUpdate) SetBody(s string) *ReplyUpdate {
	ru.mutation.SetBody(s)
	return ru
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (ru *ReplyUpdate) SetNillableBody(s *string) *ReplyUpdate {
	if s != nil {
		ru.SetBody(*s)
	}
	return ru
}

// Mutation returns the ReplyMutation object of the builder.
func (ru *ReplyUpdate) Mutation() *ReplyMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReplyUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReplyUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReplyUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReplyUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReplyUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := reply.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	if _, ok := ru.mutation.DeletedAt(); !ok && !ru.mutation.DeletedAtCleared() {
		v := reply.UpdateDefaultDeletedAt()
		ru.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReplyUpdate) check() error {
	if v, ok := ru.mutation.FromRoleID(); ok {
		if err := reply.FromRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "from_role_id", err: fmt.Errorf(`ent: validator failed for field "Reply.from_role_id": %w`, err)}
		}
	}
	if v, ok := ru.mutation.ToPostID(); ok {
		if err := reply.ToPostIDValidator(v); err != nil {
			return &ValidationError{Name: "to_post_id", err: fmt.Errorf(`ent: validator failed for field "Reply.to_post_id": %w`, err)}
		}
	}
	if v, ok := ru.mutation.LikesCount(); ok {
		if err := reply.LikesCountValidator(v); err != nil {
			return &ValidationError{Name: "likes_count", err: fmt.Errorf(`ent: validator failed for field "Reply.likes_count": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Body(); ok {
		if err := reply.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Reply.body": %w`, err)}
		}
	}
	return nil
}

func (ru *ReplyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(reply.Table, reply.Columns, sqlgraph.NewFieldSpec(reply.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(reply.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(reply.FieldDeletedAt, field.TypeTime, value)
	}
	if ru.mutation.DeletedAtCleared() {
		_spec.ClearField(reply.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ru.mutation.FromRoleID(); ok {
		_spec.SetField(reply.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedFromRoleID(); ok {
		_spec.AddField(reply.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.ToPostID(); ok {
		_spec.SetField(reply.FieldToPostID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedToPostID(); ok {
		_spec.AddField(reply.FieldToPostID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.LikesCount(); ok {
		_spec.SetField(reply.FieldLikesCount, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedLikesCount(); ok {
		_spec.AddField(reply.FieldLikesCount, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.IsTop(); ok {
		_spec.SetField(reply.FieldIsTop, field.TypeBool, value)
	}
	if value, ok := ru.mutation.Body(); ok {
		_spec.SetField(reply.FieldBody, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reply.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReplyUpdateOne is the builder for updating a single Reply entity.
type ReplyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReplyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReplyUpdateOne) SetUpdatedAt(t time.Time) *ReplyUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *ReplyUpdateOne) SetDeletedAt(t time.Time) *ReplyUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruo *ReplyUpdateOne) ClearDeletedAt() *ReplyUpdateOne {
	ruo.mutation.ClearDeletedAt()
	return ruo
}

// SetFromRoleID sets the "from_role_id" field.
func (ruo *ReplyUpdateOne) SetFromRoleID(i int64) *ReplyUpdateOne {
	ruo.mutation.ResetFromRoleID()
	ruo.mutation.SetFromRoleID(i)
	return ruo
}

// SetNillableFromRoleID sets the "from_role_id" field if the given value is not nil.
func (ruo *ReplyUpdateOne) SetNillableFromRoleID(i *int64) *ReplyUpdateOne {
	if i != nil {
		ruo.SetFromRoleID(*i)
	}
	return ruo
}

// AddFromRoleID adds i to the "from_role_id" field.
func (ruo *ReplyUpdateOne) AddFromRoleID(i int64) *ReplyUpdateOne {
	ruo.mutation.AddFromRoleID(i)
	return ruo
}

// SetToPostID sets the "to_post_id" field.
func (ruo *ReplyUpdateOne) SetToPostID(i int64) *ReplyUpdateOne {
	ruo.mutation.ResetToPostID()
	ruo.mutation.SetToPostID(i)
	return ruo
}

// SetNillableToPostID sets the "to_post_id" field if the given value is not nil.
func (ruo *ReplyUpdateOne) SetNillableToPostID(i *int64) *ReplyUpdateOne {
	if i != nil {
		ruo.SetToPostID(*i)
	}
	return ruo
}

// AddToPostID adds i to the "to_post_id" field.
func (ruo *ReplyUpdateOne) AddToPostID(i int64) *ReplyUpdateOne {
	ruo.mutation.AddToPostID(i)
	return ruo
}

// SetLikesCount sets the "likes_count" field.
func (ruo *ReplyUpdateOne) SetLikesCount(i int64) *ReplyUpdateOne {
	ruo.mutation.ResetLikesCount()
	ruo.mutation.SetLikesCount(i)
	return ruo
}

// SetNillableLikesCount sets the "likes_count" field if the given value is not nil.
func (ruo *ReplyUpdateOne) SetNillableLikesCount(i *int64) *ReplyUpdateOne {
	if i != nil {
		ruo.SetLikesCount(*i)
	}
	return ruo
}

// AddLikesCount adds i to the "likes_count" field.
func (ruo *ReplyUpdateOne) AddLikesCount(i int64) *ReplyUpdateOne {
	ruo.mutation.AddLikesCount(i)
	return ruo
}

// SetIsTop sets the "is_top" field.
func (ruo *ReplyUpdateOne) SetIsTop(b bool) *ReplyUpdateOne {
	ruo.mutation.SetIsTop(b)
	return ruo
}

// SetNillableIsTop sets the "is_top" field if the given value is not nil.
func (ruo *ReplyUpdateOne) SetNillableIsTop(b *bool) *ReplyUpdateOne {
	if b != nil {
		ruo.SetIsTop(*b)
	}
	return ruo
}

// SetBody sets the "body" field.
func (ruo *ReplyUpdateOne) SetBody(s string) *ReplyUpdateOne {
	ruo.mutation.SetBody(s)
	return ruo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (ruo *ReplyUpdateOne) SetNillableBody(s *string) *ReplyUpdateOne {
	if s != nil {
		ruo.SetBody(*s)
	}
	return ruo
}

// Mutation returns the ReplyMutation object of the builder.
func (ruo *ReplyUpdateOne) Mutation() *ReplyMutation {
	return ruo.mutation
}

// Where appends a list predicates to the ReplyUpdate builder.
func (ruo *ReplyUpdateOne) Where(ps ...predicate.Reply) *ReplyUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReplyUpdateOne) Select(field string, fields ...string) *ReplyUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reply entity.
func (ruo *ReplyUpdateOne) Save(ctx context.Context) (*Reply, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReplyUpdateOne) SaveX(ctx context.Context) *Reply {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReplyUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReplyUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReplyUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := reply.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	if _, ok := ruo.mutation.DeletedAt(); !ok && !ruo.mutation.DeletedAtCleared() {
		v := reply.UpdateDefaultDeletedAt()
		ruo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReplyUpdateOne) check() error {
	if v, ok := ruo.mutation.FromRoleID(); ok {
		if err := reply.FromRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "from_role_id", err: fmt.Errorf(`ent: validator failed for field "Reply.from_role_id": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.ToPostID(); ok {
		if err := reply.ToPostIDValidator(v); err != nil {
			return &ValidationError{Name: "to_post_id", err: fmt.Errorf(`ent: validator failed for field "Reply.to_post_id": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.LikesCount(); ok {
		if err := reply.LikesCountValidator(v); err != nil {
			return &ValidationError{Name: "likes_count", err: fmt.Errorf(`ent: validator failed for field "Reply.likes_count": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Body(); ok {
		if err := reply.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Reply.body": %w`, err)}
		}
	}
	return nil
}

func (ruo *ReplyUpdateOne) sqlSave(ctx context.Context) (_node *Reply, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(reply.Table, reply.Columns, sqlgraph.NewFieldSpec(reply.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reply.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reply.FieldID)
		for _, f := range fields {
			if !reply.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reply.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(reply.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(reply.FieldDeletedAt, field.TypeTime, value)
	}
	if ruo.mutation.DeletedAtCleared() {
		_spec.ClearField(reply.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ruo.mutation.FromRoleID(); ok {
		_spec.SetField(reply.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedFromRoleID(); ok {
		_spec.AddField(reply.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.ToPostID(); ok {
		_spec.SetField(reply.FieldToPostID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedToPostID(); ok {
		_spec.AddField(reply.FieldToPostID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.LikesCount(); ok {
		_spec.SetField(reply.FieldLikesCount, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedLikesCount(); ok {
		_spec.AddField(reply.FieldLikesCount, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.IsTop(); ok {
		_spec.SetField(reply.FieldIsTop, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.Body(); ok {
		_spec.SetField(reply.FieldBody, field.TypeString, value)
	}
	_node = &Reply{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reply.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
