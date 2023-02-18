// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/task"
)

// Loot is the model entity for the Loot schema.
type Loot struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type loot.Type `json:"type,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
	// Collectedon holds the value of the "collectedon" field.
	Collectedon time.Time `json:"collectedon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LootQuery when eager-loading is set.
	Edges       LootEdges `json:"edges"`
	rodent_loot *int
	task_loot   *int
}

// LootEdges holds the relations/edges for other nodes in the graph.
type LootEdges struct {
	// Rodent holds the value of the rodent edge.
	Rodent *Rodent `json:"rodent,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RodentOrErr returns the Rodent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LootEdges) RodentOrErr() (*Rodent, error) {
	if e.loadedTypes[0] {
		if e.Rodent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: rodent.Label}
		}
		return e.Rodent, nil
	}
	return nil, &NotLoadedError{edge: "rodent"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LootEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[1] {
		if e.Task == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Loot) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case loot.FieldData:
			values[i] = new([]byte)
		case loot.FieldID:
			values[i] = new(sql.NullInt64)
		case loot.FieldType, loot.FieldLocation:
			values[i] = new(sql.NullString)
		case loot.FieldCollectedon:
			values[i] = new(sql.NullTime)
		case loot.ForeignKeys[0]: // rodent_loot
			values[i] = new(sql.NullInt64)
		case loot.ForeignKeys[1]: // task_loot
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Loot", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Loot fields.
func (l *Loot) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loot.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case loot.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				l.Type = loot.Type(value.String)
			}
		case loot.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				l.Location = value.String
			}
		case loot.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				l.Data = *value
			}
		case loot.FieldCollectedon:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field collectedon", values[i])
			} else if value.Valid {
				l.Collectedon = value.Time
			}
		case loot.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field rodent_loot", value)
			} else if value.Valid {
				l.rodent_loot = new(int)
				*l.rodent_loot = int(value.Int64)
			}
		case loot.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_loot", value)
			} else if value.Valid {
				l.task_loot = new(int)
				*l.task_loot = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRodent queries the "rodent" edge of the Loot entity.
func (l *Loot) QueryRodent() *RodentQuery {
	return NewLootClient(l.config).QueryRodent(l)
}

// QueryTask queries the "task" edge of the Loot entity.
func (l *Loot) QueryTask() *TaskQuery {
	return NewLootClient(l.config).QueryTask(l)
}

// Update returns a builder for updating this Loot.
// Note that you need to call Loot.Unwrap() before calling this method if this Loot
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Loot) Update() *LootUpdateOne {
	return NewLootClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Loot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Loot) Unwrap() *Loot {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Loot is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Loot) String() string {
	var builder strings.Builder
	builder.WriteString("Loot(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", l.Type))
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(l.Location)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", l.Data))
	builder.WriteString(", ")
	builder.WriteString("collectedon=")
	builder.WriteString(l.Collectedon.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Loots is a parsable slice of Loot.
type Loots []*Loot
