// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/project"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/router"
	"github.com/salukikit/rodentity/ent/task"
	"github.com/salukikit/rodentity/ent/user"
)

// RodentUpdate is the builder for updating Rodent entities.
type RodentUpdate struct {
	config
	hooks    []Hook
	mutation *RodentMutation
}

// Where appends a list predicates to the RodentUpdate builder.
func (ru *RodentUpdate) Where(ps ...predicate.Rodent) *RodentUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetXid sets the "xid" field.
func (ru *RodentUpdate) SetXid(s string) *RodentUpdate {
	ru.mutation.SetXid(s)
	return ru
}

// SetType sets the "type" field.
func (ru *RodentUpdate) SetType(s string) *RodentUpdate {
	ru.mutation.SetType(s)
	return ru
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ru *RodentUpdate) SetNillableType(s *string) *RodentUpdate {
	if s != nil {
		ru.SetType(*s)
	}
	return ru
}

// SetCodename sets the "codename" field.
func (ru *RodentUpdate) SetCodename(s string) *RodentUpdate {
	ru.mutation.SetCodename(s)
	return ru
}

// SetKey sets the "key" field.
func (ru *RodentUpdate) SetKey(s string) *RodentUpdate {
	ru.mutation.SetKey(s)
	return ru
}

// SetUsercontext sets the "usercontext" field.
func (ru *RodentUpdate) SetUsercontext(s string) *RodentUpdate {
	ru.mutation.SetUsercontext(s)
	return ru
}

// SetNillableUsercontext sets the "usercontext" field if the given value is not nil.
func (ru *RodentUpdate) SetNillableUsercontext(s *string) *RodentUpdate {
	if s != nil {
		ru.SetUsercontext(*s)
	}
	return ru
}

// ClearUsercontext clears the value of the "usercontext" field.
func (ru *RodentUpdate) ClearUsercontext() *RodentUpdate {
	ru.mutation.ClearUsercontext()
	return ru
}

// SetBeacontime sets the "beacontime" field.
func (ru *RodentUpdate) SetBeacontime(s string) *RodentUpdate {
	ru.mutation.SetBeacontime(s)
	return ru
}

// SetNillableBeacontime sets the "beacontime" field if the given value is not nil.
func (ru *RodentUpdate) SetNillableBeacontime(s *string) *RodentUpdate {
	if s != nil {
		ru.SetBeacontime(*s)
	}
	return ru
}

// ClearBeacontime clears the value of the "beacontime" field.
func (ru *RodentUpdate) ClearBeacontime() *RodentUpdate {
	ru.mutation.ClearBeacontime()
	return ru
}

// SetBurned sets the "burned" field.
func (ru *RodentUpdate) SetBurned(b bool) *RodentUpdate {
	ru.mutation.SetBurned(b)
	return ru
}

// SetNillableBurned sets the "burned" field if the given value is not nil.
func (ru *RodentUpdate) SetNillableBurned(b *bool) *RodentUpdate {
	if b != nil {
		ru.SetBurned(*b)
	}
	return ru
}

// SetAlive sets the "alive" field.
func (ru *RodentUpdate) SetAlive(b bool) *RodentUpdate {
	ru.mutation.SetAlive(b)
	return ru
}

// SetNillableAlive sets the "alive" field if the given value is not nil.
func (ru *RodentUpdate) SetNillableAlive(b *bool) *RodentUpdate {
	if b != nil {
		ru.SetAlive(*b)
	}
	return ru
}

// SetJoined sets the "joined" field.
func (ru *RodentUpdate) SetJoined(t time.Time) *RodentUpdate {
	ru.mutation.SetJoined(t)
	return ru
}

// SetLastseen sets the "lastseen" field.
func (ru *RodentUpdate) SetLastseen(t time.Time) *RodentUpdate {
	ru.mutation.SetLastseen(t)
	return ru
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (ru *RodentUpdate) SetDeviceID(id int) *RodentUpdate {
	ru.mutation.SetDeviceID(id)
	return ru
}

// SetNillableDeviceID sets the "device" edge to the Device entity by ID if the given value is not nil.
func (ru *RodentUpdate) SetNillableDeviceID(id *int) *RodentUpdate {
	if id != nil {
		ru = ru.SetDeviceID(*id)
	}
	return ru
}

// SetDevice sets the "device" edge to the Device entity.
func (ru *RodentUpdate) SetDevice(d *Device) *RodentUpdate {
	return ru.SetDeviceID(d.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ru *RodentUpdate) SetUserID(id int) *RodentUpdate {
	ru.mutation.SetUserID(id)
	return ru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ru *RodentUpdate) SetNillableUserID(id *int) *RodentUpdate {
	if id != nil {
		ru = ru.SetUserID(*id)
	}
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *RodentUpdate) SetUser(u *User) *RodentUpdate {
	return ru.SetUserID(u.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ru *RodentUpdate) SetProjectID(id int) *RodentUpdate {
	ru.mutation.SetProjectID(id)
	return ru
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (ru *RodentUpdate) SetNillableProjectID(id *int) *RodentUpdate {
	if id != nil {
		ru = ru.SetProjectID(*id)
	}
	return ru
}

// SetProject sets the "project" edge to the Project entity.
func (ru *RodentUpdate) SetProject(p *Project) *RodentUpdate {
	return ru.SetProjectID(p.ID)
}

// SetRouterID sets the "router" edge to the Router entity by ID.
func (ru *RodentUpdate) SetRouterID(id int) *RodentUpdate {
	ru.mutation.SetRouterID(id)
	return ru
}

// SetNillableRouterID sets the "router" edge to the Router entity by ID if the given value is not nil.
func (ru *RodentUpdate) SetNillableRouterID(id *int) *RodentUpdate {
	if id != nil {
		ru = ru.SetRouterID(*id)
	}
	return ru
}

// SetRouter sets the "router" edge to the Router entity.
func (ru *RodentUpdate) SetRouter(r *Router) *RodentUpdate {
	return ru.SetRouterID(r.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ru *RodentUpdate) AddTaskIDs(ids ...int) *RodentUpdate {
	ru.mutation.AddTaskIDs(ids...)
	return ru
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ru *RodentUpdate) AddTasks(t ...*Task) *RodentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.AddTaskIDs(ids...)
}

// AddLootIDs adds the "loot" edge to the Loot entity by IDs.
func (ru *RodentUpdate) AddLootIDs(ids ...int) *RodentUpdate {
	ru.mutation.AddLootIDs(ids...)
	return ru
}

// AddLoot adds the "loot" edges to the Loot entity.
func (ru *RodentUpdate) AddLoot(l ...*Loot) *RodentUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.AddLootIDs(ids...)
}

// Mutation returns the RodentMutation object of the builder.
func (ru *RodentUpdate) Mutation() *RodentMutation {
	return ru.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (ru *RodentUpdate) ClearDevice() *RodentUpdate {
	ru.mutation.ClearDevice()
	return ru
}

// ClearUser clears the "user" edge to the User entity.
func (ru *RodentUpdate) ClearUser() *RodentUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearProject clears the "project" edge to the Project entity.
func (ru *RodentUpdate) ClearProject() *RodentUpdate {
	ru.mutation.ClearProject()
	return ru
}

// ClearRouter clears the "router" edge to the Router entity.
func (ru *RodentUpdate) ClearRouter() *RodentUpdate {
	ru.mutation.ClearRouter()
	return ru
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (ru *RodentUpdate) ClearTasks() *RodentUpdate {
	ru.mutation.ClearTasks()
	return ru
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (ru *RodentUpdate) RemoveTaskIDs(ids ...int) *RodentUpdate {
	ru.mutation.RemoveTaskIDs(ids...)
	return ru
}

// RemoveTasks removes "tasks" edges to Task entities.
func (ru *RodentUpdate) RemoveTasks(t ...*Task) *RodentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.RemoveTaskIDs(ids...)
}

// ClearLoot clears all "loot" edges to the Loot entity.
func (ru *RodentUpdate) ClearLoot() *RodentUpdate {
	ru.mutation.ClearLoot()
	return ru
}

// RemoveLootIDs removes the "loot" edge to Loot entities by IDs.
func (ru *RodentUpdate) RemoveLootIDs(ids ...int) *RodentUpdate {
	ru.mutation.RemoveLootIDs(ids...)
	return ru
}

// RemoveLoot removes "loot" edges to Loot entities.
func (ru *RodentUpdate) RemoveLoot(l ...*Loot) *RodentUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.RemoveLootIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RodentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RodentMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RodentUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RodentUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RodentUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RodentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rodent.Table, rodent.Columns, sqlgraph.NewFieldSpec(rodent.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Xid(); ok {
		_spec.SetField(rodent.FieldXid, field.TypeString, value)
	}
	if value, ok := ru.mutation.GetType(); ok {
		_spec.SetField(rodent.FieldType, field.TypeString, value)
	}
	if value, ok := ru.mutation.Codename(); ok {
		_spec.SetField(rodent.FieldCodename, field.TypeString, value)
	}
	if value, ok := ru.mutation.Key(); ok {
		_spec.SetField(rodent.FieldKey, field.TypeString, value)
	}
	if value, ok := ru.mutation.Usercontext(); ok {
		_spec.SetField(rodent.FieldUsercontext, field.TypeString, value)
	}
	if ru.mutation.UsercontextCleared() {
		_spec.ClearField(rodent.FieldUsercontext, field.TypeString)
	}
	if value, ok := ru.mutation.Beacontime(); ok {
		_spec.SetField(rodent.FieldBeacontime, field.TypeString, value)
	}
	if ru.mutation.BeacontimeCleared() {
		_spec.ClearField(rodent.FieldBeacontime, field.TypeString)
	}
	if value, ok := ru.mutation.Burned(); ok {
		_spec.SetField(rodent.FieldBurned, field.TypeBool, value)
	}
	if value, ok := ru.mutation.Alive(); ok {
		_spec.SetField(rodent.FieldAlive, field.TypeBool, value)
	}
	if value, ok := ru.mutation.Joined(); ok {
		_spec.SetField(rodent.FieldJoined, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Lastseen(); ok {
		_spec.SetField(rodent.FieldLastseen, field.TypeTime, value)
	}
	if ru.mutation.DeviceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DeviceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RouterCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RouterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedTasksIDs(); len(nodes) > 0 && !ru.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.LootCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedLootIDs(); len(nodes) > 0 && !ru.mutation.LootCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LootIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rodent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RodentUpdateOne is the builder for updating a single Rodent entity.
type RodentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RodentMutation
}

// SetXid sets the "xid" field.
func (ruo *RodentUpdateOne) SetXid(s string) *RodentUpdateOne {
	ruo.mutation.SetXid(s)
	return ruo
}

// SetType sets the "type" field.
func (ruo *RodentUpdateOne) SetType(s string) *RodentUpdateOne {
	ruo.mutation.SetType(s)
	return ruo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableType(s *string) *RodentUpdateOne {
	if s != nil {
		ruo.SetType(*s)
	}
	return ruo
}

// SetCodename sets the "codename" field.
func (ruo *RodentUpdateOne) SetCodename(s string) *RodentUpdateOne {
	ruo.mutation.SetCodename(s)
	return ruo
}

// SetKey sets the "key" field.
func (ruo *RodentUpdateOne) SetKey(s string) *RodentUpdateOne {
	ruo.mutation.SetKey(s)
	return ruo
}

// SetUsercontext sets the "usercontext" field.
func (ruo *RodentUpdateOne) SetUsercontext(s string) *RodentUpdateOne {
	ruo.mutation.SetUsercontext(s)
	return ruo
}

// SetNillableUsercontext sets the "usercontext" field if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableUsercontext(s *string) *RodentUpdateOne {
	if s != nil {
		ruo.SetUsercontext(*s)
	}
	return ruo
}

// ClearUsercontext clears the value of the "usercontext" field.
func (ruo *RodentUpdateOne) ClearUsercontext() *RodentUpdateOne {
	ruo.mutation.ClearUsercontext()
	return ruo
}

// SetBeacontime sets the "beacontime" field.
func (ruo *RodentUpdateOne) SetBeacontime(s string) *RodentUpdateOne {
	ruo.mutation.SetBeacontime(s)
	return ruo
}

// SetNillableBeacontime sets the "beacontime" field if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableBeacontime(s *string) *RodentUpdateOne {
	if s != nil {
		ruo.SetBeacontime(*s)
	}
	return ruo
}

// ClearBeacontime clears the value of the "beacontime" field.
func (ruo *RodentUpdateOne) ClearBeacontime() *RodentUpdateOne {
	ruo.mutation.ClearBeacontime()
	return ruo
}

// SetBurned sets the "burned" field.
func (ruo *RodentUpdateOne) SetBurned(b bool) *RodentUpdateOne {
	ruo.mutation.SetBurned(b)
	return ruo
}

// SetNillableBurned sets the "burned" field if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableBurned(b *bool) *RodentUpdateOne {
	if b != nil {
		ruo.SetBurned(*b)
	}
	return ruo
}

// SetAlive sets the "alive" field.
func (ruo *RodentUpdateOne) SetAlive(b bool) *RodentUpdateOne {
	ruo.mutation.SetAlive(b)
	return ruo
}

// SetNillableAlive sets the "alive" field if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableAlive(b *bool) *RodentUpdateOne {
	if b != nil {
		ruo.SetAlive(*b)
	}
	return ruo
}

// SetJoined sets the "joined" field.
func (ruo *RodentUpdateOne) SetJoined(t time.Time) *RodentUpdateOne {
	ruo.mutation.SetJoined(t)
	return ruo
}

// SetLastseen sets the "lastseen" field.
func (ruo *RodentUpdateOne) SetLastseen(t time.Time) *RodentUpdateOne {
	ruo.mutation.SetLastseen(t)
	return ruo
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (ruo *RodentUpdateOne) SetDeviceID(id int) *RodentUpdateOne {
	ruo.mutation.SetDeviceID(id)
	return ruo
}

// SetNillableDeviceID sets the "device" edge to the Device entity by ID if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableDeviceID(id *int) *RodentUpdateOne {
	if id != nil {
		ruo = ruo.SetDeviceID(*id)
	}
	return ruo
}

// SetDevice sets the "device" edge to the Device entity.
func (ruo *RodentUpdateOne) SetDevice(d *Device) *RodentUpdateOne {
	return ruo.SetDeviceID(d.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ruo *RodentUpdateOne) SetUserID(id int) *RodentUpdateOne {
	ruo.mutation.SetUserID(id)
	return ruo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableUserID(id *int) *RodentUpdateOne {
	if id != nil {
		ruo = ruo.SetUserID(*id)
	}
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *RodentUpdateOne) SetUser(u *User) *RodentUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ruo *RodentUpdateOne) SetProjectID(id int) *RodentUpdateOne {
	ruo.mutation.SetProjectID(id)
	return ruo
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableProjectID(id *int) *RodentUpdateOne {
	if id != nil {
		ruo = ruo.SetProjectID(*id)
	}
	return ruo
}

// SetProject sets the "project" edge to the Project entity.
func (ruo *RodentUpdateOne) SetProject(p *Project) *RodentUpdateOne {
	return ruo.SetProjectID(p.ID)
}

// SetRouterID sets the "router" edge to the Router entity by ID.
func (ruo *RodentUpdateOne) SetRouterID(id int) *RodentUpdateOne {
	ruo.mutation.SetRouterID(id)
	return ruo
}

// SetNillableRouterID sets the "router" edge to the Router entity by ID if the given value is not nil.
func (ruo *RodentUpdateOne) SetNillableRouterID(id *int) *RodentUpdateOne {
	if id != nil {
		ruo = ruo.SetRouterID(*id)
	}
	return ruo
}

// SetRouter sets the "router" edge to the Router entity.
func (ruo *RodentUpdateOne) SetRouter(r *Router) *RodentUpdateOne {
	return ruo.SetRouterID(r.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ruo *RodentUpdateOne) AddTaskIDs(ids ...int) *RodentUpdateOne {
	ruo.mutation.AddTaskIDs(ids...)
	return ruo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ruo *RodentUpdateOne) AddTasks(t ...*Task) *RodentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.AddTaskIDs(ids...)
}

// AddLootIDs adds the "loot" edge to the Loot entity by IDs.
func (ruo *RodentUpdateOne) AddLootIDs(ids ...int) *RodentUpdateOne {
	ruo.mutation.AddLootIDs(ids...)
	return ruo
}

// AddLoot adds the "loot" edges to the Loot entity.
func (ruo *RodentUpdateOne) AddLoot(l ...*Loot) *RodentUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.AddLootIDs(ids...)
}

// Mutation returns the RodentMutation object of the builder.
func (ruo *RodentUpdateOne) Mutation() *RodentMutation {
	return ruo.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (ruo *RodentUpdateOne) ClearDevice() *RodentUpdateOne {
	ruo.mutation.ClearDevice()
	return ruo
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *RodentUpdateOne) ClearUser() *RodentUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearProject clears the "project" edge to the Project entity.
func (ruo *RodentUpdateOne) ClearProject() *RodentUpdateOne {
	ruo.mutation.ClearProject()
	return ruo
}

// ClearRouter clears the "router" edge to the Router entity.
func (ruo *RodentUpdateOne) ClearRouter() *RodentUpdateOne {
	ruo.mutation.ClearRouter()
	return ruo
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (ruo *RodentUpdateOne) ClearTasks() *RodentUpdateOne {
	ruo.mutation.ClearTasks()
	return ruo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (ruo *RodentUpdateOne) RemoveTaskIDs(ids ...int) *RodentUpdateOne {
	ruo.mutation.RemoveTaskIDs(ids...)
	return ruo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (ruo *RodentUpdateOne) RemoveTasks(t ...*Task) *RodentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.RemoveTaskIDs(ids...)
}

// ClearLoot clears all "loot" edges to the Loot entity.
func (ruo *RodentUpdateOne) ClearLoot() *RodentUpdateOne {
	ruo.mutation.ClearLoot()
	return ruo
}

// RemoveLootIDs removes the "loot" edge to Loot entities by IDs.
func (ruo *RodentUpdateOne) RemoveLootIDs(ids ...int) *RodentUpdateOne {
	ruo.mutation.RemoveLootIDs(ids...)
	return ruo
}

// RemoveLoot removes "loot" edges to Loot entities.
func (ruo *RodentUpdateOne) RemoveLoot(l ...*Loot) *RodentUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.RemoveLootIDs(ids...)
}

// Where appends a list predicates to the RodentUpdate builder.
func (ruo *RodentUpdateOne) Where(ps ...predicate.Rodent) *RodentUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RodentUpdateOne) Select(field string, fields ...string) *RodentUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Rodent entity.
func (ruo *RodentUpdateOne) Save(ctx context.Context) (*Rodent, error) {
	return withHooks[*Rodent, RodentMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RodentUpdateOne) SaveX(ctx context.Context) *Rodent {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RodentUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RodentUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RodentUpdateOne) sqlSave(ctx context.Context) (_node *Rodent, err error) {
	_spec := sqlgraph.NewUpdateSpec(rodent.Table, rodent.Columns, sqlgraph.NewFieldSpec(rodent.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rodent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rodent.FieldID)
		for _, f := range fields {
			if !rodent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rodent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Xid(); ok {
		_spec.SetField(rodent.FieldXid, field.TypeString, value)
	}
	if value, ok := ruo.mutation.GetType(); ok {
		_spec.SetField(rodent.FieldType, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Codename(); ok {
		_spec.SetField(rodent.FieldCodename, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Key(); ok {
		_spec.SetField(rodent.FieldKey, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Usercontext(); ok {
		_spec.SetField(rodent.FieldUsercontext, field.TypeString, value)
	}
	if ruo.mutation.UsercontextCleared() {
		_spec.ClearField(rodent.FieldUsercontext, field.TypeString)
	}
	if value, ok := ruo.mutation.Beacontime(); ok {
		_spec.SetField(rodent.FieldBeacontime, field.TypeString, value)
	}
	if ruo.mutation.BeacontimeCleared() {
		_spec.ClearField(rodent.FieldBeacontime, field.TypeString)
	}
	if value, ok := ruo.mutation.Burned(); ok {
		_spec.SetField(rodent.FieldBurned, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.Alive(); ok {
		_spec.SetField(rodent.FieldAlive, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.Joined(); ok {
		_spec.SetField(rodent.FieldJoined, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Lastseen(); ok {
		_spec.SetField(rodent.FieldLastseen, field.TypeTime, value)
	}
	if ruo.mutation.DeviceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DeviceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RouterCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RouterIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !ruo.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.LootCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedLootIDs(); len(nodes) > 0 && !ruo.mutation.LootCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LootIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rodent{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rodent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
