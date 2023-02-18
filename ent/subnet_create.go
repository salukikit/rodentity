// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/subnet"
)

// SubnetCreate is the builder for creating a Subnet entity.
type SubnetCreate struct {
	config
	mutation *SubnetMutation
	hooks    []Hook
}

// SetCidr sets the "cidr" field.
func (sc *SubnetCreate) SetCidr(s string) *SubnetCreate {
	sc.mutation.SetCidr(s)
	return sc
}

// SetMask sets the "mask" field.
func (sc *SubnetCreate) SetMask(b []byte) *SubnetCreate {
	sc.mutation.SetMask(b)
	return sc
}

// AddHostIDs adds the "hosts" edge to the Device entity by IDs.
func (sc *SubnetCreate) AddHostIDs(ids ...int) *SubnetCreate {
	sc.mutation.AddHostIDs(ids...)
	return sc
}

// AddHosts adds the "hosts" edges to the Device entity.
func (sc *SubnetCreate) AddHosts(d ...*Device) *SubnetCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sc.AddHostIDs(ids...)
}

// Mutation returns the SubnetMutation object of the builder.
func (sc *SubnetCreate) Mutation() *SubnetMutation {
	return sc.mutation
}

// Save creates the Subnet in the database.
func (sc *SubnetCreate) Save(ctx context.Context) (*Subnet, error) {
	return withHooks[*Subnet, SubnetMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubnetCreate) SaveX(ctx context.Context) *Subnet {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubnetCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubnetCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubnetCreate) check() error {
	if _, ok := sc.mutation.Cidr(); !ok {
		return &ValidationError{Name: "cidr", err: errors.New(`ent: missing required field "Subnet.cidr"`)}
	}
	return nil
}

func (sc *SubnetCreate) sqlSave(ctx context.Context) (*Subnet, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubnetCreate) createSpec() (*Subnet, *sqlgraph.CreateSpec) {
	var (
		_node = &Subnet{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subnet.Table, sqlgraph.NewFieldSpec(subnet.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Cidr(); ok {
		_spec.SetField(subnet.FieldCidr, field.TypeString, value)
		_node.Cidr = value
	}
	if value, ok := sc.mutation.Mask(); ok {
		_spec.SetField(subnet.FieldMask, field.TypeBytes, value)
		_node.Mask = value
	}
	if nodes := sc.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subnet.HostsTable,
			Columns: subnet.HostsPrimaryKey,
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
	return _node, _spec
}

// SubnetCreateBulk is the builder for creating many Subnet entities in bulk.
type SubnetCreateBulk struct {
	config
	builders []*SubnetCreate
}

// Save creates the Subnet entities in the database.
func (scb *SubnetCreateBulk) Save(ctx context.Context) ([]*Subnet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subnet, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubnetMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubnetCreateBulk) SaveX(ctx context.Context) []*Subnet {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubnetCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubnetCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}