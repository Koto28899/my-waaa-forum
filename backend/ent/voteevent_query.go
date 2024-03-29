// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/voteevent"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoteEventQuery is the builder for querying VoteEvent entities.
type VoteEventQuery struct {
	config
	ctx        *QueryContext
	order      []voteevent.OrderOption
	inters     []Interceptor
	predicates []predicate.VoteEvent
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoteEventQuery builder.
func (veq *VoteEventQuery) Where(ps ...predicate.VoteEvent) *VoteEventQuery {
	veq.predicates = append(veq.predicates, ps...)
	return veq
}

// Limit the number of records to be returned by this query.
func (veq *VoteEventQuery) Limit(limit int) *VoteEventQuery {
	veq.ctx.Limit = &limit
	return veq
}

// Offset to start from.
func (veq *VoteEventQuery) Offset(offset int) *VoteEventQuery {
	veq.ctx.Offset = &offset
	return veq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (veq *VoteEventQuery) Unique(unique bool) *VoteEventQuery {
	veq.ctx.Unique = &unique
	return veq
}

// Order specifies how the records should be ordered.
func (veq *VoteEventQuery) Order(o ...voteevent.OrderOption) *VoteEventQuery {
	veq.order = append(veq.order, o...)
	return veq
}

// First returns the first VoteEvent entity from the query.
// Returns a *NotFoundError when no VoteEvent was found.
func (veq *VoteEventQuery) First(ctx context.Context) (*VoteEvent, error) {
	nodes, err := veq.Limit(1).All(setContextOp(ctx, veq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voteevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (veq *VoteEventQuery) FirstX(ctx context.Context) *VoteEvent {
	node, err := veq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VoteEvent ID from the query.
// Returns a *NotFoundError when no VoteEvent ID was found.
func (veq *VoteEventQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = veq.Limit(1).IDs(setContextOp(ctx, veq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voteevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (veq *VoteEventQuery) FirstIDX(ctx context.Context) int64 {
	id, err := veq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VoteEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VoteEvent entity is found.
// Returns a *NotFoundError when no VoteEvent entities are found.
func (veq *VoteEventQuery) Only(ctx context.Context) (*VoteEvent, error) {
	nodes, err := veq.Limit(2).All(setContextOp(ctx, veq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voteevent.Label}
	default:
		return nil, &NotSingularError{voteevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (veq *VoteEventQuery) OnlyX(ctx context.Context) *VoteEvent {
	node, err := veq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VoteEvent ID in the query.
// Returns a *NotSingularError when more than one VoteEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (veq *VoteEventQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = veq.Limit(2).IDs(setContextOp(ctx, veq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voteevent.Label}
	default:
		err = &NotSingularError{voteevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (veq *VoteEventQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := veq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VoteEvents.
func (veq *VoteEventQuery) All(ctx context.Context) ([]*VoteEvent, error) {
	ctx = setContextOp(ctx, veq.ctx, "All")
	if err := veq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VoteEvent, *VoteEventQuery]()
	return withInterceptors[[]*VoteEvent](ctx, veq, qr, veq.inters)
}

// AllX is like All, but panics if an error occurs.
func (veq *VoteEventQuery) AllX(ctx context.Context) []*VoteEvent {
	nodes, err := veq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VoteEvent IDs.
func (veq *VoteEventQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if veq.ctx.Unique == nil && veq.path != nil {
		veq.Unique(true)
	}
	ctx = setContextOp(ctx, veq.ctx, "IDs")
	if err = veq.Select(voteevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (veq *VoteEventQuery) IDsX(ctx context.Context) []int64 {
	ids, err := veq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (veq *VoteEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, veq.ctx, "Count")
	if err := veq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, veq, querierCount[*VoteEventQuery](), veq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (veq *VoteEventQuery) CountX(ctx context.Context) int {
	count, err := veq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (veq *VoteEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, veq.ctx, "Exist")
	switch _, err := veq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (veq *VoteEventQuery) ExistX(ctx context.Context) bool {
	exist, err := veq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoteEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (veq *VoteEventQuery) Clone() *VoteEventQuery {
	if veq == nil {
		return nil
	}
	return &VoteEventQuery{
		config:     veq.config,
		ctx:        veq.ctx.Clone(),
		order:      append([]voteevent.OrderOption{}, veq.order...),
		inters:     append([]Interceptor{}, veq.inters...),
		predicates: append([]predicate.VoteEvent{}, veq.predicates...),
		// clone intermediate query.
		sql:  veq.sql.Clone(),
		path: veq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VoteEvent.Query().
//		GroupBy(voteevent.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (veq *VoteEventQuery) GroupBy(field string, fields ...string) *VoteEventGroupBy {
	veq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VoteEventGroupBy{build: veq}
	grbuild.flds = &veq.ctx.Fields
	grbuild.label = voteevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.VoteEvent.Query().
//		Select(voteevent.FieldCreatedAt).
//		Scan(ctx, &v)
func (veq *VoteEventQuery) Select(fields ...string) *VoteEventSelect {
	veq.ctx.Fields = append(veq.ctx.Fields, fields...)
	sbuild := &VoteEventSelect{VoteEventQuery: veq}
	sbuild.label = voteevent.Label
	sbuild.flds, sbuild.scan = &veq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VoteEventSelect configured with the given aggregations.
func (veq *VoteEventQuery) Aggregate(fns ...AggregateFunc) *VoteEventSelect {
	return veq.Select().Aggregate(fns...)
}

func (veq *VoteEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range veq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, veq); err != nil {
				return err
			}
		}
	}
	for _, f := range veq.ctx.Fields {
		if !voteevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if veq.path != nil {
		prev, err := veq.path(ctx)
		if err != nil {
			return err
		}
		veq.sql = prev
	}
	return nil
}

func (veq *VoteEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VoteEvent, error) {
	var (
		nodes = []*VoteEvent{}
		_spec = veq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VoteEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VoteEvent{config: veq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(veq.modifiers) > 0 {
		_spec.Modifiers = veq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, veq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (veq *VoteEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := veq.querySpec()
	if len(veq.modifiers) > 0 {
		_spec.Modifiers = veq.modifiers
	}
	_spec.Node.Columns = veq.ctx.Fields
	if len(veq.ctx.Fields) > 0 {
		_spec.Unique = veq.ctx.Unique != nil && *veq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, veq.driver, _spec)
}

func (veq *VoteEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(voteevent.Table, voteevent.Columns, sqlgraph.NewFieldSpec(voteevent.FieldID, field.TypeInt64))
	_spec.From = veq.sql
	if unique := veq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if veq.path != nil {
		_spec.Unique = true
	}
	if fields := veq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voteevent.FieldID)
		for i := range fields {
			if fields[i] != voteevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := veq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := veq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := veq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := veq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (veq *VoteEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(veq.driver.Dialect())
	t1 := builder.Table(voteevent.Table)
	columns := veq.ctx.Fields
	if len(columns) == 0 {
		columns = voteevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if veq.sql != nil {
		selector = veq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if veq.ctx.Unique != nil && *veq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range veq.modifiers {
		m(selector)
	}
	for _, p := range veq.predicates {
		p(selector)
	}
	for _, p := range veq.order {
		p(selector)
	}
	if offset := veq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := veq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (veq *VoteEventQuery) ForUpdate(opts ...sql.LockOption) *VoteEventQuery {
	if veq.driver.Dialect() == dialect.Postgres {
		veq.Unique(false)
	}
	veq.modifiers = append(veq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return veq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (veq *VoteEventQuery) ForShare(opts ...sql.LockOption) *VoteEventQuery {
	if veq.driver.Dialect() == dialect.Postgres {
		veq.Unique(false)
	}
	veq.modifiers = append(veq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return veq
}

// VoteEventGroupBy is the group-by builder for VoteEvent entities.
type VoteEventGroupBy struct {
	selector
	build *VoteEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vegb *VoteEventGroupBy) Aggregate(fns ...AggregateFunc) *VoteEventGroupBy {
	vegb.fns = append(vegb.fns, fns...)
	return vegb
}

// Scan applies the selector query and scans the result into the given value.
func (vegb *VoteEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vegb.build.ctx, "GroupBy")
	if err := vegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoteEventQuery, *VoteEventGroupBy](ctx, vegb.build, vegb, vegb.build.inters, v)
}

func (vegb *VoteEventGroupBy) sqlScan(ctx context.Context, root *VoteEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vegb.fns))
	for _, fn := range vegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vegb.flds)+len(vegb.fns))
		for _, f := range *vegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VoteEventSelect is the builder for selecting fields of VoteEvent entities.
type VoteEventSelect struct {
	*VoteEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ves *VoteEventSelect) Aggregate(fns ...AggregateFunc) *VoteEventSelect {
	ves.fns = append(ves.fns, fns...)
	return ves
}

// Scan applies the selector query and scans the result into the given value.
func (ves *VoteEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ves.ctx, "Select")
	if err := ves.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoteEventQuery, *VoteEventSelect](ctx, ves.VoteEventQuery, ves, ves.inters, v)
}

func (ves *VoteEventSelect) sqlScan(ctx context.Context, root *VoteEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ves.fns))
	for _, fn := range ves.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ves.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ves.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
