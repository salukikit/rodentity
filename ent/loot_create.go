// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/task"
)

// LootCreate is the builder for creating a Loot entity.
type LootCreate struct {
	config
	mutation *LootMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (lc *LootCreate) SetType(l loot.Type) *LootCreate {
	lc.mutation.SetType(l)
	return lc
}

// SetLocation sets the "location" field.
func (lc *LootCreate) SetLocation(s string) *LootCreate {
	lc.mutation.SetLocation(s)
	return lc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (lc *LootCreate) SetNillableLocation(s *string) *LootCreate {
	if s != nil {
		lc.SetLocation(*s)
	}
	return lc
}

// SetData sets the "data" field.
func (lc *LootCreate) SetData(b []byte) *LootCreate {
	lc.mutation.SetData(b)
	return lc
}

// SetCollectedon sets the "collectedon" field.
func (lc *LootCreate) SetCollectedon(t time.Time) *LootCreate {
	lc.mutation.SetCollectedon(t)
	return lc
}

// SetRodentID sets the "rodent" edge to the Rodent entity by ID.
func (lc *LootCreate) SetRodentID(id int) *LootCreate {
	lc.mutation.SetRodentID(id)
	return lc
}

// SetNillableRodentID sets the "rodent" edge to the Rodent entity by ID if the given value is not nil.
func (lc *LootCreate) SetNillableRodentID(id *int) *LootCreate {
	if id != nil {
		lc = lc.SetRodentID(*id)
	}
	return lc
}

// SetRodent sets the "rodent" edge to the Rodent entity.
func (lc *LootCreate) SetRodent(r *Rodent) *LootCreate {
	return lc.SetRodentID(r.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (lc *LootCreate) SetTaskID(id int) *LootCreate {
	lc.mutation.SetTaskID(id)
	return lc
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (lc *LootCreate) SetNillableTaskID(id *int) *LootCreate {
	if id != nil {
		lc = lc.SetTaskID(*id)
	}
	return lc
}

// SetTask sets the "task" edge to the Task entity.
func (lc *LootCreate) SetTask(t *Task) *LootCreate {
	return lc.SetTaskID(t.ID)
}

// Mutation returns the LootMutation object of the builder.
func (lc *LootCreate) Mutation() *LootMutation {
	return lc.mutation
}

// Save creates the Loot in the database.
func (lc *LootCreate) Save(ctx context.Context) (*Loot, error) {
	lc.defaults()
	return withHooks[*Loot, LootMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LootCreate) SaveX(ctx context.Context) *Loot {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LootCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LootCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LootCreate) defaults() {
	if _, ok := lc.mutation.Location(); !ok {
		v := loot.DefaultLocation
		lc.mutation.SetLocation(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LootCreate) check() error {
	if _, ok := lc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Loot.type"`)}
	}
	if v, ok := lc.mutation.GetType(); ok {
		if err := loot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Loot.type": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "Loot.location"`)}
	}
	if _, ok := lc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Loot.data"`)}
	}
	if _, ok := lc.mutation.Collectedon(); !ok {
		return &ValidationError{Name: "collectedon", err: errors.New(`ent: missing required field "Loot.collectedon"`)}
	}
	return nil
}

func (lc *LootCreate) sqlSave(ctx context.Context) (*Loot, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LootCreate) createSpec() (*Loot, *sqlgraph.CreateSpec) {
	var (
		_node = &Loot{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(loot.Table, sqlgraph.NewFieldSpec(loot.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.GetType(); ok {
		_spec.SetField(loot.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := lc.mutation.Location(); ok {
		_spec.SetField(loot.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := lc.mutation.Data(); ok {
		_spec.SetField(loot.FieldData, field.TypeBytes, value)
		_node.Data = value
	}
	if value, ok := lc.mutation.Collectedon(); ok {
		_spec.SetField(loot.FieldCollectedon, field.TypeTime, value)
		_node.Collectedon = value
	}
	if nodes := lc.mutation.RodentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loot.RodentTable,
			Columns: []string{loot.RodentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rodent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.rodent_loot = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loot.TaskTable,
			Columns: []string{loot.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_loot = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LootCreateBulk is the builder for creating many Loot entities in bulk.
type LootCreateBulk struct {
	config
	builders []*LootCreate
}

// Save creates the Loot entities in the database.
func (lcb *LootCreateBulk) Save(ctx context.Context) ([]*Loot, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Loot, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LootMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LootCreateBulk) SaveX(ctx context.Context) []*Loot {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LootCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LootCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
