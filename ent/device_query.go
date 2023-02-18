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
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/group"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/subnet"
	"github.com/salukikit/rodentity/ent/user"
)

// DeviceQuery is the builder for querying Device entities.
type DeviceQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.Device
	withUsers   *UserQuery
	withRodents *RodentQuery
	withGroups  *GroupQuery
	withDomain  *DomainQuery
	withSubnets *SubnetQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceQuery builder.
func (dq *DeviceQuery) Where(ps ...predicate.Device) *DeviceQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DeviceQuery) Limit(limit int) *DeviceQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DeviceQuery) Offset(offset int) *DeviceQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeviceQuery) Unique(unique bool) *DeviceQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DeviceQuery) Order(o ...OrderFunc) *DeviceQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryUsers chains the current query on the "users" edge.
func (dq *DeviceQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, device.UsersTable, device.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRodents chains the current query on the "rodents" edge.
func (dq *DeviceQuery) QueryRodents() *RodentQuery {
	query := (&RodentClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(rodent.Table, rodent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, device.RodentsTable, device.RodentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (dq *DeviceQuery) QueryGroups() *GroupQuery {
	query := (&GroupClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, device.GroupsTable, device.GroupsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDomain chains the current query on the "domain" edge.
func (dq *DeviceQuery) QueryDomain() *DomainQuery {
	query := (&DomainClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(domain.Table, domain.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, device.DomainTable, device.DomainColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubnets chains the current query on the "subnets" edge.
func (dq *DeviceQuery) QuerySubnets() *SubnetQuery {
	query := (&SubnetClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(subnet.Table, subnet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, device.SubnetsTable, device.SubnetsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Device entity from the query.
// Returns a *NotFoundError when no Device was found.
func (dq *DeviceQuery) First(ctx context.Context) (*Device, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{device.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeviceQuery) FirstX(ctx context.Context) *Device {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Device ID from the query.
// Returns a *NotFoundError when no Device ID was found.
func (dq *DeviceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{device.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeviceQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Device entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Device entity is found.
// Returns a *NotFoundError when no Device entities are found.
func (dq *DeviceQuery) Only(ctx context.Context) (*Device, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{device.Label}
	default:
		return nil, &NotSingularError{device.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeviceQuery) OnlyX(ctx context.Context) *Device {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Device ID in the query.
// Returns a *NotSingularError when more than one Device ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeviceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{device.Label}
	default:
		err = &NotSingularError{device.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeviceQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Devices.
func (dq *DeviceQuery) All(ctx context.Context) ([]*Device, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Device, *DeviceQuery]()
	return withInterceptors[[]*Device](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeviceQuery) AllX(ctx context.Context) []*Device {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Device IDs.
func (dq *DeviceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(device.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeviceQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeviceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DeviceQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeviceQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeviceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeviceQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeviceQuery) Clone() *DeviceQuery {
	if dq == nil {
		return nil
	}
	return &DeviceQuery{
		config:      dq.config,
		ctx:         dq.ctx.Clone(),
		order:       append([]OrderFunc{}, dq.order...),
		inters:      append([]Interceptor{}, dq.inters...),
		predicates:  append([]predicate.Device{}, dq.predicates...),
		withUsers:   dq.withUsers.Clone(),
		withRodents: dq.withRodents.Clone(),
		withGroups:  dq.withGroups.Clone(),
		withDomain:  dq.withDomain.Clone(),
		withSubnets: dq.withSubnets.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithUsers(opts ...func(*UserQuery)) *DeviceQuery {
	query := (&UserClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withUsers = query
	return dq
}

// WithRodents tells the query-builder to eager-load the nodes that are connected to
// the "rodents" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithRodents(opts ...func(*RodentQuery)) *DeviceQuery {
	query := (&RodentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withRodents = query
	return dq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithGroups(opts ...func(*GroupQuery)) *DeviceQuery {
	query := (&GroupClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withGroups = query
	return dq
}

// WithDomain tells the query-builder to eager-load the nodes that are connected to
// the "domain" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithDomain(opts ...func(*DomainQuery)) *DeviceQuery {
	query := (&DomainClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDomain = query
	return dq
}

// WithSubnets tells the query-builder to eager-load the nodes that are connected to
// the "subnets" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithSubnets(opts ...func(*SubnetQuery)) *DeviceQuery {
	query := (&SubnetClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withSubnets = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hostname string `json:"hostname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Device.Query().
//		GroupBy(device.FieldHostname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DeviceQuery) GroupBy(field string, fields ...string) *DeviceGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeviceGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = device.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Hostname string `json:"hostname,omitempty"`
//	}
//
//	client.Device.Query().
//		Select(device.FieldHostname).
//		Scan(ctx, &v)
func (dq *DeviceQuery) Select(fields ...string) *DeviceSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DeviceSelect{DeviceQuery: dq}
	sbuild.label = device.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeviceSelect configured with the given aggregations.
func (dq *DeviceQuery) Aggregate(fns ...AggregateFunc) *DeviceSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DeviceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !device.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeviceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Device, error) {
	var (
		nodes       = []*Device{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [5]bool{
			dq.withUsers != nil,
			dq.withRodents != nil,
			dq.withGroups != nil,
			dq.withDomain != nil,
			dq.withSubnets != nil,
		}
	)
	if dq.withDomain != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, device.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Device).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Device{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withUsers; query != nil {
		if err := dq.loadUsers(ctx, query, nodes,
			func(n *Device) { n.Edges.Users = []*User{} },
			func(n *Device, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withRodents; query != nil {
		if err := dq.loadRodents(ctx, query, nodes,
			func(n *Device) { n.Edges.Rodents = []*Rodent{} },
			func(n *Device, e *Rodent) { n.Edges.Rodents = append(n.Edges.Rodents, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withGroups; query != nil {
		if err := dq.loadGroups(ctx, query, nodes,
			func(n *Device) { n.Edges.Groups = []*Group{} },
			func(n *Device, e *Group) { n.Edges.Groups = append(n.Edges.Groups, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withDomain; query != nil {
		if err := dq.loadDomain(ctx, query, nodes, nil,
			func(n *Device, e *Domain) { n.Edges.Domain = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withSubnets; query != nil {
		if err := dq.loadSubnets(ctx, query, nodes,
			func(n *Device) { n.Edges.Subnets = []*Subnet{} },
			func(n *Device, e *Subnet) { n.Edges.Subnets = append(n.Edges.Subnets, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeviceQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Device, init func(*Device), assign func(*Device, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Device)
	nids := make(map[int]map[*Device]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(device.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(device.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(device.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(device.UsersPrimaryKey[0]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Device]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DeviceQuery) loadRodents(ctx context.Context, query *RodentQuery, nodes []*Device, init func(*Device), assign func(*Device, *Rodent)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Device)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Rodent(func(s *sql.Selector) {
		s.Where(sql.InValues(device.RodentsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.device_rodents
		if fk == nil {
			return fmt.Errorf(`foreign-key "device_rodents" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_rodents" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DeviceQuery) loadGroups(ctx context.Context, query *GroupQuery, nodes []*Device, init func(*Device), assign func(*Device, *Group)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Device)
	nids := make(map[int]map[*Device]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(device.GroupsTable)
		s.Join(joinT).On(s.C(group.FieldID), joinT.C(device.GroupsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(device.GroupsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(device.GroupsPrimaryKey[1]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Device]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Group](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "groups" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DeviceQuery) loadDomain(ctx context.Context, query *DomainQuery, nodes []*Device, init func(*Device), assign func(*Device, *Domain)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Device)
	for i := range nodes {
		if nodes[i].domain_devices == nil {
			continue
		}
		fk := *nodes[i].domain_devices
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(domain.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "domain_devices" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DeviceQuery) loadSubnets(ctx context.Context, query *SubnetQuery, nodes []*Device, init func(*Device), assign func(*Device, *Subnet)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Device)
	nids := make(map[int]map[*Device]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(device.SubnetsTable)
		s.Join(joinT).On(s.C(subnet.FieldID), joinT.C(device.SubnetsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(device.SubnetsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(device.SubnetsPrimaryKey[1]))
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
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Device]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Subnet](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "subnets" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (dq *DeviceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeviceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for i := range fields {
			if fields[i] != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeviceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(device.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = device.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceGroupBy is the group-by builder for Device entities.
type DeviceGroupBy struct {
	selector
	build *DeviceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeviceGroupBy) Aggregate(fns ...AggregateFunc) *DeviceGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DeviceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceQuery, *DeviceGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DeviceGroupBy) sqlScan(ctx context.Context, root *DeviceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeviceSelect is the builder for selecting fields of Device entities.
type DeviceSelect struct {
	*DeviceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DeviceSelect) Aggregate(fns ...AggregateFunc) *DeviceSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeviceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceQuery, *DeviceSelect](ctx, ds.DeviceQuery, ds, ds.inters, v)
}

func (ds *DeviceSelect) sqlScan(ctx context.Context, root *DeviceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
