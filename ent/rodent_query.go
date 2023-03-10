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
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/project"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/router"
	"github.com/salukikit/rodentity/ent/task"
	"github.com/salukikit/rodentity/ent/user"
)

// RodentQuery is the builder for querying Rodent entities.
type RodentQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.Rodent
	withDevice  *DeviceQuery
	withUser    *UserQuery
	withProject *ProjectQuery
	withRouter  *RouterQuery
	withTasks   *TaskQuery
	withLoot    *LootQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RodentQuery builder.
func (rq *RodentQuery) Where(ps ...predicate.Rodent) *RodentQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RodentQuery) Limit(limit int) *RodentQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RodentQuery) Offset(offset int) *RodentQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RodentQuery) Unique(unique bool) *RodentQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RodentQuery) Order(o ...OrderFunc) *RodentQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryDevice chains the current query on the "device" edge.
func (rq *RodentQuery) QueryDevice() *DeviceQuery {
	query := (&DeviceClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rodent.Table, rodent.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rodent.DeviceTable, rodent.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (rq *RodentQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rodent.Table, rodent.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rodent.UserTable, rodent.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProject chains the current query on the "project" edge.
func (rq *RodentQuery) QueryProject() *ProjectQuery {
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
			sqlgraph.From(rodent.Table, rodent.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rodent.ProjectTable, rodent.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRouter chains the current query on the "router" edge.
func (rq *RodentQuery) QueryRouter() *RouterQuery {
	query := (&RouterClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rodent.Table, rodent.FieldID, selector),
			sqlgraph.To(router.Table, router.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, rodent.RouterTable, rodent.RouterPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTasks chains the current query on the "tasks" edge.
func (rq *RodentQuery) QueryTasks() *TaskQuery {
	query := (&TaskClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rodent.Table, rodent.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, rodent.TasksTable, rodent.TasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLoot chains the current query on the "loot" edge.
func (rq *RodentQuery) QueryLoot() *LootQuery {
	query := (&LootClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rodent.Table, rodent.FieldID, selector),
			sqlgraph.To(loot.Table, loot.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, rodent.LootTable, rodent.LootColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Rodent entity from the query.
// Returns a *NotFoundError when no Rodent was found.
func (rq *RodentQuery) First(ctx context.Context) (*Rodent, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rodent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RodentQuery) FirstX(ctx context.Context) *Rodent {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Rodent ID from the query.
// Returns a *NotFoundError when no Rodent ID was found.
func (rq *RodentQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rodent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RodentQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Rodent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Rodent entity is found.
// Returns a *NotFoundError when no Rodent entities are found.
func (rq *RodentQuery) Only(ctx context.Context) (*Rodent, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rodent.Label}
	default:
		return nil, &NotSingularError{rodent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RodentQuery) OnlyX(ctx context.Context) *Rodent {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Rodent ID in the query.
// Returns a *NotSingularError when more than one Rodent ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RodentQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rodent.Label}
	default:
		err = &NotSingularError{rodent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RodentQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Rodents.
func (rq *RodentQuery) All(ctx context.Context) ([]*Rodent, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Rodent, *RodentQuery]()
	return withInterceptors[[]*Rodent](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RodentQuery) AllX(ctx context.Context) []*Rodent {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Rodent IDs.
func (rq *RodentQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(rodent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RodentQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RodentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RodentQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RodentQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RodentQuery) Exist(ctx context.Context) (bool, error) {
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
func (rq *RodentQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RodentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RodentQuery) Clone() *RodentQuery {
	if rq == nil {
		return nil
	}
	return &RodentQuery{
		config:      rq.config,
		ctx:         rq.ctx.Clone(),
		order:       append([]OrderFunc{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Rodent{}, rq.predicates...),
		withDevice:  rq.withDevice.Clone(),
		withUser:    rq.withUser.Clone(),
		withProject: rq.withProject.Clone(),
		withRouter:  rq.withRouter.Clone(),
		withTasks:   rq.withTasks.Clone(),
		withLoot:    rq.withLoot.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RodentQuery) WithDevice(opts ...func(*DeviceQuery)) *RodentQuery {
	query := (&DeviceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withDevice = query
	return rq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RodentQuery) WithUser(opts ...func(*UserQuery)) *RodentQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RodentQuery) WithProject(opts ...func(*ProjectQuery)) *RodentQuery {
	query := (&ProjectClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withProject = query
	return rq
}

// WithRouter tells the query-builder to eager-load the nodes that are connected to
// the "router" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RodentQuery) WithRouter(opts ...func(*RouterQuery)) *RodentQuery {
	query := (&RouterClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withRouter = query
	return rq
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RodentQuery) WithTasks(opts ...func(*TaskQuery)) *RodentQuery {
	query := (&TaskClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withTasks = query
	return rq
}

// WithLoot tells the query-builder to eager-load the nodes that are connected to
// the "loot" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RodentQuery) WithLoot(opts ...func(*LootQuery)) *RodentQuery {
	query := (&LootClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withLoot = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Rodent.Query().
//		GroupBy(rodent.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RodentQuery) GroupBy(field string, fields ...string) *RodentGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RodentGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = rodent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Rodent.Query().
//		Select(rodent.FieldName).
//		Scan(ctx, &v)
func (rq *RodentQuery) Select(fields ...string) *RodentSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RodentSelect{RodentQuery: rq}
	sbuild.label = rodent.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RodentSelect configured with the given aggregations.
func (rq *RodentQuery) Aggregate(fns ...AggregateFunc) *RodentSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RodentQuery) prepareQuery(ctx context.Context) error {
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
		if !rodent.ValidColumn(f) {
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

func (rq *RodentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Rodent, error) {
	var (
		nodes       = []*Rodent{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [6]bool{
			rq.withDevice != nil,
			rq.withUser != nil,
			rq.withProject != nil,
			rq.withRouter != nil,
			rq.withTasks != nil,
			rq.withLoot != nil,
		}
	)
	if rq.withDevice != nil || rq.withUser != nil || rq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, rodent.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Rodent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Rodent{config: rq.config}
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
	if query := rq.withDevice; query != nil {
		if err := rq.loadDevice(ctx, query, nodes, nil,
			func(n *Rodent, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Rodent, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withProject; query != nil {
		if err := rq.loadProject(ctx, query, nodes, nil,
			func(n *Rodent, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withRouter; query != nil {
		if err := rq.loadRouter(ctx, query, nodes,
			func(n *Rodent) { n.Edges.Router = []*Router{} },
			func(n *Rodent, e *Router) { n.Edges.Router = append(n.Edges.Router, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withTasks; query != nil {
		if err := rq.loadTasks(ctx, query, nodes,
			func(n *Rodent) { n.Edges.Tasks = []*Task{} },
			func(n *Rodent, e *Task) { n.Edges.Tasks = append(n.Edges.Tasks, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withLoot; query != nil {
		if err := rq.loadLoot(ctx, query, nodes,
			func(n *Rodent) { n.Edges.Loot = []*Loot{} },
			func(n *Rodent, e *Loot) { n.Edges.Loot = append(n.Edges.Loot, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RodentQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*Rodent, init func(*Rodent), assign func(*Rodent, *Device)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Rodent)
	for i := range nodes {
		if nodes[i].device_rodents == nil {
			continue
		}
		fk := *nodes[i].device_rodents
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_rodents" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RodentQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Rodent, init func(*Rodent), assign func(*Rodent, *User)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Rodent)
	for i := range nodes {
		if nodes[i].user_rodents == nil {
			continue
		}
		fk := *nodes[i].user_rodents
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_rodents" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RodentQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Rodent, init func(*Rodent), assign func(*Rodent, *Project)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Rodent)
	for i := range nodes {
		if nodes[i].project_rodents == nil {
			continue
		}
		fk := *nodes[i].project_rodents
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
			return fmt.Errorf(`unexpected foreign-key "project_rodents" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RodentQuery) loadRouter(ctx context.Context, query *RouterQuery, nodes []*Rodent, init func(*Rodent), assign func(*Rodent, *Router)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Rodent)
	nids := make(map[xid.ID]map[*Rodent]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(rodent.RouterTable)
		s.Join(joinT).On(s.C(router.FieldID), joinT.C(rodent.RouterPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(rodent.RouterPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(rodent.RouterPrimaryKey[1]))
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
					nids[inValue] = map[*Rodent]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Router](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "router" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RodentQuery) loadTasks(ctx context.Context, query *TaskQuery, nodes []*Rodent, init func(*Rodent), assign func(*Rodent, *Task)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Rodent)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Task(func(s *sql.Selector) {
		s.Where(sql.InValues(rodent.TasksColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.rodent_tasks
		if fk == nil {
			return fmt.Errorf(`foreign-key "rodent_tasks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rodent_tasks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *RodentQuery) loadLoot(ctx context.Context, query *LootQuery, nodes []*Rodent, init func(*Rodent), assign func(*Rodent, *Loot)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Rodent)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Loot(func(s *sql.Selector) {
		s.Where(sql.InValues(rodent.LootColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.rodent_loot
		if fk == nil {
			return fmt.Errorf(`foreign-key "rodent_loot" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rodent_loot" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RodentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RodentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rodent.Table, rodent.Columns, sqlgraph.NewFieldSpec(rodent.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rodent.FieldID)
		for i := range fields {
			if fields[i] != rodent.FieldID {
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

func (rq *RodentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(rodent.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = rodent.Columns
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

// RodentGroupBy is the group-by builder for Rodent entities.
type RodentGroupBy struct {
	selector
	build *RodentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RodentGroupBy) Aggregate(fns ...AggregateFunc) *RodentGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RodentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RodentQuery, *RodentGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RodentGroupBy) sqlScan(ctx context.Context, root *RodentQuery, v any) error {
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

// RodentSelect is the builder for selecting fields of Rodent entities.
type RodentSelect struct {
	*RodentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RodentSelect) Aggregate(fns ...AggregateFunc) *RodentSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RodentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RodentQuery, *RodentSelect](ctx, rs.RodentQuery, rs, rs.inters, v)
}

func (rs *RodentSelect) sqlScan(ctx context.Context, root *RodentQuery, v any) error {
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
