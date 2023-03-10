// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/group"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/services"
	"github.com/salukikit/rodentity/ent/subnet"
	"github.com/salukikit/rodentity/ent/user"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
}

// SetHostname sets the "hostname" field.
func (dc *DeviceCreate) SetHostname(s string) *DeviceCreate {
	dc.mutation.SetHostname(s)
	return dc
}

// SetOs sets the "os" field.
func (dc *DeviceCreate) SetOs(s string) *DeviceCreate {
	dc.mutation.SetOs(s)
	return dc
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableOs(s *string) *DeviceCreate {
	if s != nil {
		dc.SetOs(*s)
	}
	return dc
}

// SetArch sets the "arch" field.
func (dc *DeviceCreate) SetArch(s string) *DeviceCreate {
	dc.mutation.SetArch(s)
	return dc
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableArch(s *string) *DeviceCreate {
	if s != nil {
		dc.SetArch(*s)
	}
	return dc
}

// SetVersion sets the "version" field.
func (dc *DeviceCreate) SetVersion(s string) *DeviceCreate {
	dc.mutation.SetVersion(s)
	return dc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableVersion(s *string) *DeviceCreate {
	if s != nil {
		dc.SetVersion(*s)
	}
	return dc
}

// SetNetInterfaces sets the "net_interfaces" field.
func (dc *DeviceCreate) SetNetInterfaces(s []string) *DeviceCreate {
	dc.mutation.SetNetInterfaces(s)
	return dc
}

// SetMachinepass sets the "machinepass" field.
func (dc *DeviceCreate) SetMachinepass(s string) *DeviceCreate {
	dc.mutation.SetMachinepass(s)
	return dc
}

// SetNillableMachinepass sets the "machinepass" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableMachinepass(s *string) *DeviceCreate {
	if s != nil {
		dc.SetMachinepass(*s)
	}
	return dc
}

// SetCertificates sets the "certificates" field.
func (dc *DeviceCreate) SetCertificates(s []string) *DeviceCreate {
	dc.mutation.SetCertificates(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DeviceCreate) SetID(x xid.ID) *DeviceCreate {
	dc.mutation.SetID(x)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableID(x *xid.ID) *DeviceCreate {
	if x != nil {
		dc.SetID(*x)
	}
	return dc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (dc *DeviceCreate) AddUserIDs(ids ...xid.ID) *DeviceCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the "users" edges to the User entity.
func (dc *DeviceCreate) AddUsers(u ...*User) *DeviceCreate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddRodentIDs adds the "rodents" edge to the Rodent entity by IDs.
func (dc *DeviceCreate) AddRodentIDs(ids ...xid.ID) *DeviceCreate {
	dc.mutation.AddRodentIDs(ids...)
	return dc
}

// AddRodents adds the "rodents" edges to the Rodent entity.
func (dc *DeviceCreate) AddRodents(r ...*Rodent) *DeviceCreate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dc.AddRodentIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (dc *DeviceCreate) AddGroupIDs(ids ...xid.ID) *DeviceCreate {
	dc.mutation.AddGroupIDs(ids...)
	return dc
}

// AddGroups adds the "groups" edges to the Group entity.
func (dc *DeviceCreate) AddGroups(g ...*Group) *DeviceCreate {
	ids := make([]xid.ID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return dc.AddGroupIDs(ids...)
}

// SetDomainID sets the "domain" edge to the Domain entity by ID.
func (dc *DeviceCreate) SetDomainID(id xid.ID) *DeviceCreate {
	dc.mutation.SetDomainID(id)
	return dc
}

// SetNillableDomainID sets the "domain" edge to the Domain entity by ID if the given value is not nil.
func (dc *DeviceCreate) SetNillableDomainID(id *xid.ID) *DeviceCreate {
	if id != nil {
		dc = dc.SetDomainID(*id)
	}
	return dc
}

// SetDomain sets the "domain" edge to the Domain entity.
func (dc *DeviceCreate) SetDomain(d *Domain) *DeviceCreate {
	return dc.SetDomainID(d.ID)
}

// AddSubnetIDs adds the "subnets" edge to the Subnet entity by IDs.
func (dc *DeviceCreate) AddSubnetIDs(ids ...xid.ID) *DeviceCreate {
	dc.mutation.AddSubnetIDs(ids...)
	return dc
}

// AddSubnets adds the "subnets" edges to the Subnet entity.
func (dc *DeviceCreate) AddSubnets(s ...*Subnet) *DeviceCreate {
	ids := make([]xid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddSubnetIDs(ids...)
}

// AddServiceIDs adds the "services" edge to the Services entity by IDs.
func (dc *DeviceCreate) AddServiceIDs(ids ...int) *DeviceCreate {
	dc.mutation.AddServiceIDs(ids...)
	return dc
}

// AddServices adds the "services" edges to the Services entity.
func (dc *DeviceCreate) AddServices(s ...*Services) *DeviceCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddServiceIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	dc.defaults()
	return withHooks[*Device, DeviceMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeviceCreate) defaults() {
	if _, ok := dc.mutation.Os(); !ok {
		v := device.DefaultOs
		dc.mutation.SetOs(v)
	}
	if _, ok := dc.mutation.Arch(); !ok {
		v := device.DefaultArch
		dc.mutation.SetArch(v)
	}
	if _, ok := dc.mutation.Version(); !ok {
		v := device.DefaultVersion
		dc.mutation.SetVersion(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := device.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	if _, ok := dc.mutation.Hostname(); !ok {
		return &ValidationError{Name: "hostname", err: errors.New(`ent: missing required field "Device.hostname"`)}
	}
	if _, ok := dc.mutation.Os(); !ok {
		return &ValidationError{Name: "os", err: errors.New(`ent: missing required field "Device.os"`)}
	}
	if _, ok := dc.mutation.Arch(); !ok {
		return &ValidationError{Name: "arch", err: errors.New(`ent: missing required field "Device.arch"`)}
	}
	if _, ok := dc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Device.version"`)}
	}
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(device.Table, sqlgraph.NewFieldSpec(device.FieldID, field.TypeString))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.Hostname(); ok {
		_spec.SetField(device.FieldHostname, field.TypeString, value)
		_node.Hostname = value
	}
	if value, ok := dc.mutation.Os(); ok {
		_spec.SetField(device.FieldOs, field.TypeString, value)
		_node.Os = value
	}
	if value, ok := dc.mutation.Arch(); ok {
		_spec.SetField(device.FieldArch, field.TypeString, value)
		_node.Arch = value
	}
	if value, ok := dc.mutation.Version(); ok {
		_spec.SetField(device.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := dc.mutation.NetInterfaces(); ok {
		_spec.SetField(device.FieldNetInterfaces, field.TypeJSON, value)
		_node.NetInterfaces = value
	}
	if value, ok := dc.mutation.Machinepass(); ok {
		_spec.SetField(device.FieldMachinepass, field.TypeString, value)
		_node.Machinepass = value
	}
	if value, ok := dc.mutation.Certificates(); ok {
		_spec.SetField(device.FieldCertificates, field.TypeJSON, value)
		_node.Certificates = value
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   device.UsersTable,
			Columns: device.UsersPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.RodentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.RodentsTable,
			Columns: []string{device.RodentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: rodent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   device.GroupsTable,
			Columns: device.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.DomainTable,
			Columns: []string{device.DomainColumn},
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
		_node.domain_devices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.SubnetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   device.SubnetsTable,
			Columns: device.SubnetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: subnet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   device.ServicesTable,
			Columns: device.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: services.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	builders []*DeviceCreate
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
