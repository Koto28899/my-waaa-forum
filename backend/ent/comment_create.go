// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/comment"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentCreate) SetCreatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommentCreate) SetUpdatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CommentCreate) SetDeletedAt(t time.Time) *CommentCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableDeletedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetFromRoleID sets the "from_role_id" field.
func (cc *CommentCreate) SetFromRoleID(i int64) *CommentCreate {
	cc.mutation.SetFromRoleID(i)
	return cc
}

// SetToPostID sets the "to_post_id" field.
func (cc *CommentCreate) SetToPostID(i int64) *CommentCreate {
	cc.mutation.SetToPostID(i)
	return cc
}

// SetLikesCount sets the "likes_count" field.
func (cc *CommentCreate) SetLikesCount(i int64) *CommentCreate {
	cc.mutation.SetLikesCount(i)
	return cc
}

// SetNillableLikesCount sets the "likes_count" field if the given value is not nil.
func (cc *CommentCreate) SetNillableLikesCount(i *int64) *CommentCreate {
	if i != nil {
		cc.SetLikesCount(*i)
	}
	return cc
}

// SetIsTop sets the "is_top" field.
func (cc *CommentCreate) SetIsTop(b bool) *CommentCreate {
	cc.mutation.SetIsTop(b)
	return cc
}

// SetNillableIsTop sets the "is_top" field if the given value is not nil.
func (cc *CommentCreate) SetNillableIsTop(b *bool) *CommentCreate {
	if b != nil {
		cc.SetIsTop(*b)
	}
	return cc
}

// SetBody sets the "body" field.
func (cc *CommentCreate) SetBody(s string) *CommentCreate {
	cc.mutation.SetBody(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(i int64) *CommentCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := comment.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := comment.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.LikesCount(); !ok {
		v := comment.DefaultLikesCount
		cc.mutation.SetLikesCount(v)
	}
	if _, ok := cc.mutation.IsTop(); !ok {
		v := comment.DefaultIsTop
		cc.mutation.SetIsTop(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Comment.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Comment.updated_at"`)}
	}
	if _, ok := cc.mutation.FromRoleID(); !ok {
		return &ValidationError{Name: "from_role_id", err: errors.New(`ent: missing required field "Comment.from_role_id"`)}
	}
	if v, ok := cc.mutation.FromRoleID(); ok {
		if err := comment.FromRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "from_role_id", err: fmt.Errorf(`ent: validator failed for field "Comment.from_role_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ToPostID(); !ok {
		return &ValidationError{Name: "to_post_id", err: errors.New(`ent: missing required field "Comment.to_post_id"`)}
	}
	if v, ok := cc.mutation.ToPostID(); ok {
		if err := comment.ToPostIDValidator(v); err != nil {
			return &ValidationError{Name: "to_post_id", err: fmt.Errorf(`ent: validator failed for field "Comment.to_post_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.LikesCount(); !ok {
		return &ValidationError{Name: "likes_count", err: errors.New(`ent: missing required field "Comment.likes_count"`)}
	}
	if v, ok := cc.mutation.LikesCount(); ok {
		if err := comment.LikesCountValidator(v); err != nil {
			return &ValidationError{Name: "likes_count", err: fmt.Errorf(`ent: validator failed for field "Comment.likes_count": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IsTop(); !ok {
		return &ValidationError{Name: "is_top", err: errors.New(`ent: missing required field "Comment.is_top"`)}
	}
	if _, ok := cc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Comment.body"`)}
	}
	if v, ok := cc.mutation.Body(); ok {
		if err := comment.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Comment.body": %w`, err)}
		}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := comment.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Comment.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(comment.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := cc.mutation.FromRoleID(); ok {
		_spec.SetField(comment.FieldFromRoleID, field.TypeInt64, value)
		_node.FromRoleID = value
	}
	if value, ok := cc.mutation.ToPostID(); ok {
		_spec.SetField(comment.FieldToPostID, field.TypeInt64, value)
		_node.ToPostID = value
	}
	if value, ok := cc.mutation.LikesCount(); ok {
		_spec.SetField(comment.FieldLikesCount, field.TypeInt64, value)
		_node.LikesCount = value
	}
	if value, ok := cc.mutation.IsTop(); ok {
		_spec.SetField(comment.FieldIsTop, field.TypeBool, value)
		_node.IsTop = value
	}
	if value, ok := cc.mutation.Body(); ok {
		_spec.SetField(comment.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}