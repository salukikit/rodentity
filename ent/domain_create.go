// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/group"
	"github.com/salukikit/rodentity/ent/user"
)

// DomainCreate is the builder for creating a Domain entity.
type DomainCreate struct {
	config
	mutation *DomainMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DomainCreate) SetName(s string) *DomainCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetAD sets the "AD" field.
func (dc *DomainCreate) SetAD(b bool) *DomainCreate {
	dc.mutation.SetAD(b)
	return dc
}

// SetNillableAD sets the "AD" field if the given value is not nil.
func (dc *DomainCreate) SetNillableAD(b *bool) *DomainCreate {
	if b != nil {
		dc.SetAD(*b)
	}
	return dc
}

// SetOwned sets the "owned" field.
func (dc *DomainCreate) SetOwned(b bool) *DomainCreate {
	dc.mutation.SetOwned(b)
	return dc
}

// SetNillableOwned sets the "owned" field if the given value is not nil.
func (dc *DomainCreate) SetNillableOwned(b *bool) *DomainCreate {
	if b != nil {
		dc.SetOwned(*b)
	}
	return dc
}

// SetCloud sets the "cloud" field.
func (dc *DomainCreate) SetCloud(s string) *DomainCreate {
	dc.mutation.SetCloud(s)
	return dc
}

// SetNillableCloud sets the "cloud" field if the given value is not nil.
func (dc *DomainCreate) SetNillableCloud(s *string) *DomainCreate {
	if s != nil {
		dc.SetCloud(*s)
	}
	return dc
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (dc *DomainCreate) AddDeviceIDs(ids ...int) *DomainCreate {
	dc.mutation.AddDeviceIDs(ids...)
	return dc
}

// AddDevices adds the "devices" edges to the Device entity.
func (dc *DomainCreate) AddDevices(d ...*Device) *DomainCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDeviceIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (dc *DomainCreate) AddUserIDs(ids ...int) *DomainCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the "users" edges to the User entity.
func (dc *DomainCreate) AddUsers(u ...*User) *DomainCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (dc *DomainCreate) AddGroupIDs(ids ...int) *DomainCreate {
	dc.mutation.AddGroupIDs(ids...)
	return dc
}

// AddGroups adds the "groups" edges to the Group entity.
func (dc *DomainCreate) AddGroups(g ...*Group) *DomainCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return dc.AddGroupIDs(ids...)
}

// AddChilddomainIDs adds the "childdomains" edge to the Domain entity by IDs.
func (dc *DomainCreate) AddChilddomainIDs(ids ...int) *DomainCreate {
	dc.mutation.AddChilddomainIDs(ids...)
	return dc
}

// AddChilddomains adds the "childdomains" edges to the Domain entity.
func (dc *DomainCreate) AddChilddomains(d ...*Domain) *DomainCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChilddomainIDs(ids...)
}

// SetParentdomainID sets the "parentdomain" edge to the Domain entity by ID.
func (dc *DomainCreate) SetParentdomainID(id int) *DomainCreate {
	dc.mutation.SetParentdomainID(id)
	return dc
}

// SetNillableParentdomainID sets the "parentdomain" edge to the Domain entity by ID if the given value is not nil.
func (dc *DomainCreate) SetNillableParentdomainID(id *int) *DomainCreate {
	if id != nil {
		dc = dc.SetParentdomainID(*id)
	}
	return dc
}

// SetParentdomain sets the "parentdomain" edge to the Domain entity.
func (dc *DomainCreate) SetParentdomain(d *Domain) *DomainCreate {
	return dc.SetParentdomainID(d.ID)
}

// Mutation returns the DomainMutation object of the builder.
func (dc *DomainCreate) Mutation() *DomainMutation {
	return dc.mutation
}

// Save creates the Domain in the database.
func (dc *DomainCreate) Save(ctx context.Context) (*Domain, error) {
	dc.defaults()
	return withHooks[*Domain, DomainMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DomainCreate) SaveX(ctx context.Context) *Domain {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DomainCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DomainCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DomainCreate) defaults() {
	if _, ok := dc.mutation.AD(); !ok {
		v := domain.DefaultAD
		dc.mutation.SetAD(v)
	}
	if _, ok := dc.mutation.Owned(); !ok {
		v := domain.DefaultOwned
		dc.mutation.SetOwned(v)
	}
	if _, ok := dc.mutation.Cloud(); !ok {
		v := domain.DefaultCloud
		dc.mutation.SetCloud(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DomainCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Domain.name"`)}
	}
	if _, ok := dc.mutation.AD(); !ok {
		return &ValidationError{Name: "AD", err: errors.New(`ent: missing required field "Domain.AD"`)}
	}
	if _, ok := dc.mutation.Owned(); !ok {
		return &ValidationError{Name: "owned", err: errors.New(`ent: missing required field "Domain.owned"`)}
	}
	if _, ok := dc.mutation.Cloud(); !ok {
		return &ValidationError{Name: "cloud", err: errors.New(`ent: missing required field "Domain.cloud"`)}
	}
	return nil
}

func (dc *DomainCreate) sqlSave(ctx context.Context) (*Domain, error) {
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
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DomainCreate) createSpec() (*Domain, *sqlgraph.CreateSpec) {
	var (
		_node = &Domain{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(domain.Table, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt))
	)
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(domain.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.AD(); ok {
		_spec.SetField(domain.FieldAD, field.TypeBool, value)
		_node.AD = value
	}
	if value, ok := dc.mutation.Owned(); ok {
		_spec.SetField(domain.FieldOwned, field.TypeBool, value)
		_node.Owned = value
	}
	if value, ok := dc.mutation.Cloud(); ok {
		_spec.SetField(domain.FieldCloud, field.TypeString, value)
		_node.Cloud = value
	}
	if nodes := dc.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
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
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChilddomainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ParentdomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   domain.ParentdomainTable,
			Columns: []string{domain.ParentdomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.domain_childdomains = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DomainCreateBulk is the builder for creating many Domain entities in bulk.
type DomainCreateBulk struct {
	config
	builders []*DomainCreate
}

// Save creates the Domain entities in the database.
func (dcb *DomainCreateBulk) Save(ctx context.Context) ([]*Domain, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Domain, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DomainMutation)
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DomainCreateBulk) SaveX(ctx context.Context) []*Domain {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DomainCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DomainCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
