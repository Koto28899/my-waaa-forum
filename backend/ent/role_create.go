// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/role"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoleCreate) SetCreatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoleCreate) SetUpdatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *RoleCreate) SetDeletedAt(t time.Time) *RoleCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDeletedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// SetEmail sets the "email" field.
func (rc *RoleCreate) SetEmail(s string) *RoleCreate {
	rc.mutation.SetEmail(s)
	return rc
}

// SetHashedPwd sets the "hashed_pwd" field.
func (rc *RoleCreate) SetHashedPwd(s string) *RoleCreate {
	rc.mutation.SetHashedPwd(s)
	return rc
}

// SetPasswordChangedAt sets the "password_changed_at" field.
func (rc *RoleCreate) SetPasswordChangedAt(t time.Time) *RoleCreate {
	rc.mutation.SetPasswordChangedAt(t)
	return rc
}

// SetNillablePasswordChangedAt sets the "password_changed_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillablePasswordChangedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetPasswordChangedAt(*t)
	}
	return rc
}

// SetRoleName sets the "role_name" field.
func (rc *RoleCreate) SetRoleName(s string) *RoleCreate {
	rc.mutation.SetRoleName(s)
	return rc
}

// SetType sets the "type" field.
func (rc *RoleCreate) SetType(r role.Type) *RoleCreate {
	rc.mutation.SetType(r)
	return rc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rc *RoleCreate) SetNillableType(r *role.Type) *RoleCreate {
	if r != nil {
		rc.SetType(*r)
	}
	return rc
}

// SetStatement sets the "statement" field.
func (rc *RoleCreate) SetStatement(s string) *RoleCreate {
	rc.mutation.SetStatement(s)
	return rc
}

// SetNillableStatement sets the "statement" field if the given value is not nil.
func (rc *RoleCreate) SetNillableStatement(s *string) *RoleCreate {
	if s != nil {
		rc.SetStatement(*s)
	}
	return rc
}

// SetPostsCount sets the "posts_count" field.
func (rc *RoleCreate) SetPostsCount(i int64) *RoleCreate {
	rc.mutation.SetPostsCount(i)
	return rc
}

// SetNillablePostsCount sets the "posts_count" field if the given value is not nil.
func (rc *RoleCreate) SetNillablePostsCount(i *int64) *RoleCreate {
	if i != nil {
		rc.SetPostsCount(*i)
	}
	return rc
}

// SetCommentsCount sets the "comments_count" field.
func (rc *RoleCreate) SetCommentsCount(i int64) *RoleCreate {
	rc.mutation.SetCommentsCount(i)
	return rc
}

// SetNillableCommentsCount sets the "comments_count" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCommentsCount(i *int64) *RoleCreate {
	if i != nil {
		rc.SetCommentsCount(*i)
	}
	return rc
}

// SetRepliesCount sets the "replies_count" field.
func (rc *RoleCreate) SetRepliesCount(i int64) *RoleCreate {
	rc.mutation.SetRepliesCount(i)
	return rc
}

// SetNillableRepliesCount sets the "replies_count" field if the given value is not nil.
func (rc *RoleCreate) SetNillableRepliesCount(i *int64) *RoleCreate {
	if i != nil {
		rc.SetRepliesCount(*i)
	}
	return rc
}

// SetLikesCount sets the "likes_count" field.
func (rc *RoleCreate) SetLikesCount(i int64) *RoleCreate {
	rc.mutation.SetLikesCount(i)
	return rc
}

// SetNillableLikesCount sets the "likes_count" field if the given value is not nil.
func (rc *RoleCreate) SetNillableLikesCount(i *int64) *RoleCreate {
	if i != nil {
		rc.SetLikesCount(*i)
	}
	return rc
}

// SetHelpsCount sets the "helps_count" field.
func (rc *RoleCreate) SetHelpsCount(i int64) *RoleCreate {
	rc.mutation.SetHelpsCount(i)
	return rc
}

// SetNillableHelpsCount sets the "helps_count" field if the given value is not nil.
func (rc *RoleCreate) SetNillableHelpsCount(i *int64) *RoleCreate {
	if i != nil {
		rc.SetHelpsCount(*i)
	}
	return rc
}

// SetFansCount sets the "fans_count" field.
func (rc *RoleCreate) SetFansCount(i int64) *RoleCreate {
	rc.mutation.SetFansCount(i)
	return rc
}

// SetNillableFansCount sets the "fans_count" field if the given value is not nil.
func (rc *RoleCreate) SetNillableFansCount(i *int64) *RoleCreate {
	if i != nil {
		rc.SetFansCount(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(i int64) *RoleCreate {
	rc.mutation.SetID(i)
	return rc
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := role.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := role.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.PasswordChangedAt(); !ok {
		v := role.DefaultPasswordChangedAt()
		rc.mutation.SetPasswordChangedAt(v)
	}
	if _, ok := rc.mutation.GetType(); !ok {
		v := role.DefaultType
		rc.mutation.SetType(v)
	}
	if _, ok := rc.mutation.PostsCount(); !ok {
		v := role.DefaultPostsCount
		rc.mutation.SetPostsCount(v)
	}
	if _, ok := rc.mutation.CommentsCount(); !ok {
		v := role.DefaultCommentsCount
		rc.mutation.SetCommentsCount(v)
	}
	if _, ok := rc.mutation.RepliesCount(); !ok {
		v := role.DefaultRepliesCount
		rc.mutation.SetRepliesCount(v)
	}
	if _, ok := rc.mutation.LikesCount(); !ok {
		v := role.DefaultLikesCount
		rc.mutation.SetLikesCount(v)
	}
	if _, ok := rc.mutation.HelpsCount(); !ok {
		v := role.DefaultHelpsCount
		rc.mutation.SetHelpsCount(v)
	}
	if _, ok := rc.mutation.FansCount(); !ok {
		v := role.DefaultFansCount
		rc.mutation.SetFansCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Role.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Role.updated_at"`)}
	}
	if _, ok := rc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Role.email"`)}
	}
	if v, ok := rc.mutation.Email(); ok {
		if err := role.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Role.email": %w`, err)}
		}
	}
	if _, ok := rc.mutation.HashedPwd(); !ok {
		return &ValidationError{Name: "hashed_pwd", err: errors.New(`ent: missing required field "Role.hashed_pwd"`)}
	}
	if v, ok := rc.mutation.HashedPwd(); ok {
		if err := role.HashedPwdValidator(v); err != nil {
			return &ValidationError{Name: "hashed_pwd", err: fmt.Errorf(`ent: validator failed for field "Role.hashed_pwd": %w`, err)}
		}
	}
	if _, ok := rc.mutation.PasswordChangedAt(); !ok {
		return &ValidationError{Name: "password_changed_at", err: errors.New(`ent: missing required field "Role.password_changed_at"`)}
	}
	if _, ok := rc.mutation.RoleName(); !ok {
		return &ValidationError{Name: "role_name", err: errors.New(`ent: missing required field "Role.role_name"`)}
	}
	if v, ok := rc.mutation.RoleName(); ok {
		if err := role.RoleNameValidator(v); err != nil {
			return &ValidationError{Name: "role_name", err: fmt.Errorf(`ent: validator failed for field "Role.role_name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Role.type"`)}
	}
	if v, ok := rc.mutation.GetType(); ok {
		if err := role.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Role.type": %w`, err)}
		}
	}
	if v, ok := rc.mutation.Statement(); ok {
		if err := role.StatementValidator(v); err != nil {
			return &ValidationError{Name: "statement", err: fmt.Errorf(`ent: validator failed for field "Role.statement": %w`, err)}
		}
	}
	if _, ok := rc.mutation.PostsCount(); !ok {
		return &ValidationError{Name: "posts_count", err: errors.New(`ent: missing required field "Role.posts_count"`)}
	}
	if v, ok := rc.mutation.PostsCount(); ok {
		if err := role.PostsCountValidator(v); err != nil {
			return &ValidationError{Name: "posts_count", err: fmt.Errorf(`ent: validator failed for field "Role.posts_count": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CommentsCount(); !ok {
		return &ValidationError{Name: "comments_count", err: errors.New(`ent: missing required field "Role.comments_count"`)}
	}
	if v, ok := rc.mutation.CommentsCount(); ok {
		if err := role.CommentsCountValidator(v); err != nil {
			return &ValidationError{Name: "comments_count", err: fmt.Errorf(`ent: validator failed for field "Role.comments_count": %w`, err)}
		}
	}
	if _, ok := rc.mutation.RepliesCount(); !ok {
		return &ValidationError{Name: "replies_count", err: errors.New(`ent: missing required field "Role.replies_count"`)}
	}
	if v, ok := rc.mutation.RepliesCount(); ok {
		if err := role.RepliesCountValidator(v); err != nil {
			return &ValidationError{Name: "replies_count", err: fmt.Errorf(`ent: validator failed for field "Role.replies_count": %w`, err)}
		}
	}
	if _, ok := rc.mutation.LikesCount(); !ok {
		return &ValidationError{Name: "likes_count", err: errors.New(`ent: missing required field "Role.likes_count"`)}
	}
	if v, ok := rc.mutation.LikesCount(); ok {
		if err := role.LikesCountValidator(v); err != nil {
			return &ValidationError{Name: "likes_count", err: fmt.Errorf(`ent: validator failed for field "Role.likes_count": %w`, err)}
		}
	}
	if _, ok := rc.mutation.HelpsCount(); !ok {
		return &ValidationError{Name: "helps_count", err: errors.New(`ent: missing required field "Role.helps_count"`)}
	}
	if v, ok := rc.mutation.HelpsCount(); ok {
		if err := role.HelpsCountValidator(v); err != nil {
			return &ValidationError{Name: "helps_count", err: fmt.Errorf(`ent: validator failed for field "Role.helps_count": %w`, err)}
		}
	}
	if _, ok := rc.mutation.FansCount(); !ok {
		return &ValidationError{Name: "fans_count", err: errors.New(`ent: missing required field "Role.fans_count"`)}
	}
	if v, ok := rc.mutation.FansCount(); ok {
		if err := role.FansCountValidator(v); err != nil {
			return &ValidationError{Name: "fans_count", err: fmt.Errorf(`ent: validator failed for field "Role.fans_count": %w`, err)}
		}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := role.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Role.id": %w`, err)}
		}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := rc.mutation.Email(); ok {
		_spec.SetField(role.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := rc.mutation.HashedPwd(); ok {
		_spec.SetField(role.FieldHashedPwd, field.TypeString, value)
		_node.HashedPwd = value
	}
	if value, ok := rc.mutation.PasswordChangedAt(); ok {
		_spec.SetField(role.FieldPasswordChangedAt, field.TypeTime, value)
		_node.PasswordChangedAt = value
	}
	if value, ok := rc.mutation.RoleName(); ok {
		_spec.SetField(role.FieldRoleName, field.TypeString, value)
		_node.RoleName = value
	}
	if value, ok := rc.mutation.GetType(); ok {
		_spec.SetField(role.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := rc.mutation.Statement(); ok {
		_spec.SetField(role.FieldStatement, field.TypeString, value)
		_node.Statement = &value
	}
	if value, ok := rc.mutation.PostsCount(); ok {
		_spec.SetField(role.FieldPostsCount, field.TypeInt64, value)
		_node.PostsCount = value
	}
	if value, ok := rc.mutation.CommentsCount(); ok {
		_spec.SetField(role.FieldCommentsCount, field.TypeInt64, value)
		_node.CommentsCount = value
	}
	if value, ok := rc.mutation.RepliesCount(); ok {
		_spec.SetField(role.FieldRepliesCount, field.TypeInt64, value)
		_node.RepliesCount = value
	}
	if value, ok := rc.mutation.LikesCount(); ok {
		_spec.SetField(role.FieldLikesCount, field.TypeInt64, value)
		_node.LikesCount = value
	}
	if value, ok := rc.mutation.HelpsCount(); ok {
		_spec.SetField(role.FieldHelpsCount, field.TypeInt64, value)
		_node.HelpsCount = value
	}
	if value, ok := rc.mutation.FansCount(); ok {
		_spec.SetField(role.FieldFansCount, field.TypeInt64, value)
		_node.FansCount = value
	}
	return _node, _spec
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	err      error
	builders []*RoleCreate
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
