// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/salukikit/rodentity/ent/operator"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Xid holds the value of the "xid" field.
	Xid string `json:"xid,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Args holds the value of the "args" field.
	Args []string `json:"args,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
	// Result holds the value of the "result" field.
	Result []byte `json:"result,omitempty"`
	// Executed holds the value of the "Executed" field.
	Executed bool `json:"Executed,omitempty"`
	// Looted holds the value of the "looted" field.
	Looted bool `json:"looted,omitempty"`
	// Requestedat holds the value of the "requestedat" field.
	Requestedat time.Time `json:"requestedat,omitempty"`
	// Completedat holds the value of the "completedat" field.
	Completedat time.Time `json:"completedat,omitempty"`
	// TTPs holds the value of the "TTPs" field.
	TTPs []string `json:"TTPs,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges          TaskEdges `json:"edges"`
	operator_tasks *int
	rodent_tasks   *int
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Rodent holds the value of the rodent edge.
	Rodent *Rodent `json:"rodent,omitempty"`
	// Operator holds the value of the operator edge.
	Operator *Operator `json:"operator,omitempty"`
	// Loot holds the value of the loot edge.
	Loot []*Loot `json:"loot,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RodentOrErr returns the Rodent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) RodentOrErr() (*Rodent, error) {
	if e.loadedTypes[0] {
		if e.Rodent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: rodent.Label}
		}
		return e.Rodent, nil
	}
	return nil, &NotLoadedError{edge: "rodent"}
}

// OperatorOrErr returns the Operator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) OperatorOrErr() (*Operator, error) {
	if e.loadedTypes[1] {
		if e.Operator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: operator.Label}
		}
		return e.Operator, nil
	}
	return nil, &NotLoadedError{edge: "operator"}
}

// LootOrErr returns the Loot value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) LootOrErr() ([]*Loot, error) {
	if e.loadedTypes[2] {
		return e.Loot, nil
	}
	return nil, &NotLoadedError{edge: "loot"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldArgs, task.FieldData, task.FieldResult, task.FieldTTPs:
			values[i] = new([]byte)
		case task.FieldExecuted, task.FieldLooted:
			values[i] = new(sql.NullBool)
		case task.FieldID:
			values[i] = new(sql.NullInt64)
		case task.FieldXid, task.FieldType:
			values[i] = new(sql.NullString)
		case task.FieldRequestedat, task.FieldCompletedat:
			values[i] = new(sql.NullTime)
		case task.ForeignKeys[0]: // operator_tasks
			values[i] = new(sql.NullInt64)
		case task.ForeignKeys[1]: // rodent_tasks
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case task.FieldXid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field xid", values[i])
			} else if value.Valid {
				t.Xid = value.String
			}
		case task.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = value.String
			}
		case task.FieldArgs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field args", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Args); err != nil {
					return fmt.Errorf("unmarshal field args: %w", err)
				}
			}
		case task.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				t.Data = *value
			}
		case task.FieldResult:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value != nil {
				t.Result = *value
			}
		case task.FieldExecuted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Executed", values[i])
			} else if value.Valid {
				t.Executed = value.Bool
			}
		case task.FieldLooted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field looted", values[i])
			} else if value.Valid {
				t.Looted = value.Bool
			}
		case task.FieldRequestedat:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field requestedat", values[i])
			} else if value.Valid {
				t.Requestedat = value.Time
			}
		case task.FieldCompletedat:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field completedat", values[i])
			} else if value.Valid {
				t.Completedat = value.Time
			}
		case task.FieldTTPs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field TTPs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.TTPs); err != nil {
					return fmt.Errorf("unmarshal field TTPs: %w", err)
				}
			}
		case task.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field operator_tasks", value)
			} else if value.Valid {
				t.operator_tasks = new(int)
				*t.operator_tasks = int(value.Int64)
			}
		case task.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field rodent_tasks", value)
			} else if value.Valid {
				t.rodent_tasks = new(int)
				*t.rodent_tasks = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRodent queries the "rodent" edge of the Task entity.
func (t *Task) QueryRodent() *RodentQuery {
	return NewTaskClient(t.config).QueryRodent(t)
}

// QueryOperator queries the "operator" edge of the Task entity.
func (t *Task) QueryOperator() *OperatorQuery {
	return NewTaskClient(t.config).QueryOperator(t)
}

// QueryLoot queries the "loot" edge of the Task entity.
func (t *Task) QueryLoot() *LootQuery {
	return NewTaskClient(t.config).QueryLoot(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return NewTaskClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("xid=")
	builder.WriteString(t.Xid)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(t.Type)
	builder.WriteString(", ")
	builder.WriteString("args=")
	builder.WriteString(fmt.Sprintf("%v", t.Args))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", t.Data))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", t.Result))
	builder.WriteString(", ")
	builder.WriteString("Executed=")
	builder.WriteString(fmt.Sprintf("%v", t.Executed))
	builder.WriteString(", ")
	builder.WriteString("looted=")
	builder.WriteString(fmt.Sprintf("%v", t.Looted))
	builder.WriteString(", ")
	builder.WriteString("requestedat=")
	builder.WriteString(t.Requestedat.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("completedat=")
	builder.WriteString(t.Completedat.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("TTPs=")
	builder.WriteString(fmt.Sprintf("%v", t.TTPs))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task
