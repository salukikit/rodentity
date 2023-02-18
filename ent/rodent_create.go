// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/project"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/router"
	"github.com/salukikit/rodentity/ent/task"
	"github.com/salukikit/rodentity/ent/user"
)

// RodentCreate is the builder for creating a Rodent entity.
type RodentCreate struct {
	config
	mutation *RodentMutation
	hooks    []Hook
}

// SetXid sets the "xid" field.
func (rc *RodentCreate) SetXid(s string) *RodentCreate {
	rc.mutation.SetXid(s)
	return rc
}

// SetType sets the "type" field.
func (rc *RodentCreate) SetType(s string) *RodentCreate {
	rc.mutation.SetType(s)
	return rc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rc *RodentCreate) SetNillableType(s *string) *RodentCreate {
	if s != nil {
		rc.SetType(*s)
	}
	return rc
}

// SetCodename sets the "codename" field.
func (rc *RodentCreate) SetCodename(s string) *RodentCreate {
	rc.mutation.SetCodename(s)
	return rc
}

// SetKey sets the "key" field.
func (rc *RodentCreate) SetKey(s string) *RodentCreate {
	rc.mutation.SetKey(s)
	return rc
}

// SetUsercontext sets the "usercontext" field.
func (rc *RodentCreate) SetUsercontext(s string) *RodentCreate {
	rc.mutation.SetUsercontext(s)
	return rc
}

// SetNillableUsercontext sets the "usercontext" field if the given value is not nil.
func (rc *RodentCreate) SetNillableUsercontext(s *string) *RodentCreate {
	if s != nil {
		rc.SetUsercontext(*s)
	}
	return rc
}

// SetBeacontime sets the "beacontime" field.
func (rc *RodentCreate) SetBeacontime(s string) *RodentCreate {
	rc.mutation.SetBeacontime(s)
	return rc
}

// SetNillableBeacontime sets the "beacontime" field if the given value is not nil.
func (rc *RodentCreate) SetNillableBeacontime(s *string) *RodentCreate {
	if s != nil {
		rc.SetBeacontime(*s)
	}
	return rc
}

// SetBurned sets the "burned" field.
func (rc *RodentCreate) SetBurned(b bool) *RodentCreate {
	rc.mutation.SetBurned(b)
	return rc
}

// SetNillableBurned sets the "burned" field if the given value is not nil.
func (rc *RodentCreate) SetNillableBurned(b *bool) *RodentCreate {
	if b != nil {
		rc.SetBurned(*b)
	}
	return rc
}

// SetAlive sets the "alive" field.
func (rc *RodentCreate) SetAlive(b bool) *RodentCreate {
	rc.mutation.SetAlive(b)
	return rc
}

// SetNillableAlive sets the "alive" field if the given value is not nil.
func (rc *RodentCreate) SetNillableAlive(b *bool) *RodentCreate {
	if b != nil {
		rc.SetAlive(*b)
	}
	return rc
}

// SetJoined sets the "joined" field.
func (rc *RodentCreate) SetJoined(t time.Time) *RodentCreate {
	rc.mutation.SetJoined(t)
	return rc
}

// SetLastseen sets the "lastseen" field.
func (rc *RodentCreate) SetLastseen(t time.Time) *RodentCreate {
	rc.mutation.SetLastseen(t)
	return rc
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (rc *RodentCreate) SetDeviceID(id int) *RodentCreate {
	rc.mutation.SetDeviceID(id)
	return rc
}

// SetNillableDeviceID sets the "device" edge to the Device entity by ID if the given value is not nil.
func (rc *RodentCreate) SetNillableDeviceID(id *int) *RodentCreate {
	if id != nil {
		rc = rc.SetDeviceID(*id)
	}
	return rc
}

// SetDevice sets the "device" edge to the Device entity.
func (rc *RodentCreate) SetDevice(d *Device) *RodentCreate {
	return rc.SetDeviceID(d.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rc *RodentCreate) SetUserID(id int) *RodentCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (rc *RodentCreate) SetNillableUserID(id *int) *RodentCreate {
	if id != nil {
		rc = rc.SetUserID(*id)
	}
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RodentCreate) SetUser(u *User) *RodentCreate {
	return rc.SetUserID(u.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (rc *RodentCreate) SetProjectID(id int) *RodentCreate {
	rc.mutation.SetProjectID(id)
	return rc
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (rc *RodentCreate) SetNillableProjectID(id *int) *RodentCreate {
	if id != nil {
		rc = rc.SetProjectID(*id)
	}
	return rc
}

// SetProject sets the "project" edge to the Project entity.
func (rc *RodentCreate) SetProject(p *Project) *RodentCreate {
	return rc.SetProjectID(p.ID)
}

// SetRouterID sets the "router" edge to the Router entity by ID.
func (rc *RodentCreate) SetRouterID(id int) *RodentCreate {
	rc.mutation.SetRouterID(id)
	return rc
}

// SetNillableRouterID sets the "router" edge to the Router entity by ID if the given value is not nil.
func (rc *RodentCreate) SetNillableRouterID(id *int) *RodentCreate {
	if id != nil {
		rc = rc.SetRouterID(*id)
	}
	return rc
}

// SetRouter sets the "router" edge to the Router entity.
func (rc *RodentCreate) SetRouter(r *Router) *RodentCreate {
	return rc.SetRouterID(r.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (rc *RodentCreate) AddTaskIDs(ids ...int) *RodentCreate {
	rc.mutation.AddTaskIDs(ids...)
	return rc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (rc *RodentCreate) AddTasks(t ...*Task) *RodentCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddTaskIDs(ids...)
}

// AddLootIDs adds the "loot" edge to the Loot entity by IDs.
func (rc *RodentCreate) AddLootIDs(ids ...int) *RodentCreate {
	rc.mutation.AddLootIDs(ids...)
	return rc
}

// AddLoot adds the "loot" edges to the Loot entity.
func (rc *RodentCreate) AddLoot(l ...*Loot) *RodentCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return rc.AddLootIDs(ids...)
}

// Mutation returns the RodentMutation object of the builder.
func (rc *RodentCreate) Mutation() *RodentMutation {
	return rc.mutation
}

// Save creates the Rodent in the database.
func (rc *RodentCreate) Save(ctx context.Context) (*Rodent, error) {
	rc.defaults()
	return withHooks[*Rodent, RodentMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RodentCreate) SaveX(ctx context.Context) *Rodent {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RodentCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RodentCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RodentCreate) defaults() {
	if _, ok := rc.mutation.GetType(); !ok {
		v := rodent.DefaultType
		rc.mutation.SetType(v)
	}
	if _, ok := rc.mutation.Burned(); !ok {
		v := rodent.DefaultBurned
		rc.mutation.SetBurned(v)
	}
	if _, ok := rc.mutation.Alive(); !ok {
		v := rodent.DefaultAlive
		rc.mutation.SetAlive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RodentCreate) check() error {
	if _, ok := rc.mutation.Xid(); !ok {
		return &ValidationError{Name: "xid", err: errors.New(`ent: missing required field "Rodent.xid"`)}
	}
	if _, ok := rc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Rodent.type"`)}
	}
	if _, ok := rc.mutation.Codename(); !ok {
		return &ValidationError{Name: "codename", err: errors.New(`ent: missing required field "Rodent.codename"`)}
	}
	if _, ok := rc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Rodent.key"`)}
	}
	if _, ok := rc.mutation.Burned(); !ok {
		return &ValidationError{Name: "burned", err: errors.New(`ent: missing required field "Rodent.burned"`)}
	}
	if _, ok := rc.mutation.Alive(); !ok {
		return &ValidationError{Name: "alive", err: errors.New(`ent: missing required field "Rodent.alive"`)}
	}
	if _, ok := rc.mutation.Joined(); !ok {
		return &ValidationError{Name: "joined", err: errors.New(`ent: missing required field "Rodent.joined"`)}
	}
	if _, ok := rc.mutation.Lastseen(); !ok {
		return &ValidationError{Name: "lastseen", err: errors.New(`ent: missing required field "Rodent.lastseen"`)}
	}
	return nil
}

func (rc *RodentCreate) sqlSave(ctx context.Context) (*Rodent, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RodentCreate) createSpec() (*Rodent, *sqlgraph.CreateSpec) {
	var (
		_node = &Rodent{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(rodent.Table, sqlgraph.NewFieldSpec(rodent.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Xid(); ok {
		_spec.SetField(rodent.FieldXid, field.TypeString, value)
		_node.Xid = value
	}
	if value, ok := rc.mutation.GetType(); ok {
		_spec.SetField(rodent.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := rc.mutation.Codename(); ok {
		_spec.SetField(rodent.FieldCodename, field.TypeString, value)
		_node.Codename = value
	}
	if value, ok := rc.mutation.Key(); ok {
		_spec.SetField(rodent.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := rc.mutation.Usercontext(); ok {
		_spec.SetField(rodent.FieldUsercontext, field.TypeString, value)
		_node.Usercontext = value
	}
	if value, ok := rc.mutation.Beacontime(); ok {
		_spec.SetField(rodent.FieldBeacontime, field.TypeString, value)
		_node.Beacontime = value
	}
	if value, ok := rc.mutation.Burned(); ok {
		_spec.SetField(rodent.FieldBurned, field.TypeBool, value)
		_node.Burned = value
	}
	if value, ok := rc.mutation.Alive(); ok {
		_spec.SetField(rodent.FieldAlive, field.TypeBool, value)
		_node.Alive = value
	}
	if value, ok := rc.mutation.Joined(); ok {
		_spec.SetField(rodent.FieldJoined, field.TypeTime, value)
		_node.Joined = value
	}
	if value, ok := rc.mutation.Lastseen(); ok {
		_spec.SetField(rodent.FieldLastseen, field.TypeTime, value)
		_node.Lastseen = value
	}
	if nodes := rc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rodent.DeviceTable,
			Columns: []string{rodent.DeviceColumn},
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
		_node.device_rodents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rodent.UserTable,
			Columns: []string{rodent.UserColumn},
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
		_node.user_rodents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rodent.ProjectTable,
			Columns: []string{rodent.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.project_rodents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RouterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rodent.RouterTable,
			Columns: []string{rodent.RouterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: router.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.router_rodents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rodent.TasksTable,
			Columns: []string{rodent.TasksColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.LootIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rodent.LootTable,
			Columns: []string{rodent.LootColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: loot.FieldID,
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

// RodentCreateBulk is the builder for creating many Rodent entities in bulk.
type RodentCreateBulk struct {
	config
	builders []*RodentCreate
}

// Save creates the Rodent entities in the database.
func (rcb *RodentCreateBulk) Save(ctx context.Context) ([]*Rodent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rodent, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RodentMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RodentCreateBulk) SaveX(ctx context.Context) []*Rodent {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RodentCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RodentCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
