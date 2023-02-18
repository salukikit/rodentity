// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/project"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/router"
)

// RouterQuery is the builder for querying Router entities.
type RouterQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.Router
	withRodents *RodentQuery
	withProject *ProjectQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RouterQuery builder.
func (rq *RouterQuery) Where(ps ...predicate.Router) *RouterQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RouterQuery) Limit(limit int) *RouterQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RouterQuery) Offset(offset int) *RouterQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RouterQuery) Unique(unique bool) *RouterQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RouterQuery) Order(o ...OrderFunc) *RouterQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryRodents chains the current query on the "rodents" edge.
func (rq *RouterQuery) QueryRodents() *RodentQuery {
	query := (&RodentClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(router.Table, router.FieldID, selector),
			sqlgraph.To(rodent.Table, rodent.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, router.RodentsTable, router.RodentsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProject chains the current query on the "project" edge.
func (rq *RouterQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(router.Table, router.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, router.ProjectTable, router.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Router entity from the query.
// Returns a *NotFoundError when no Router was found.
func (rq *RouterQuery) First(ctx context.Context) (*Router, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{router.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RouterQuery) FirstX(ctx context.Context) *Router {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Router ID from the query.
// Returns a *NotFoundError when no Router ID was found.
func (rq *RouterQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{router.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RouterQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Router entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Router entity is found.
// Returns a *NotFoundError when no Router entities are found.
func (rq *RouterQuery) Only(ctx context.Context) (*Router, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{router.Label}
	default:
		return nil, &NotSingularError{router.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RouterQuery) OnlyX(ctx context.Context) *Router {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Router ID in the query.
// Returns a *NotSingularError when more than one Router ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RouterQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{router.Label}
	default:
		err = &NotSingularError{router.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RouterQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Routers.
func (rq *RouterQuery) All(ctx context.Context) ([]*Router, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Router, *RouterQuery]()
	return withInterceptors[[]*Router](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RouterQuery) AllX(ctx context.Context) []*Router {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Router IDs.
func (rq *RouterQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(router.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RouterQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RouterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RouterQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RouterQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RouterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RouterQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RouterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RouterQuery) Clone() *RouterQuery {
	if rq == nil {
		return nil
	}
	return &RouterQuery{
		config:      rq.config,
		ctx:         rq.ctx.Clone(),
		order:       append([]OrderFunc{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Router{}, rq.predicates...),
		withRodents: rq.withRodents.Clone(),
		withProject: rq.withProject.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithRodents tells the query-builder to eager-load the nodes that are connected to
// the "rodents" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RouterQuery) WithRodents(opts ...func(*RodentQuery)) *RouterQuery {
	query := (&RodentClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withRodents = query
	return rq
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RouterQuery) WithProject(opts ...func(*ProjectQuery)) *RouterQuery {
	query := (&ProjectClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withProject = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rname string `json:"rname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Router.Query().
//		GroupBy(router.FieldRname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RouterQuery) GroupBy(field string, fields ...string) *RouterGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RouterGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = router.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rname string `json:"rname,omitempty"`
//	}
//
//	client.Router.Query().
//		Select(router.FieldRname).
//		Scan(ctx, &v)
func (rq *RouterQuery) Select(fields ...string) *RouterSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RouterSelect{RouterQuery: rq}
	sbuild.label = router.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RouterSelect configured with the given aggregations.
func (rq *RouterQuery) Aggregate(fns ...AggregateFunc) *RouterSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RouterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !router.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RouterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Router, error) {
	var (
		nodes       = []*Router{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withRodents != nil,
			rq.withProject != nil,
		}
	)
	if rq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, router.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Router).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Router{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withRodents; query != nil {
		if err := rq.loadRodents(ctx, query, nodes,
			func(n *Router) { n.Edges.Rodents = []*Rodent{} },
			func(n *Router, e *Rodent) { n.Edges.Rodents = append(n.Edges.Rodents, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withProject; query != nil {
		if err := rq.loadProject(ctx, query, nodes, nil,
			func(n *Router, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RouterQuery) loadRodents(ctx context.Context, query *RodentQuery, nodes []*Router, init func(*Router), assign func(*Router, *Rodent)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Router)
	nids := make(map[xid.ID]map[*Router]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(router.RodentsTable)
		s.Join(joinT).On(s.C(rodent.FieldID), joinT.C(router.RodentsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(router.RodentsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(router.RodentsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Router]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Rodent](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "rodents" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RouterQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Router, init func(*Router), assign func(*Router, *Project)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Router)
	for i := range nodes {
		if nodes[i].project_routers == nil {
			continue
		}
		fk := *nodes[i].project_routers
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_routers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RouterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RouterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(router.Table, router.Columns, sqlgraph.NewFieldSpec(router.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, router.FieldID)
		for i := range fields {
			if fields[i] != router.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RouterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(router.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = router.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RouterGroupBy is the group-by builder for Router entities.
type RouterGroupBy struct {
	selector
	build *RouterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RouterGroupBy) Aggregate(fns ...AggregateFunc) *RouterGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RouterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouterQuery, *RouterGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RouterGroupBy) sqlScan(ctx context.Context, root *RouterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RouterSelect is the builder for selecting fields of Router entities.
type RouterSelect struct {
	*RouterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RouterSelect) Aggregate(fns ...AggregateFunc) *RouterSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RouterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouterQuery, *RouterSelect](ctx, rs.RouterQuery, rs, rs.inters, v)
}

func (rs *RouterSelect) sqlScan(ctx context.Context, root *RouterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
