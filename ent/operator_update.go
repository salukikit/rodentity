// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/salukikit/rodentity/ent/operator"
	"github.com/salukikit/rodentity/ent/predicate"
	"github.com/salukikit/rodentity/ent/project"
	"github.com/salukikit/rodentity/ent/task"
)

// OperatorUpdate is the builder for updating Operator entities.
type OperatorUpdate struct {
	config
	hooks    []Hook
	mutation *OperatorMutation
}

// Where appends a list predicates to the OperatorUpdate builder.
func (ou *OperatorUpdate) Where(ps ...predicate.Operator) *OperatorUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUsername sets the "username" field.
func (ou *OperatorUpdate) SetUsername(s string) *OperatorUpdate {
	ou.mutation.SetUsername(s)
	return ou
}

// SetPrivkey sets the "privkey" field.
func (ou *OperatorUpdate) SetPrivkey(b []byte) *OperatorUpdate {
	ou.mutation.SetPrivkey(b)
	return ou
}

// ClearPrivkey clears the value of the "privkey" field.
func (ou *OperatorUpdate) ClearPrivkey() *OperatorUpdate {
	ou.mutation.ClearPrivkey()
	return ou
}

// SetCert sets the "cert" field.
func (ou *OperatorUpdate) SetCert(b []byte) *OperatorUpdate {
	ou.mutation.SetCert(b)
	return ou
}

// ClearCert clears the value of the "cert" field.
func (ou *OperatorUpdate) ClearCert() *OperatorUpdate {
	ou.mutation.ClearCert()
	return ou
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ou *OperatorUpdate) AddProjectIDs(ids ...int) *OperatorUpdate {
	ou.mutation.AddProjectIDs(ids...)
	return ou
}

// AddProjects adds the "projects" edges to the Project entity.
func (ou *OperatorUpdate) AddProjects(p ...*Project) *OperatorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddProjectIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ou *OperatorUpdate) AddTaskIDs(ids ...int) *OperatorUpdate {
	ou.mutation.AddTaskIDs(ids...)
	return ou
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ou *OperatorUpdate) AddTasks(t ...*Task) *OperatorUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ou.AddTaskIDs(ids...)
}

// Mutation returns the OperatorMutation object of the builder.
func (ou *OperatorUpdate) Mutation() *OperatorMutation {
	return ou.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (ou *OperatorUpdate) ClearProjects() *OperatorUpdate {
	ou.mutation.ClearProjects()
	return ou
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (ou *OperatorUpdate) RemoveProjectIDs(ids ...int) *OperatorUpdate {
	ou.mutation.RemoveProjectIDs(ids...)
	return ou
}

// RemoveProjects removes "projects" edges to Project entities.
func (ou *OperatorUpdate) RemoveProjects(p ...*Project) *OperatorUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemoveProjectIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (ou *OperatorUpdate) ClearTasks() *OperatorUpdate {
	ou.mutation.ClearTasks()
	return ou
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (ou *OperatorUpdate) RemoveTaskIDs(ids ...int) *OperatorUpdate {
	ou.mutation.RemoveTaskIDs(ids...)
	return ou
}

// RemoveTasks removes "tasks" edges to Task entities.
func (ou *OperatorUpdate) RemoveTasks(t ...*Task) *OperatorUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ou.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OperatorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OperatorMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OperatorUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OperatorUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OperatorUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OperatorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(operator.Table, operator.Columns, sqlgraph.NewFieldSpec(operator.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Username(); ok {
		_spec.SetField(operator.FieldUsername, field.TypeString, value)
	}
	if value, ok := ou.mutation.Privkey(); ok {
		_spec.SetField(operator.FieldPrivkey, field.TypeBytes, value)
	}
	if ou.mutation.PrivkeyCleared() {
		_spec.ClearField(operator.FieldPrivkey, field.TypeBytes)
	}
	if value, ok := ou.mutation.Cert(); ok {
		_spec.SetField(operator.FieldCert, field.TypeBytes, value)
	}
	if ou.mutation.CertCleared() {
		_spec.ClearField(operator.FieldCert, field.TypeBytes)
	}
	if ou.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   operator.ProjectsTable,
			Columns: operator.ProjectsPrimaryKey,
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
	if nodes := ou.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !ou.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   operator.ProjectsTable,
			Columns: operator.ProjectsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   operator.ProjectsTable,
			Columns: operator.ProjectsPrimaryKey,
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
	if ou.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operator.TasksTable,
			Columns: []string{operator.TasksColumn},
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
	if nodes := ou.mutation.RemovedTasksIDs(); len(nodes) > 0 && !ou.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operator.TasksTable,
			Columns: []string{operator.TasksColumn},
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
	if nodes := ou.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operator.TasksTable,
			Columns: []string{operator.TasksColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{operator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OperatorUpdateOne is the builder for updating a single Operator entity.
type OperatorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OperatorMutation
}

// SetUsername sets the "username" field.
func (ouo *OperatorUpdateOne) SetUsername(s string) *OperatorUpdateOne {
	ouo.mutation.SetUsername(s)
	return ouo
}

// SetPrivkey sets the "privkey" field.
func (ouo *OperatorUpdateOne) SetPrivkey(b []byte) *OperatorUpdateOne {
	ouo.mutation.SetPrivkey(b)
	return ouo
}

// ClearPrivkey clears the value of the "privkey" field.
func (ouo *OperatorUpdateOne) ClearPrivkey() *OperatorUpdateOne {
	ouo.mutation.ClearPrivkey()
	return ouo
}

// SetCert sets the "cert" field.
func (ouo *OperatorUpdateOne) SetCert(b []byte) *OperatorUpdateOne {
	ouo.mutation.SetCert(b)
	return ouo
}

// ClearCert clears the value of the "cert" field.
func (ouo *OperatorUpdateOne) ClearCert() *OperatorUpdateOne {
	ouo.mutation.ClearCert()
	return ouo
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (ouo *OperatorUpdateOne) AddProjectIDs(ids ...int) *OperatorUpdateOne {
	ouo.mutation.AddProjectIDs(ids...)
	return ouo
}

// AddProjects adds the "projects" edges to the Project entity.
func (ouo *OperatorUpdateOne) AddProjects(p ...*Project) *OperatorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddProjectIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ouo *OperatorUpdateOne) AddTaskIDs(ids ...int) *OperatorUpdateOne {
	ouo.mutation.AddTaskIDs(ids...)
	return ouo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ouo *OperatorUpdateOne) AddTasks(t ...*Task) *OperatorUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ouo.AddTaskIDs(ids...)
}

// Mutation returns the OperatorMutation object of the builder.
func (ouo *OperatorUpdateOne) Mutation() *OperatorMutation {
	return ouo.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (ouo *OperatorUpdateOne) ClearProjects() *OperatorUpdateOne {
	ouo.mutation.ClearProjects()
	return ouo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (ouo *OperatorUpdateOne) RemoveProjectIDs(ids ...int) *OperatorUpdateOne {
	ouo.mutation.RemoveProjectIDs(ids...)
	return ouo
}

// RemoveProjects removes "projects" edges to Project entities.
func (ouo *OperatorUpdateOne) RemoveProjects(p ...*Project) *OperatorUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemoveProjectIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (ouo *OperatorUpdateOne) ClearTasks() *OperatorUpdateOne {
	ouo.mutation.ClearTasks()
	return ouo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (ouo *OperatorUpdateOne) RemoveTaskIDs(ids ...int) *OperatorUpdateOne {
	ouo.mutation.RemoveTaskIDs(ids...)
	return ouo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (ouo *OperatorUpdateOne) RemoveTasks(t ...*Task) *OperatorUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ouo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the OperatorUpdate builder.
func (ouo *OperatorUpdateOne) Where(ps ...predicate.Operator) *OperatorUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OperatorUpdateOne) Select(field string, fields ...string) *OperatorUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Operator entity.
func (ouo *OperatorUpdateOne) Save(ctx context.Context) (*Operator, error) {
	return withHooks[*Operator, OperatorMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OperatorUpdateOne) SaveX(ctx context.Context) *Operator {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OperatorUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OperatorUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OperatorUpdateOne) sqlSave(ctx context.Context) (_node *Operator, err error) {
	_spec := sqlgraph.NewUpdateSpec(operator.Table, operator.Columns, sqlgraph.NewFieldSpec(operator.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Operator.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, operator.FieldID)
		for _, f := range fields {
			if !operator.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != operator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Username(); ok {
		_spec.SetField(operator.FieldUsername, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Privkey(); ok {
		_spec.SetField(operator.FieldPrivkey, field.TypeBytes, value)
	}
	if ouo.mutation.PrivkeyCleared() {
		_spec.ClearField(operator.FieldPrivkey, field.TypeBytes)
	}
	if value, ok := ouo.mutation.Cert(); ok {
		_spec.SetField(operator.FieldCert, field.TypeBytes, value)
	}
	if ouo.mutation.CertCleared() {
		_spec.ClearField(operator.FieldCert, field.TypeBytes)
	}
	if ouo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   operator.ProjectsTable,
			Columns: operator.ProjectsPrimaryKey,
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
	if nodes := ouo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !ouo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   operator.ProjectsTable,
			Columns: operator.ProjectsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   operator.ProjectsTable,
			Columns: operator.ProjectsPrimaryKey,
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
	if ouo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operator.TasksTable,
			Columns: []string{operator.TasksColumn},
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
	if nodes := ouo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !ouo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operator.TasksTable,
			Columns: []string{operator.TasksColumn},
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
	if nodes := ouo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   operator.TasksTable,
			Columns: []string{operator.TasksColumn},
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
	_node = &Operator{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{operator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
