// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/salukikit/rodentity/ent/operator"
)

// Operator is the model entity for the Operator schema.
type Operator struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Privkey holds the value of the "privkey" field.
	Privkey []byte `json:"privkey,omitempty"`
	// Cert holds the value of the "cert" field.
	Cert []byte `json:"cert,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OperatorQuery when eager-loading is set.
	Edges OperatorEdges `json:"edges"`
}

// OperatorEdges holds the relations/edges for other nodes in the graph.
type OperatorEdges struct {
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e OperatorEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Operator) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case operator.FieldPrivkey, operator.FieldCert:
			values[i] = new([]byte)
		case operator.FieldID:
			values[i] = new(sql.NullInt64)
		case operator.FieldUsername:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Operator", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Operator fields.
func (o *Operator) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case operator.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case operator.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				o.Username = value.String
			}
		case operator.FieldPrivkey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field privkey", values[i])
			} else if value != nil {
				o.Privkey = *value
			}
		case operator.FieldCert:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cert", values[i])
			} else if value != nil {
				o.Cert = *value
			}
		}
	}
	return nil
}

// QueryProjects queries the "projects" edge of the Operator entity.
func (o *Operator) QueryProjects() *ProjectQuery {
	return NewOperatorClient(o.config).QueryProjects(o)
}

// Update returns a builder for updating this Operator.
// Note that you need to call Operator.Unwrap() before calling this method if this Operator
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Operator) Update() *OperatorUpdateOne {
	return NewOperatorClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Operator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Operator) Unwrap() *Operator {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Operator is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Operator) String() string {
	var builder strings.Builder
	builder.WriteString("Operator(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("username=")
	builder.WriteString(o.Username)
	builder.WriteString(", ")
	builder.WriteString("privkey=")
	builder.WriteString(fmt.Sprintf("%v", o.Privkey))
	builder.WriteString(", ")
	builder.WriteString("cert=")
	builder.WriteString(fmt.Sprintf("%v", o.Cert))
	builder.WriteByte(')')
	return builder.String()
}

// Operators is a parsable slice of Operator.
type Operators []*Operator