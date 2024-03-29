// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/help"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HelpCreate is the builder for creating a Help entity.
type HelpCreate struct {
	config
	mutation *HelpMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hc *HelpCreate) SetCreatedAt(t time.Time) *HelpCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hc *HelpCreate) SetNillableCreatedAt(t *time.Time) *HelpCreate {
	if t != nil {
		hc.SetCreatedAt(*t)
	}
	return hc
}

// SetUpdatedAt sets the "updated_at" field.
func (hc *HelpCreate) SetUpdatedAt(t time.Time) *HelpCreate {
	hc.mutation.SetUpdatedAt(t)
	return hc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hc *HelpCreate) SetNillableUpdatedAt(t *time.Time) *HelpCreate {
	if t != nil {
		hc.SetUpdatedAt(*t)
	}
	return hc
}

// SetDeletedAt sets the "deleted_at" field.
func (hc *HelpCreate) SetDeletedAt(t time.Time) *HelpCreate {
	hc.mutation.SetDeletedAt(t)
	return hc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (hc *HelpCreate) SetNillableDeletedAt(t *time.Time) *HelpCreate {
	if t != nil {
		hc.SetDeletedAt(*t)
	}
	return hc
}

// SetFromPostID sets the "from_post_id" field.
func (hc *HelpCreate) SetFromPostID(i int64) *HelpCreate {
	hc.mutation.SetFromPostID(i)
	return hc
}

// SetAdoptCommentID sets the "adopt_comment_id" field.
func (hc *HelpCreate) SetAdoptCommentID(i int64) *HelpCreate {
	hc.mutation.SetAdoptCommentID(i)
	return hc
}

// SetNillableAdoptCommentID sets the "adopt_comment_id" field if the given value is not nil.
func (hc *HelpCreate) SetNillableAdoptCommentID(i *int64) *HelpCreate {
	if i != nil {
		hc.SetAdoptCommentID(*i)
	}
	return hc
}

// SetID sets the "id" field.
func (hc *HelpCreate) SetID(i int64) *HelpCreate {
	hc.mutation.SetID(i)
	return hc
}

// Mutation returns the HelpMutation object of the builder.
func (hc *HelpCreate) Mutation() *HelpMutation {
	return hc.mutation
}

// Save creates the Help in the database.
func (hc *HelpCreate) Save(ctx context.Context) (*Help, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HelpCreate) SaveX(ctx context.Context) *Help {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HelpCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HelpCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HelpCreate) defaults() {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		v := help.DefaultCreatedAt()
		hc.mutation.SetCreatedAt(v)
	}
	if _, ok := hc.mutation.UpdatedAt(); !ok {
		v := help.DefaultUpdatedAt()
		hc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HelpCreate) check() error {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Help.created_at"`)}
	}
	if _, ok := hc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Help.updated_at"`)}
	}
	if _, ok := hc.mutation.FromPostID(); !ok {
		return &ValidationError{Name: "from_post_id", err: errors.New(`ent: missing required field "Help.from_post_id"`)}
	}
	if v, ok := hc.mutation.FromPostID(); ok {
		if err := help.FromPostIDValidator(v); err != nil {
			return &ValidationError{Name: "from_post_id", err: fmt.Errorf(`ent: validator failed for field "Help.from_post_id": %w`, err)}
		}
	}
	if v, ok := hc.mutation.AdoptCommentID(); ok {
		if err := help.AdoptCommentIDValidator(v); err != nil {
			return &ValidationError{Name: "adopt_comment_id", err: fmt.Errorf(`ent: validator failed for field "Help.adopt_comment_id": %w`, err)}
		}
	}
	if v, ok := hc.mutation.ID(); ok {
		if err := help.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Help.id": %w`, err)}
		}
	}
	return nil
}

func (hc *HelpCreate) sqlSave(ctx context.Context) (*Help, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HelpCreate) createSpec() (*Help, *sqlgraph.CreateSpec) {
	var (
		_node = &Help{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(help.Table, sqlgraph.NewFieldSpec(help.FieldID, field.TypeInt64))
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.SetField(help.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hc.mutation.UpdatedAt(); ok {
		_spec.SetField(help.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hc.mutation.DeletedAt(); ok {
		_spec.SetField(help.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := hc.mutation.FromPostID(); ok {
		_spec.SetField(help.FieldFromPostID, field.TypeInt64, value)
		_node.FromPostID = value
	}
	if value, ok := hc.mutation.AdoptCommentID(); ok {
		_spec.SetField(help.FieldAdoptCommentID, field.TypeInt64, value)
		_node.AdoptCommentID = &value
	}
	return _node, _spec
}

// HelpCreateBulk is the builder for creating many Help entities in bulk.
type HelpCreateBulk struct {
	config
	err      error
	builders []*HelpCreate
}

// Save creates the Help entities in the database.
func (hcb *HelpCreateBulk) Save(ctx context.Context) ([]*Help, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Help, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HelpMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HelpCreateBulk) SaveX(ctx context.Context) []*Help {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HelpCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HelpCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
