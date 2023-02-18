// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/services"
	"github.com/salukikit/rodentity/ent/subnet"
)

// ServicesUpdate is the builder for updating Services entities.
type ServicesUpdate struct {
	config
	hooks    []Hook
	mutation *ServicesMutation
}

// Where appends a list predicates to the ServicesUpdate builder.
func (su *ServicesUpdate) Where(ps ...predicate.Services) *ServicesUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetServiceName sets the "service_name" field.
func (su *ServicesUpdate) SetServiceName(s string) *ServicesUpdate {
	su.mutation.SetServiceName(s)
	return su
}

// SetPort sets the "port" field.
func (su *ServicesUpdate) SetPort(s string) *ServicesUpdate {
	su.mutation.SetPort(s)
	return su
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (su *ServicesUpdate) AddDeviceIDs(ids ...int) *ServicesUpdate {
	su.mutation.AddDeviceIDs(ids...)
	return su
}

// AddDevices adds the "devices" edges to the Device entity.
func (su *ServicesUpdate) AddDevices(d ...*Device) *ServicesUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.AddDeviceIDs(ids...)
}

// AddSubnetIDs adds the "subnet" edge to the Subnet entity by IDs.
func (su *ServicesUpdate) AddSubnetIDs(ids ...int) *ServicesUpdate {
	su.mutation.AddSubnetIDs(ids...)
	return su
}

// AddSubnet adds the "subnet" edges to the Subnet entity.
func (su *ServicesUpdate) AddSubnet(s ...*Subnet) *ServicesUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSubnetIDs(ids...)
}

// Mutation returns the ServicesMutation object of the builder.
func (su *ServicesUpdate) Mutation() *ServicesMutation {
	return su.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (su *ServicesUpdate) ClearDevices() *ServicesUpdate {
	su.mutation.ClearDevices()
	return su
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (su *ServicesUpdate) RemoveDeviceIDs(ids ...int) *ServicesUpdate {
	su.mutation.RemoveDeviceIDs(ids...)
	return su
}

// RemoveDevices removes "devices" edges to Device entities.
func (su *ServicesUpdate) RemoveDevices(d ...*Device) *ServicesUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.RemoveDeviceIDs(ids...)
}

// ClearSubnet clears all "subnet" edges to the Subnet entity.
func (su *ServicesUpdate) ClearSubnet() *ServicesUpdate {
	su.mutation.ClearSubnet()
	return su
}

// RemoveSubnetIDs removes the "subnet" edge to Subnet entities by IDs.
func (su *ServicesUpdate) RemoveSubnetIDs(ids ...int) *ServicesUpdate {
	su.mutation.RemoveSubnetIDs(ids...)
	return su
}

// RemoveSubnet removes "subnet" edges to Subnet entities.
func (su *ServicesUpdate) RemoveSubnet(s ...*Subnet) *ServicesUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSubnetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServicesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ServicesMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServicesUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServicesUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServicesUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ServicesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(services.Table, services.Columns, sqlgraph.NewFieldSpec(services.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.ServiceName(); ok {
		_spec.SetField(services.FieldServiceName, field.TypeString, value)
	}
	if value, ok := su.mutation.Port(); ok {
		_spec.SetField(services.FieldPort, field.TypeString, value)
	}
	if su.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   services.DevicesTable,
			Columns: services.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !su.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   services.DevicesTable,
			Columns: services.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   services.DevicesTable,
			Columns: services.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.SubnetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   services.SubnetTable,
			Columns: []string{services.SubnetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subnet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSubnetIDs(); len(nodes) > 0 && !su.mutation.SubnetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   services.SubnetTable,
			Columns: []string{services.SubnetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subnet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SubnetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   services.SubnetTable,
			Columns: []string{services.SubnetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subnet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{services.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ServicesUpdateOne is the builder for updating a single Services entity.
type ServicesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServicesMutation
}

// SetServiceName sets the "service_name" field.
func (suo *ServicesUpdateOne) SetServiceName(s string) *ServicesUpdateOne {
	suo.mutation.SetServiceName(s)
	return suo
}

// SetPort sets the "port" field.
func (suo *ServicesUpdateOne) SetPort(s string) *ServicesUpdateOne {
	suo.mutation.SetPort(s)
	return suo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (suo *ServicesUpdateOne) AddDeviceIDs(ids ...int) *ServicesUpdateOne {
	suo.mutation.AddDeviceIDs(ids...)
	return suo
}

// AddDevices adds the "devices" edges to the Device entity.
func (suo *ServicesUpdateOne) AddDevices(d ...*Device) *ServicesUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.AddDeviceIDs(ids...)
}

// AddSubnetIDs adds the "subnet" edge to the Subnet entity by IDs.
func (suo *ServicesUpdateOne) AddSubnetIDs(ids ...int) *ServicesUpdateOne {
	suo.mutation.AddSubnetIDs(ids...)
	return suo
}

// AddSubnet adds the "subnet" edges to the Subnet entity.
func (suo *ServicesUpdateOne) AddSubnet(s ...*Subnet) *ServicesUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSubnetIDs(ids...)
}

// Mutation returns the ServicesMutation object of the builder.
func (suo *ServicesUpdateOne) Mutation() *ServicesMutation {
	return suo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (suo *ServicesUpdateOne) ClearDevices() *ServicesUpdateOne {
	suo.mutation.ClearDevices()
	return suo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (suo *ServicesUpdateOne) RemoveDeviceIDs(ids ...int) *ServicesUpdateOne {
	suo.mutation.RemoveDeviceIDs(ids...)
	return suo
}

// RemoveDevices removes "devices" edges to Device entities.
func (suo *ServicesUpdateOne) RemoveDevices(d ...*Device) *ServicesUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.RemoveDeviceIDs(ids...)
}

// ClearSubnet clears all "subnet" edges to the Subnet entity.
func (suo *ServicesUpdateOne) ClearSubnet() *ServicesUpdateOne {
	suo.mutation.ClearSubnet()
	return suo
}

// RemoveSubnetIDs removes the "subnet" edge to Subnet entities by IDs.
func (suo *ServicesUpdateOne) RemoveSubnetIDs(ids ...int) *ServicesUpdateOne {
	suo.mutation.RemoveSubnetIDs(ids...)
	return suo
}

// RemoveSubnet removes "subnet" edges to Subnet entities.
func (suo *ServicesUpdateOne) RemoveSubnet(s ...*Subnet) *ServicesUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSubnetIDs(ids...)
}

// Where appends a list predicates to the ServicesUpdate builder.
func (suo *ServicesUpdateOne) Where(ps ...predicate.Services) *ServicesUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServicesUpdateOne) Select(field string, fields ...string) *ServicesUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Services entity.
func (suo *ServicesUpdateOne) Save(ctx context.Context) (*Services, error) {
	return withHooks[*Services, ServicesMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServicesUpdateOne) SaveX(ctx context.Context) *Services {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServicesUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServicesUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ServicesUpdateOne) sqlSave(ctx context.Context) (_node *Services, err error) {
	_spec := sqlgraph.NewUpdateSpec(services.Table, services.Columns, sqlgraph.NewFieldSpec(services.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Services.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, services.FieldID)
		for _, f := range fields {
			if !services.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != services.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.ServiceName(); ok {
		_spec.SetField(services.FieldServiceName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Port(); ok {
		_spec.SetField(services.FieldPort, field.TypeString, value)
	}
	if suo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   services.DevicesTable,
			Columns: services.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !suo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   services.DevicesTable,
			Columns: services.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   services.DevicesTable,
			Columns: services.DevicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.SubnetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   services.SubnetTable,
			Columns: []string{services.SubnetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subnet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSubnetIDs(); len(nodes) > 0 && !suo.mutation.SubnetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   services.SubnetTable,
			Columns: []string{services.SubnetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subnet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SubnetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   services.SubnetTable,
			Columns: []string{services.SubnetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subnet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Services{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{services.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}