// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/task"
)

// LootQuery is the builder for querying Loot entities.
type LootQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Loot
	withRodent *RodentQuery
	withTask   *TaskQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LootQuery builder.
func (lq *LootQuery) Where(ps ...predicate.Loot) *LootQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LootQuery) Limit(limit int) *LootQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LootQuery) Offset(offset int) *LootQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LootQuery) Unique(unique bool) *LootQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LootQuery) Order(o ...OrderFunc) *LootQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryRodent chains the current query on the "rodent" edge.
func (lq *LootQuery) QueryRodent() *RodentQuery {
	query := (&RodentClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loot.Table, loot.FieldID, selector),
			sqlgraph.To(rodent.Table, rodent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, loot.RodentTable, loot.RodentColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTask chains the current query on the "task" edge.
func (lq *LootQuery) QueryTask() *TaskQuery {
	query := (&TaskClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loot.Table, loot.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, loot.TaskTable, loot.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Loot entity from the query.
// Returns a *NotFoundError when no Loot was found.
func (lq *LootQuery) First(ctx context.Context) (*Loot, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loot.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LootQuery) FirstX(ctx context.Context) *Loot {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Loot ID from the query.
// Returns a *NotFoundError when no Loot ID was found.
func (lq *LootQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loot.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LootQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Loot entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Loot entity is found.
// Returns a *NotFoundError when no Loot entities are found.
func (lq *LootQuery) Only(ctx context.Context) (*Loot, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loot.Label}
	default:
		return nil, &NotSingularError{loot.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LootQuery) OnlyX(ctx context.Context) *Loot {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Loot ID in the query.
// Returns a *NotSingularError when more than one Loot ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LootQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loot.Label}
	default:
		err = &NotSingularError{loot.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LootQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Loots.
func (lq *LootQuery) All(ctx context.Context) ([]*Loot, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Loot, *LootQuery]()
	return withInterceptors[[]*Loot](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LootQuery) AllX(ctx context.Context) []*Loot {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Loot IDs.
func (lq *LootQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(loot.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LootQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LootQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LootQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LootQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LootQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LootQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LootQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LootQuery) Clone() *LootQuery {
	if lq == nil {
		return nil
	}
	return &LootQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]OrderFunc{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.Loot{}, lq.predicates...),
		withRodent: lq.withRodent.Clone(),
		withTask:   lq.withTask.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithRodent tells the query-builder to eager-load the nodes that are connected to
// the "rodent" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LootQuery) WithRodent(opts ...func(*RodentQuery)) *LootQuery {
	query := (&RodentClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withRodent = query
	return lq
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LootQuery) WithTask(opts ...func(*TaskQuery)) *LootQuery {
	query := (&TaskClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withTask = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type loot.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Loot.Query().
//		GroupBy(loot.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LootQuery) GroupBy(field string, fields ...string) *LootGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LootGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = loot.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type loot.Type `json:"type,omitempty"`
//	}
//
//	client.Loot.Query().
//		Select(loot.FieldType).
//		Scan(ctx, &v)
func (lq *LootQuery) Select(fields ...string) *LootSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LootSelect{LootQuery: lq}
	sbuild.label = loot.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LootSelect configured with the given aggregations.
func (lq *LootQuery) Aggregate(fns ...AggregateFunc) *LootSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LootQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !loot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LootQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Loot, error) {
	var (
		nodes       = []*Loot{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [2]bool{
			lq.withRodent != nil,
			lq.withTask != nil,
		}
	)
	if lq.withRodent != nil || lq.withTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, loot.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Loot).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Loot{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withRodent; query != nil {
		if err := lq.loadRodent(ctx, query, nodes, nil,
			func(n *Loot, e *Rodent) { n.Edges.Rodent = e }); err != nil {
			return nil, err
		}
	}
	if query := lq.withTask; query != nil {
		if err := lq.loadTask(ctx, query, nodes, nil,
			func(n *Loot, e *Task) { n.Edges.Task = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LootQuery) loadRodent(ctx context.Context, query *RodentQuery, nodes []*Loot, init func(*Loot), assign func(*Loot, *Rodent)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Loot)
	for i := range nodes {
		if nodes[i].rodent_loot == nil {
			continue
		}
		fk := *nodes[i].rodent_loot
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(rodent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rodent_loot" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lq *LootQuery) loadTask(ctx context.Context, query *TaskQuery, nodes []*Loot, init func(*Loot), assign func(*Loot, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Loot)
	for i := range nodes {
		if nodes[i].task_loot == nil {
			continue
		}
		fk := *nodes[i].task_loot
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "task_loot" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LootQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LootQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loot.Table, loot.Columns, sqlgraph.NewFieldSpec(loot.FieldID, field.TypeInt))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loot.FieldID)
		for i := range fields {
			if fields[i] != loot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LootQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(loot.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = loot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LootGroupBy is the group-by builder for Loot entities.
type LootGroupBy struct {
	selector
	build *LootQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LootGroupBy) Aggregate(fns ...AggregateFunc) *LootGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LootGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LootQuery, *LootGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LootGroupBy) sqlScan(ctx context.Context, root *LootQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LootSelect is the builder for selecting fields of Loot entities.
type LootSelect struct {
	*LootQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LootSelect) Aggregate(fns ...AggregateFunc) *LootSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LootSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LootQuery, *LootSelect](ctx, ls.LootQuery, ls, ls.inters, v)
}

func (ls *LootSelect) sqlScan(ctx context.Context, root *LootQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
