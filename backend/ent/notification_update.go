// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/notification"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotificationUpdate) SetUpdatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetDeletedAt sets the "deleted_at" field.
func (nu *NotificationUpdate) SetDeletedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetDeletedAt(t)
	return nu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nu *NotificationUpdate) ClearDeletedAt() *NotificationUpdate {
	nu.mutation.ClearDeletedAt()
	return nu
}

// SetNotificationType sets the "notification_type" field.
func (nu *NotificationUpdate) SetNotificationType(nt notification.NotificationType) *NotificationUpdate {
	nu.mutation.SetNotificationType(nt)
	return nu
}

// SetNillableNotificationType sets the "notification_type" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableNotificationType(nt *notification.NotificationType) *NotificationUpdate {
	if nt != nil {
		nu.SetNotificationType(*nt)
	}
	return nu
}

// SetFromRoleID sets the "from_role_id" field.
func (nu *NotificationUpdate) SetFromRoleID(i int64) *NotificationUpdate {
	nu.mutation.ResetFromRoleID()
	nu.mutation.SetFromRoleID(i)
	return nu
}

// SetNillableFromRoleID sets the "from_role_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableFromRoleID(i *int64) *NotificationUpdate {
	if i != nil {
		nu.SetFromRoleID(*i)
	}
	return nu
}

// AddFromRoleID adds i to the "from_role_id" field.
func (nu *NotificationUpdate) AddFromRoleID(i int64) *NotificationUpdate {
	nu.mutation.AddFromRoleID(i)
	return nu
}

// SetSecneType sets the "secne_type" field.
func (nu *NotificationUpdate) SetSecneType(nt notification.SecneType) *NotificationUpdate {
	nu.mutation.SetSecneType(nt)
	return nu
}

// SetNillableSecneType sets the "secne_type" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableSecneType(nt *notification.SecneType) *NotificationUpdate {
	if nt != nil {
		nu.SetSecneType(*nt)
	}
	return nu
}

// SetSenceID sets the "sence_id" field.
func (nu *NotificationUpdate) SetSenceID(i int64) *NotificationUpdate {
	nu.mutation.ResetSenceID()
	nu.mutation.SetSenceID(i)
	return nu
}

// SetNillableSenceID sets the "sence_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableSenceID(i *int64) *NotificationUpdate {
	if i != nil {
		nu.SetSenceID(*i)
	}
	return nu
}

// AddSenceID adds i to the "sence_id" field.
func (nu *NotificationUpdate) AddSenceID(i int64) *NotificationUpdate {
	nu.mutation.AddSenceID(i)
	return nu
}

// SetToRoleID sets the "to_role_id" field.
func (nu *NotificationUpdate) SetToRoleID(i int64) *NotificationUpdate {
	nu.mutation.ResetToRoleID()
	nu.mutation.SetToRoleID(i)
	return nu
}

// SetNillableToRoleID sets the "to_role_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableToRoleID(i *int64) *NotificationUpdate {
	if i != nil {
		nu.SetToRoleID(*i)
	}
	return nu
}

// AddToRoleID adds i to the "to_role_id" field.
func (nu *NotificationUpdate) AddToRoleID(i int64) *NotificationUpdate {
	nu.mutation.AddToRoleID(i)
	return nu
}

// SetIsChecked sets the "is_checked" field.
func (nu *NotificationUpdate) SetIsChecked(b bool) *NotificationUpdate {
	nu.mutation.SetIsChecked(b)
	return nu
}

// SetNillableIsChecked sets the "is_checked" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableIsChecked(b *bool) *NotificationUpdate {
	if b != nil {
		nu.SetIsChecked(*b)
	}
	return nu
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NotificationUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
	if _, ok := nu.mutation.DeletedAt(); !ok && !nu.mutation.DeletedAtCleared() {
		v := notification.UpdateDefaultDeletedAt()
		nu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NotificationUpdate) check() error {
	if v, ok := nu.mutation.NotificationType(); ok {
		if err := notification.NotificationTypeValidator(v); err != nil {
			return &ValidationError{Name: "notification_type", err: fmt.Errorf(`ent: validator failed for field "Notification.notification_type": %w`, err)}
		}
	}
	if v, ok := nu.mutation.FromRoleID(); ok {
		if err := notification.FromRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "from_role_id", err: fmt.Errorf(`ent: validator failed for field "Notification.from_role_id": %w`, err)}
		}
	}
	if v, ok := nu.mutation.SecneType(); ok {
		if err := notification.SecneTypeValidator(v); err != nil {
			return &ValidationError{Name: "secne_type", err: fmt.Errorf(`ent: validator failed for field "Notification.secne_type": %w`, err)}
		}
	}
	if v, ok := nu.mutation.SenceID(); ok {
		if err := notification.SenceIDValidator(v); err != nil {
			return &ValidationError{Name: "sence_id", err: fmt.Errorf(`ent: validator failed for field "Notification.sence_id": %w`, err)}
		}
	}
	if v, ok := nu.mutation.ToRoleID(); ok {
		if err := notification.ToRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "to_role_id", err: fmt.Errorf(`ent: validator failed for field "Notification.to_role_id": %w`, err)}
		}
	}
	return nil
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt64))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.DeletedAt(); ok {
		_spec.SetField(notification.FieldDeletedAt, field.TypeTime, value)
	}
	if nu.mutation.DeletedAtCleared() {
		_spec.ClearField(notification.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nu.mutation.NotificationType(); ok {
		_spec.SetField(notification.FieldNotificationType, field.TypeEnum, value)
	}
	if value, ok := nu.mutation.FromRoleID(); ok {
		_spec.SetField(notification.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.AddedFromRoleID(); ok {
		_spec.AddField(notification.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.SecneType(); ok {
		_spec.SetField(notification.FieldSecneType, field.TypeEnum, value)
	}
	if value, ok := nu.mutation.SenceID(); ok {
		_spec.SetField(notification.FieldSenceID, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.AddedSenceID(); ok {
		_spec.AddField(notification.FieldSenceID, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.ToRoleID(); ok {
		_spec.SetField(notification.FieldToRoleID, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.AddedToRoleID(); ok {
		_spec.AddField(notification.FieldToRoleID, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.IsChecked(); ok {
		_spec.SetField(notification.FieldIsChecked, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotificationUpdateOne) SetUpdatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetDeletedAt sets the "deleted_at" field.
func (nuo *NotificationUpdateOne) SetDeletedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetDeletedAt(t)
	return nuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nuo *NotificationUpdateOne) ClearDeletedAt() *NotificationUpdateOne {
	nuo.mutation.ClearDeletedAt()
	return nuo
}

// SetNotificationType sets the "notification_type" field.
func (nuo *NotificationUpdateOne) SetNotificationType(nt notification.NotificationType) *NotificationUpdateOne {
	nuo.mutation.SetNotificationType(nt)
	return nuo
}

// SetNillableNotificationType sets the "notification_type" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableNotificationType(nt *notification.NotificationType) *NotificationUpdateOne {
	if nt != nil {
		nuo.SetNotificationType(*nt)
	}
	return nuo
}

// SetFromRoleID sets the "from_role_id" field.
func (nuo *NotificationUpdateOne) SetFromRoleID(i int64) *NotificationUpdateOne {
	nuo.mutation.ResetFromRoleID()
	nuo.mutation.SetFromRoleID(i)
	return nuo
}

// SetNillableFromRoleID sets the "from_role_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableFromRoleID(i *int64) *NotificationUpdateOne {
	if i != nil {
		nuo.SetFromRoleID(*i)
	}
	return nuo
}

// AddFromRoleID adds i to the "from_role_id" field.
func (nuo *NotificationUpdateOne) AddFromRoleID(i int64) *NotificationUpdateOne {
	nuo.mutation.AddFromRoleID(i)
	return nuo
}

// SetSecneType sets the "secne_type" field.
func (nuo *NotificationUpdateOne) SetSecneType(nt notification.SecneType) *NotificationUpdateOne {
	nuo.mutation.SetSecneType(nt)
	return nuo
}

// SetNillableSecneType sets the "secne_type" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableSecneType(nt *notification.SecneType) *NotificationUpdateOne {
	if nt != nil {
		nuo.SetSecneType(*nt)
	}
	return nuo
}

// SetSenceID sets the "sence_id" field.
func (nuo *NotificationUpdateOne) SetSenceID(i int64) *NotificationUpdateOne {
	nuo.mutation.ResetSenceID()
	nuo.mutation.SetSenceID(i)
	return nuo
}

// SetNillableSenceID sets the "sence_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableSenceID(i *int64) *NotificationUpdateOne {
	if i != nil {
		nuo.SetSenceID(*i)
	}
	return nuo
}

// AddSenceID adds i to the "sence_id" field.
func (nuo *NotificationUpdateOne) AddSenceID(i int64) *NotificationUpdateOne {
	nuo.mutation.AddSenceID(i)
	return nuo
}

// SetToRoleID sets the "to_role_id" field.
func (nuo *NotificationUpdateOne) SetToRoleID(i int64) *NotificationUpdateOne {
	nuo.mutation.ResetToRoleID()
	nuo.mutation.SetToRoleID(i)
	return nuo
}

// SetNillableToRoleID sets the "to_role_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableToRoleID(i *int64) *NotificationUpdateOne {
	if i != nil {
		nuo.SetToRoleID(*i)
	}
	return nuo
}

// AddToRoleID adds i to the "to_role_id" field.
func (nuo *NotificationUpdateOne) AddToRoleID(i int64) *NotificationUpdateOne {
	nuo.mutation.AddToRoleID(i)
	return nuo
}

// SetIsChecked sets the "is_checked" field.
func (nuo *NotificationUpdateOne) SetIsChecked(b bool) *NotificationUpdateOne {
	nuo.mutation.SetIsChecked(b)
	return nuo
}

// SetNillableIsChecked sets the "is_checked" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableIsChecked(b *bool) *NotificationUpdateOne {
	if b != nil {
		nuo.SetIsChecked(*b)
	}
	return nuo
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NotificationUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := nuo.mutation.DeletedAt(); !ok && !nuo.mutation.DeletedAtCleared() {
		v := notification.UpdateDefaultDeletedAt()
		nuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NotificationUpdateOne) check() error {
	if v, ok := nuo.mutation.NotificationType(); ok {
		if err := notification.NotificationTypeValidator(v); err != nil {
			return &ValidationError{Name: "notification_type", err: fmt.Errorf(`ent: validator failed for field "Notification.notification_type": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.FromRoleID(); ok {
		if err := notification.FromRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "from_role_id", err: fmt.Errorf(`ent: validator failed for field "Notification.from_role_id": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.SecneType(); ok {
		if err := notification.SecneTypeValidator(v); err != nil {
			return &ValidationError{Name: "secne_type", err: fmt.Errorf(`ent: validator failed for field "Notification.secne_type": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.SenceID(); ok {
		if err := notification.SenceIDValidator(v); err != nil {
			return &ValidationError{Name: "sence_id", err: fmt.Errorf(`ent: validator failed for field "Notification.sence_id": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.ToRoleID(); ok {
		if err := notification.ToRoleIDValidator(v); err != nil {
			return &ValidationError{Name: "to_role_id", err: fmt.Errorf(`ent: validator failed for field "Notification.to_role_id": %w`, err)}
		}
	}
	return nil
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt64))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.DeletedAt(); ok {
		_spec.SetField(notification.FieldDeletedAt, field.TypeTime, value)
	}
	if nuo.mutation.DeletedAtCleared() {
		_spec.ClearField(notification.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nuo.mutation.NotificationType(); ok {
		_spec.SetField(notification.FieldNotificationType, field.TypeEnum, value)
	}
	if value, ok := nuo.mutation.FromRoleID(); ok {
		_spec.SetField(notification.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.AddedFromRoleID(); ok {
		_spec.AddField(notification.FieldFromRoleID, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.SecneType(); ok {
		_spec.SetField(notification.FieldSecneType, field.TypeEnum, value)
	}
	if value, ok := nuo.mutation.SenceID(); ok {
		_spec.SetField(notification.FieldSenceID, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.AddedSenceID(); ok {
		_spec.AddField(notification.FieldSenceID, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.ToRoleID(); ok {
		_spec.SetField(notification.FieldToRoleID, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.AddedToRoleID(); ok {
		_spec.AddField(notification.FieldToRoleID, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.IsChecked(); ok {
		_spec.SetField(notification.FieldIsChecked, field.TypeBool, value)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
