// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/group"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/user"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetDescription sets the "description" field.
func (gu *GroupUpdate) SetDescription(s string) *GroupUpdate {
	gu.mutation.SetDescription(s)
	return gu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDescription(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDescription(*s)
	}
	return gu
}

// SetPermissions sets the "permissions" field.
func (gu *GroupUpdate) SetPermissions(s string) *GroupUpdate {
	gu.mutation.SetPermissions(s)
	return gu
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (gu *GroupUpdate) SetNillablePermissions(s *string) *GroupUpdate {
	if s != nil {
		gu.SetPermissions(*s)
	}
	return gu
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (gu *GroupUpdate) AddDeviceIDs(ids ...xid.ID) *GroupUpdate {
	gu.mutation.AddDeviceIDs(ids...)
	return gu
}

// AddDevices adds the "devices" edges to the Device entity.
func (gu *GroupUpdate) AddDevices(d ...*Device) *GroupUpdate {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.AddDeviceIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gu *GroupUpdate) AddUserIDs(ids ...xid.ID) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the "users" edges to the User entity.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetDomainID sets the "domain" edge to the Domain entity by ID.
func (gu *GroupUpdate) SetDomainID(id xid.ID) *GroupUpdate {
	gu.mutation.SetDomainID(id)
	return gu
}

// SetNillableDomainID sets the "domain" edge to the Domain entity by ID if the given value is not nil.
func (gu *GroupUpdate) SetNillableDomainID(id *xid.ID) *GroupUpdate {
	if id != nil {
		gu = gu.SetDomainID(*id)
	}
	return gu
}

// SetDomain sets the "domain" edge to the Domain entity.
func (gu *GroupUpdate) SetDomain(d *Domain) *GroupUpdate {
	return gu.SetDomainID(d.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (gu *GroupUpdate) ClearDevices() *GroupUpdate {
	gu.mutation.ClearDevices()
	return gu
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (gu *GroupUpdate) RemoveDeviceIDs(ids ...xid.ID) *GroupUpdate {
	gu.mutation.RemoveDeviceIDs(ids...)
	return gu
}

// RemoveDevices removes "devices" edges to Device entities.
func (gu *GroupUpdate) RemoveDevices(d ...*Device) *GroupUpdate {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.RemoveDeviceIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (gu *GroupUpdate) ClearUsers() *GroupUpdate {
	gu.mutation.ClearUsers()
	return gu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gu *GroupUpdate) RemoveUserIDs(ids ...xid.ID) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes "users" edges to User entities.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearDomain clears the "domain" edge to the Domain entity.
func (gu *GroupUpdate) ClearDomain() *GroupUpdate {
	gu.mutation.ClearDomain()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GroupMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := gu.mutation.Permissions(); ok {
		_spec.SetField(group.FieldPermissions, field.TypeString, value)
	}
	if gu.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.DevicesTable,
			Columns: group.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !gu.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.DevicesTable,
			Columns: group.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.DevicesTable,
			Columns: group.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.DomainTable,
			Columns: []string{group.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.DomainTable,
			Columns: []string{group.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetDescription sets the "description" field.
func (guo *GroupUpdateOne) SetDescription(s string) *GroupUpdateOne {
	guo.mutation.SetDescription(s)
	return guo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDescription(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDescription(*s)
	}
	return guo
}

// SetPermissions sets the "permissions" field.
func (guo *GroupUpdateOne) SetPermissions(s string) *GroupUpdateOne {
	guo.mutation.SetPermissions(s)
	return guo
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillablePermissions(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetPermissions(*s)
	}
	return guo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (guo *GroupUpdateOne) AddDeviceIDs(ids ...xid.ID) *GroupUpdateOne {
	guo.mutation.AddDeviceIDs(ids...)
	return guo
}

// AddDevices adds the "devices" edges to the Device entity.
func (guo *GroupUpdateOne) AddDevices(d ...*Device) *GroupUpdateOne {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.AddDeviceIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (guo *GroupUpdateOne) AddUserIDs(ids ...xid.ID) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the "users" edges to the User entity.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetDomainID sets the "domain" edge to the Domain entity by ID.
func (guo *GroupUpdateOne) SetDomainID(id xid.ID) *GroupUpdateOne {
	guo.mutation.SetDomainID(id)
	return guo
}

// SetNillableDomainID sets the "domain" edge to the Domain entity by ID if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDomainID(id *xid.ID) *GroupUpdateOne {
	if id != nil {
		guo = guo.SetDomainID(*id)
	}
	return guo
}

// SetDomain sets the "domain" edge to the Domain entity.
func (guo *GroupUpdateOne) SetDomain(d *Domain) *GroupUpdateOne {
	return guo.SetDomainID(d.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (guo *GroupUpdateOne) ClearDevices() *GroupUpdateOne {
	guo.mutation.ClearDevices()
	return guo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (guo *GroupUpdateOne) RemoveDeviceIDs(ids ...xid.ID) *GroupUpdateOne {
	guo.mutation.RemoveDeviceIDs(ids...)
	return guo
}

// RemoveDevices removes "devices" edges to Device entities.
func (guo *GroupUpdateOne) RemoveDevices(d ...*Device) *GroupUpdateOne {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.RemoveDeviceIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (guo *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	guo.mutation.ClearUsers()
	return guo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...xid.ID) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes "users" edges to User entities.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearDomain clears the "domain" edge to the Domain entity.
func (guo *GroupUpdateOne) ClearDomain() *GroupUpdateOne {
	guo.mutation.ClearDomain()
	return guo
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return withHooks[*Group, GroupMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := guo.mutation.Permissions(); ok {
		_spec.SetField(group.FieldPermissions, field.TypeString, value)
	}
	if guo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.DevicesTable,
			Columns: group.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !guo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.DevicesTable,
			Columns: group.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.DevicesTable,
			Columns: group.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.DomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.DomainTable,
			Columns: []string{group.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.DomainTable,
			Columns: []string{group.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
