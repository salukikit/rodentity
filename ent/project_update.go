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
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/operator"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/project"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/router"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetObjective sets the "objective" field.
func (pu *ProjectUpdate) SetObjective(s string) *ProjectUpdate {
	pu.mutation.SetObjective(s)
	return pu
}

// SetNillableObjective sets the "objective" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableObjective(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetObjective(*s)
	}
	return pu
}

// ClearObjective clears the value of the "objective" field.
func (pu *ProjectUpdate) ClearObjective() *ProjectUpdate {
	pu.mutation.ClearObjective()
	return pu
}

// SetEndDate sets the "end_date" field.
func (pu *ProjectUpdate) SetEndDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetEndDate(t)
	return pu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableEndDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetEndDate(*t)
	}
	return pu
}

// ClearEndDate clears the value of the "end_date" field.
func (pu *ProjectUpdate) ClearEndDate() *ProjectUpdate {
	pu.mutation.ClearEndDate()
	return pu
}

// SetStartDate sets the "start_date" field.
func (pu *ProjectUpdate) SetStartDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetStartDate(t)
	return pu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStartDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetStartDate(*t)
	}
	return pu
}

// ClearStartDate clears the value of the "start_date" field.
func (pu *ProjectUpdate) ClearStartDate() *ProjectUpdate {
	pu.mutation.ClearStartDate()
	return pu
}

// AddOperatorIDs adds the "operators" edge to the Operator entity by IDs.
func (pu *ProjectUpdate) AddOperatorIDs(ids ...xid.ID) *ProjectUpdate {
	pu.mutation.AddOperatorIDs(ids...)
	return pu
}

// AddOperators adds the "operators" edges to the Operator entity.
func (pu *ProjectUpdate) AddOperators(o ...*Operator) *ProjectUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddOperatorIDs(ids...)
}

// AddRodentIDs adds the "rodents" edge to the Rodent entity by IDs.
func (pu *ProjectUpdate) AddRodentIDs(ids ...xid.ID) *ProjectUpdate {
	pu.mutation.AddRodentIDs(ids...)
	return pu
}

// AddRodents adds the "rodents" edges to the Rodent entity.
func (pu *ProjectUpdate) AddRodents(r ...*Rodent) *ProjectUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRodentIDs(ids...)
}

// AddRouterIDs adds the "routers" edge to the Router entity by IDs.
func (pu *ProjectUpdate) AddRouterIDs(ids ...xid.ID) *ProjectUpdate {
	pu.mutation.AddRouterIDs(ids...)
	return pu
}

// AddRouters adds the "routers" edges to the Router entity.
func (pu *ProjectUpdate) AddRouters(r ...*Router) *ProjectUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRouterIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearOperators clears all "operators" edges to the Operator entity.
func (pu *ProjectUpdate) ClearOperators() *ProjectUpdate {
	pu.mutation.ClearOperators()
	return pu
}

// RemoveOperatorIDs removes the "operators" edge to Operator entities by IDs.
func (pu *ProjectUpdate) RemoveOperatorIDs(ids ...xid.ID) *ProjectUpdate {
	pu.mutation.RemoveOperatorIDs(ids...)
	return pu
}

// RemoveOperators removes "operators" edges to Operator entities.
func (pu *ProjectUpdate) RemoveOperators(o ...*Operator) *ProjectUpdate {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveOperatorIDs(ids...)
}

// ClearRodents clears all "rodents" edges to the Rodent entity.
func (pu *ProjectUpdate) ClearRodents() *ProjectUpdate {
	pu.mutation.ClearRodents()
	return pu
}

// RemoveRodentIDs removes the "rodents" edge to Rodent entities by IDs.
func (pu *ProjectUpdate) RemoveRodentIDs(ids ...xid.ID) *ProjectUpdate {
	pu.mutation.RemoveRodentIDs(ids...)
	return pu
}

// RemoveRodents removes "rodents" edges to Rodent entities.
func (pu *ProjectUpdate) RemoveRodents(r ...*Rodent) *ProjectUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRodentIDs(ids...)
}

// ClearRouters clears all "routers" edges to the Router entity.
func (pu *ProjectUpdate) ClearRouters() *ProjectUpdate {
	pu.mutation.ClearRouters()
	return pu
}

// RemoveRouterIDs removes the "routers" edge to Router entities by IDs.
func (pu *ProjectUpdate) RemoveRouterIDs(ids ...xid.ID) *ProjectUpdate {
	pu.mutation.RemoveRouterIDs(ids...)
	return pu
}

// RemoveRouters removes "routers" edges to Router entities.
func (pu *ProjectUpdate) RemoveRouters(r ...*Router) *ProjectUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRouterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ProjectMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Objective(); ok {
		_spec.SetField(project.FieldObjective, field.TypeString, value)
	}
	if pu.mutation.ObjectiveCleared() {
		_spec.ClearField(project.FieldObjective, field.TypeString)
	}
	if value, ok := pu.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if pu.mutation.EndDateCleared() {
		_spec.ClearField(project.FieldEndDate, field.TypeTime)
	}
	if value, ok := pu.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if pu.mutation.StartDateCleared() {
		_spec.ClearField(project.FieldStartDate, field.TypeTime)
	}
	if pu.mutation.OperatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.OperatorsTable,
			Columns: project.OperatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: operator.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedOperatorsIDs(); len(nodes) > 0 && !pu.mutation.OperatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.OperatorsTable,
			Columns: project.OperatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: operator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OperatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.OperatorsTable,
			Columns: project.OperatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: operator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.RodentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RodentsTable,
			Columns: []string{project.RodentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: rodent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRodentsIDs(); len(nodes) > 0 && !pu.mutation.RodentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RodentsTable,
			Columns: []string{project.RodentsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RodentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RodentsTable,
			Columns: []string{project.RodentsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.RoutersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoutersTable,
			Columns: []string{project.RoutersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: router.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRoutersIDs(); len(nodes) > 0 && !pu.mutation.RoutersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoutersTable,
			Columns: []string{project.RoutersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: router.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RoutersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoutersTable,
			Columns: []string{project.RoutersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: router.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetObjective sets the "objective" field.
func (puo *ProjectUpdateOne) SetObjective(s string) *ProjectUpdateOne {
	puo.mutation.SetObjective(s)
	return puo
}

// SetNillableObjective sets the "objective" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableObjective(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetObjective(*s)
	}
	return puo
}

// ClearObjective clears the value of the "objective" field.
func (puo *ProjectUpdateOne) ClearObjective() *ProjectUpdateOne {
	puo.mutation.ClearObjective()
	return puo
}

// SetEndDate sets the "end_date" field.
func (puo *ProjectUpdateOne) SetEndDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetEndDate(t)
	return puo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableEndDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetEndDate(*t)
	}
	return puo
}

// ClearEndDate clears the value of the "end_date" field.
func (puo *ProjectUpdateOne) ClearEndDate() *ProjectUpdateOne {
	puo.mutation.ClearEndDate()
	return puo
}

// SetStartDate sets the "start_date" field.
func (puo *ProjectUpdateOne) SetStartDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetStartDate(t)
	return puo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStartDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetStartDate(*t)
	}
	return puo
}

// ClearStartDate clears the value of the "start_date" field.
func (puo *ProjectUpdateOne) ClearStartDate() *ProjectUpdateOne {
	puo.mutation.ClearStartDate()
	return puo
}

// AddOperatorIDs adds the "operators" edge to the Operator entity by IDs.
func (puo *ProjectUpdateOne) AddOperatorIDs(ids ...xid.ID) *ProjectUpdateOne {
	puo.mutation.AddOperatorIDs(ids...)
	return puo
}

// AddOperators adds the "operators" edges to the Operator entity.
func (puo *ProjectUpdateOne) AddOperators(o ...*Operator) *ProjectUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddOperatorIDs(ids...)
}

// AddRodentIDs adds the "rodents" edge to the Rodent entity by IDs.
func (puo *ProjectUpdateOne) AddRodentIDs(ids ...xid.ID) *ProjectUpdateOne {
	puo.mutation.AddRodentIDs(ids...)
	return puo
}

// AddRodents adds the "rodents" edges to the Rodent entity.
func (puo *ProjectUpdateOne) AddRodents(r ...*Rodent) *ProjectUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRodentIDs(ids...)
}

// AddRouterIDs adds the "routers" edge to the Router entity by IDs.
func (puo *ProjectUpdateOne) AddRouterIDs(ids ...xid.ID) *ProjectUpdateOne {
	puo.mutation.AddRouterIDs(ids...)
	return puo
}

// AddRouters adds the "routers" edges to the Router entity.
func (puo *ProjectUpdateOne) AddRouters(r ...*Router) *ProjectUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRouterIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearOperators clears all "operators" edges to the Operator entity.
func (puo *ProjectUpdateOne) ClearOperators() *ProjectUpdateOne {
	puo.mutation.ClearOperators()
	return puo
}

// RemoveOperatorIDs removes the "operators" edge to Operator entities by IDs.
func (puo *ProjectUpdateOne) RemoveOperatorIDs(ids ...xid.ID) *ProjectUpdateOne {
	puo.mutation.RemoveOperatorIDs(ids...)
	return puo
}

// RemoveOperators removes "operators" edges to Operator entities.
func (puo *ProjectUpdateOne) RemoveOperators(o ...*Operator) *ProjectUpdateOne {
	ids := make([]xid.ID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveOperatorIDs(ids...)
}

// ClearRodents clears all "rodents" edges to the Rodent entity.
func (puo *ProjectUpdateOne) ClearRodents() *ProjectUpdateOne {
	puo.mutation.ClearRodents()
	return puo
}

// RemoveRodentIDs removes the "rodents" edge to Rodent entities by IDs.
func (puo *ProjectUpdateOne) RemoveRodentIDs(ids ...xid.ID) *ProjectUpdateOne {
	puo.mutation.RemoveRodentIDs(ids...)
	return puo
}

// RemoveRodents removes "rodents" edges to Rodent entities.
func (puo *ProjectUpdateOne) RemoveRodents(r ...*Rodent) *ProjectUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRodentIDs(ids...)
}

// ClearRouters clears all "routers" edges to the Router entity.
func (puo *ProjectUpdateOne) ClearRouters() *ProjectUpdateOne {
	puo.mutation.ClearRouters()
	return puo
}

// RemoveRouterIDs removes the "routers" edge to Router entities by IDs.
func (puo *ProjectUpdateOne) RemoveRouterIDs(ids ...xid.ID) *ProjectUpdateOne {
	puo.mutation.RemoveRouterIDs(ids...)
	return puo
}

// RemoveRouters removes "routers" edges to Router entities.
func (puo *ProjectUpdateOne) RemoveRouters(r ...*Router) *ProjectUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRouterIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks[*Project, ProjectMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Objective(); ok {
		_spec.SetField(project.FieldObjective, field.TypeString, value)
	}
	if puo.mutation.ObjectiveCleared() {
		_spec.ClearField(project.FieldObjective, field.TypeString)
	}
	if value, ok := puo.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if puo.mutation.EndDateCleared() {
		_spec.ClearField(project.FieldEndDate, field.TypeTime)
	}
	if value, ok := puo.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if puo.mutation.StartDateCleared() {
		_spec.ClearField(project.FieldStartDate, field.TypeTime)
	}
	if puo.mutation.OperatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.OperatorsTable,
			Columns: project.OperatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: operator.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedOperatorsIDs(); len(nodes) > 0 && !puo.mutation.OperatorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.OperatorsTable,
			Columns: project.OperatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: operator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OperatorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.OperatorsTable,
			Columns: project.OperatorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: operator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.RodentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RodentsTable,
			Columns: []string{project.RodentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: rodent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRodentsIDs(); len(nodes) > 0 && !puo.mutation.RodentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RodentsTable,
			Columns: []string{project.RodentsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RodentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RodentsTable,
			Columns: []string{project.RodentsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.RoutersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoutersTable,
			Columns: []string{project.RoutersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: router.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRoutersIDs(); len(nodes) > 0 && !puo.mutation.RoutersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoutersTable,
			Columns: []string{project.RoutersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: router.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RoutersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoutersTable,
			Columns: []string{project.RoutersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: router.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
