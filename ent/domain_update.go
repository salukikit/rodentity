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

// DomainUpdate is the builder for updating Domain entities.
type DomainUpdate struct {
	config
	hooks    []Hook
	mutation *DomainMutation
}

// Where appends a list predicates to the DomainUpdate builder.
func (du *DomainUpdate) Where(ps ...predicate.Domain) *DomainUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DomainUpdate) SetName(s string) *DomainUpdate {
	du.mutation.SetName(s)
	return du
}

// SetAD sets the "AD" field.
func (du *DomainUpdate) SetAD(b bool) *DomainUpdate {
	du.mutation.SetAD(b)
	return du
}

// SetNillableAD sets the "AD" field if the given value is not nil.
func (du *DomainUpdate) SetNillableAD(b *bool) *DomainUpdate {
	if b != nil {
		du.SetAD(*b)
	}
	return du
}

// SetOwned sets the "owned" field.
func (du *DomainUpdate) SetOwned(b bool) *DomainUpdate {
	du.mutation.SetOwned(b)
	return du
}

// SetNillableOwned sets the "owned" field if the given value is not nil.
func (du *DomainUpdate) SetNillableOwned(b *bool) *DomainUpdate {
	if b != nil {
		du.SetOwned(*b)
	}
	return du
}

// SetCloud sets the "cloud" field.
func (du *DomainUpdate) SetCloud(s string) *DomainUpdate {
	du.mutation.SetCloud(s)
	return du
}

// SetNillableCloud sets the "cloud" field if the given value is not nil.
func (du *DomainUpdate) SetNillableCloud(s *string) *DomainUpdate {
	if s != nil {
		du.SetCloud(*s)
	}
	return du
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (du *DomainUpdate) AddDeviceIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.AddDeviceIDs(ids...)
	return du
}

// AddDevices adds the "devices" edges to the Device entity.
func (du *DomainUpdate) AddDevices(d ...*Device) *DomainUpdate {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDeviceIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (du *DomainUpdate) AddUserIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.AddUserIDs(ids...)
	return du
}

// AddUsers adds the "users" edges to the User entity.
func (du *DomainUpdate) AddUsers(u ...*User) *DomainUpdate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddUserIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (du *DomainUpdate) AddGroupIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.AddGroupIDs(ids...)
	return du
}

// AddGroups adds the "groups" edges to the Group entity.
func (du *DomainUpdate) AddGroups(g ...*Group) *DomainUpdate {
	ids := make([]xid.ID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return du.AddGroupIDs(ids...)
}

// AddChilddomainIDs adds the "childdomains" edge to the Domain entity by IDs.
func (du *DomainUpdate) AddChilddomainIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.AddChilddomainIDs(ids...)
	return du
}

// AddChilddomains adds the "childdomains" edges to the Domain entity.
func (du *DomainUpdate) AddChilddomains(d ...*Domain) *DomainUpdate {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddChilddomainIDs(ids...)
}

// SetParentdomainID sets the "parentdomain" edge to the Domain entity by ID.
func (du *DomainUpdate) SetParentdomainID(id xid.ID) *DomainUpdate {
	du.mutation.SetParentdomainID(id)
	return du
}

// SetNillableParentdomainID sets the "parentdomain" edge to the Domain entity by ID if the given value is not nil.
func (du *DomainUpdate) SetNillableParentdomainID(id *xid.ID) *DomainUpdate {
	if id != nil {
		du = du.SetParentdomainID(*id)
	}
	return du
}

// SetParentdomain sets the "parentdomain" edge to the Domain entity.
func (du *DomainUpdate) SetParentdomain(d *Domain) *DomainUpdate {
	return du.SetParentdomainID(d.ID)
}

// Mutation returns the DomainMutation object of the builder.
func (du *DomainUpdate) Mutation() *DomainMutation {
	return du.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (du *DomainUpdate) ClearDevices() *DomainUpdate {
	du.mutation.ClearDevices()
	return du
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (du *DomainUpdate) RemoveDeviceIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.RemoveDeviceIDs(ids...)
	return du
}

// RemoveDevices removes "devices" edges to Device entities.
func (du *DomainUpdate) RemoveDevices(d ...*Device) *DomainUpdate {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDeviceIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (du *DomainUpdate) ClearUsers() *DomainUpdate {
	du.mutation.ClearUsers()
	return du
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (du *DomainUpdate) RemoveUserIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.RemoveUserIDs(ids...)
	return du
}

// RemoveUsers removes "users" edges to User entities.
func (du *DomainUpdate) RemoveUsers(u ...*User) *DomainUpdate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveUserIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (du *DomainUpdate) ClearGroups() *DomainUpdate {
	du.mutation.ClearGroups()
	return du
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (du *DomainUpdate) RemoveGroupIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.RemoveGroupIDs(ids...)
	return du
}

// RemoveGroups removes "groups" edges to Group entities.
func (du *DomainUpdate) RemoveGroups(g ...*Group) *DomainUpdate {
	ids := make([]xid.ID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return du.RemoveGroupIDs(ids...)
}

// ClearChilddomains clears all "childdomains" edges to the Domain entity.
func (du *DomainUpdate) ClearChilddomains() *DomainUpdate {
	du.mutation.ClearChilddomains()
	return du
}

// RemoveChilddomainIDs removes the "childdomains" edge to Domain entities by IDs.
func (du *DomainUpdate) RemoveChilddomainIDs(ids ...xid.ID) *DomainUpdate {
	du.mutation.RemoveChilddomainIDs(ids...)
	return du
}

// RemoveChilddomains removes "childdomains" edges to Domain entities.
func (du *DomainUpdate) RemoveChilddomains(d ...*Domain) *DomainUpdate {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveChilddomainIDs(ids...)
}

// ClearParentdomain clears the "parentdomain" edge to the Domain entity.
func (du *DomainUpdate) ClearParentdomain() *DomainUpdate {
	du.mutation.ClearParentdomain()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DomainUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, DomainMutation](ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DomainUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DomainUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DomainUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DomainUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(domain.Table, domain.Columns, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeString))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(domain.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.AD(); ok {
		_spec.SetField(domain.FieldAD, field.TypeBool, value)
	}
	if value, ok := du.mutation.Owned(); ok {
		_spec.SetField(domain.FieldOwned, field.TypeBool, value)
	}
	if value, ok := du.mutation.Cloud(); ok {
		_spec.SetField(domain.FieldCloud, field.TypeString, value)
	}
	if du.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
	if nodes := du.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !du.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
	if nodes := du.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
	if du.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
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
	if nodes := du.mutation.RemovedUsersIDs(); len(nodes) > 0 && !du.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
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
	if nodes := du.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
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
	if du.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !du.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ChilddomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedChilddomainsIDs(); len(nodes) > 0 && !du.mutation.ChilddomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ChilddomainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
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
	if du.mutation.ParentdomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   domain.ParentdomainTable,
			Columns: []string{domain.ParentdomainColumn},
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
	if nodes := du.mutation.ParentdomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   domain.ParentdomainTable,
			Columns: []string{domain.ParentdomainColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DomainUpdateOne is the builder for updating a single Domain entity.
type DomainUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DomainMutation
}

// SetName sets the "name" field.
func (duo *DomainUpdateOne) SetName(s string) *DomainUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetAD sets the "AD" field.
func (duo *DomainUpdateOne) SetAD(b bool) *DomainUpdateOne {
	duo.mutation.SetAD(b)
	return duo
}

// SetNillableAD sets the "AD" field if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableAD(b *bool) *DomainUpdateOne {
	if b != nil {
		duo.SetAD(*b)
	}
	return duo
}

// SetOwned sets the "owned" field.
func (duo *DomainUpdateOne) SetOwned(b bool) *DomainUpdateOne {
	duo.mutation.SetOwned(b)
	return duo
}

// SetNillableOwned sets the "owned" field if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableOwned(b *bool) *DomainUpdateOne {
	if b != nil {
		duo.SetOwned(*b)
	}
	return duo
}

// SetCloud sets the "cloud" field.
func (duo *DomainUpdateOne) SetCloud(s string) *DomainUpdateOne {
	duo.mutation.SetCloud(s)
	return duo
}

// SetNillableCloud sets the "cloud" field if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableCloud(s *string) *DomainUpdateOne {
	if s != nil {
		duo.SetCloud(*s)
	}
	return duo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (duo *DomainUpdateOne) AddDeviceIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.AddDeviceIDs(ids...)
	return duo
}

// AddDevices adds the "devices" edges to the Device entity.
func (duo *DomainUpdateOne) AddDevices(d ...*Device) *DomainUpdateOne {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDeviceIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (duo *DomainUpdateOne) AddUserIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.AddUserIDs(ids...)
	return duo
}

// AddUsers adds the "users" edges to the User entity.
func (duo *DomainUpdateOne) AddUsers(u ...*User) *DomainUpdateOne {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddUserIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (duo *DomainUpdateOne) AddGroupIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.AddGroupIDs(ids...)
	return duo
}

// AddGroups adds the "groups" edges to the Group entity.
func (duo *DomainUpdateOne) AddGroups(g ...*Group) *DomainUpdateOne {
	ids := make([]xid.ID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return duo.AddGroupIDs(ids...)
}

// AddChilddomainIDs adds the "childdomains" edge to the Domain entity by IDs.
func (duo *DomainUpdateOne) AddChilddomainIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.AddChilddomainIDs(ids...)
	return duo
}

// AddChilddomains adds the "childdomains" edges to the Domain entity.
func (duo *DomainUpdateOne) AddChilddomains(d ...*Domain) *DomainUpdateOne {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddChilddomainIDs(ids...)
}

// SetParentdomainID sets the "parentdomain" edge to the Domain entity by ID.
func (duo *DomainUpdateOne) SetParentdomainID(id xid.ID) *DomainUpdateOne {
	duo.mutation.SetParentdomainID(id)
	return duo
}

// SetNillableParentdomainID sets the "parentdomain" edge to the Domain entity by ID if the given value is not nil.
func (duo *DomainUpdateOne) SetNillableParentdomainID(id *xid.ID) *DomainUpdateOne {
	if id != nil {
		duo = duo.SetParentdomainID(*id)
	}
	return duo
}

// SetParentdomain sets the "parentdomain" edge to the Domain entity.
func (duo *DomainUpdateOne) SetParentdomain(d *Domain) *DomainUpdateOne {
	return duo.SetParentdomainID(d.ID)
}

// Mutation returns the DomainMutation object of the builder.
func (duo *DomainUpdateOne) Mutation() *DomainMutation {
	return duo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (duo *DomainUpdateOne) ClearDevices() *DomainUpdateOne {
	duo.mutation.ClearDevices()
	return duo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (duo *DomainUpdateOne) RemoveDeviceIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.RemoveDeviceIDs(ids...)
	return duo
}

// RemoveDevices removes "devices" edges to Device entities.
func (duo *DomainUpdateOne) RemoveDevices(d ...*Device) *DomainUpdateOne {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDeviceIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (duo *DomainUpdateOne) ClearUsers() *DomainUpdateOne {
	duo.mutation.ClearUsers()
	return duo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (duo *DomainUpdateOne) RemoveUserIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.RemoveUserIDs(ids...)
	return duo
}

// RemoveUsers removes "users" edges to User entities.
func (duo *DomainUpdateOne) RemoveUsers(u ...*User) *DomainUpdateOne {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveUserIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (duo *DomainUpdateOne) ClearGroups() *DomainUpdateOne {
	duo.mutation.ClearGroups()
	return duo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (duo *DomainUpdateOne) RemoveGroupIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.RemoveGroupIDs(ids...)
	return duo
}

// RemoveGroups removes "groups" edges to Group entities.
func (duo *DomainUpdateOne) RemoveGroups(g ...*Group) *DomainUpdateOne {
	ids := make([]xid.ID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return duo.RemoveGroupIDs(ids...)
}

// ClearChilddomains clears all "childdomains" edges to the Domain entity.
func (duo *DomainUpdateOne) ClearChilddomains() *DomainUpdateOne {
	duo.mutation.ClearChilddomains()
	return duo
}

// RemoveChilddomainIDs removes the "childdomains" edge to Domain entities by IDs.
func (duo *DomainUpdateOne) RemoveChilddomainIDs(ids ...xid.ID) *DomainUpdateOne {
	duo.mutation.RemoveChilddomainIDs(ids...)
	return duo
}

// RemoveChilddomains removes "childdomains" edges to Domain entities.
func (duo *DomainUpdateOne) RemoveChilddomains(d ...*Domain) *DomainUpdateOne {
	ids := make([]xid.ID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveChilddomainIDs(ids...)
}

// ClearParentdomain clears the "parentdomain" edge to the Domain entity.
func (duo *DomainUpdateOne) ClearParentdomain() *DomainUpdateOne {
	duo.mutation.ClearParentdomain()
	return duo
}

// Where appends a list predicates to the DomainUpdate builder.
func (duo *DomainUpdateOne) Where(ps ...predicate.Domain) *DomainUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DomainUpdateOne) Select(field string, fields ...string) *DomainUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Domain entity.
func (duo *DomainUpdateOne) Save(ctx context.Context) (*Domain, error) {
	return withHooks[*Domain, DomainMutation](ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DomainUpdateOne) SaveX(ctx context.Context) *Domain {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DomainUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DomainUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DomainUpdateOne) sqlSave(ctx context.Context) (_node *Domain, err error) {
	_spec := sqlgraph.NewUpdateSpec(domain.Table, domain.Columns, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeString))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Domain.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, domain.FieldID)
		for _, f := range fields {
			if !domain.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != domain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(domain.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.AD(); ok {
		_spec.SetField(domain.FieldAD, field.TypeBool, value)
	}
	if value, ok := duo.mutation.Owned(); ok {
		_spec.SetField(domain.FieldOwned, field.TypeBool, value)
	}
	if value, ok := duo.mutation.Cloud(); ok {
		_spec.SetField(domain.FieldCloud, field.TypeString, value)
	}
	if duo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
	if nodes := duo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !duo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
	if nodes := duo.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.DevicesTable,
			Columns: []string{domain.DevicesColumn},
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
	if duo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
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
	if nodes := duo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !duo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
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
	if nodes := duo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: []string{domain.UsersColumn},
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
	if duo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !duo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.GroupsTable,
			Columns: []string{domain.GroupsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ChilddomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: domain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedChilddomainsIDs(); len(nodes) > 0 && !duo.mutation.ChilddomainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ChilddomainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.ChilddomainsTable,
			Columns: []string{domain.ChilddomainsColumn},
			Bidi:    true,
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
	if duo.mutation.ParentdomainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   domain.ParentdomainTable,
			Columns: []string{domain.ParentdomainColumn},
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
	if nodes := duo.mutation.ParentdomainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   domain.ParentdomainTable,
			Columns: []string{domain.ParentdomainColumn},
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
	_node = &Domain{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
