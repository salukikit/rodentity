// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/salukikit/rodentity/ent/router"
)

// Router is the model entity for the Router schema.
type Router struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Privkey holds the value of the "privkey" field.
	Privkey []byte `json:"privkey,omitempty"`
	// Cert holds the value of the "cert" field.
	Cert []byte `json:"cert,omitempty"`
	// Commands holds the value of the "commands" field.
	Commands []string `json:"commands,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RouterQuery when eager-loading is set.
	Edges RouterEdges `json:"edges"`
}

// RouterEdges holds the relations/edges for other nodes in the graph.
type RouterEdges struct {
	// Rodents holds the value of the rodents edge.
	Rodents []*Rodent `json:"rodents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RodentsOrErr returns the Rodents value or an error if the edge
// was not loaded in eager-loading.
func (e RouterEdges) RodentsOrErr() ([]*Rodent, error) {
	if e.loadedTypes[0] {
		return e.Rodents, nil
	}
	return nil, &NotLoadedError{edge: "rodents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Router) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case router.FieldPrivkey, router.FieldCert, router.FieldCommands:
			values[i] = new([]byte)
		case router.FieldID:
			values[i] = new(sql.NullInt64)
		case router.FieldUsername:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Router", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Router fields.
func (r *Router) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case router.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case router.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				r.Username = value.String
			}
		case router.FieldPrivkey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field privkey", values[i])
			} else if value != nil {
				r.Privkey = *value
			}
		case router.FieldCert:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cert", values[i])
			} else if value != nil {
				r.Cert = *value
			}
		case router.FieldCommands:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field commands", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Commands); err != nil {
					return fmt.Errorf("unmarshal field commands: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryRodents queries the "rodents" edge of the Router entity.
func (r *Router) QueryRodents() *RodentQuery {
	return NewRouterClient(r.config).QueryRodents(r)
}

// Update returns a builder for updating this Router.
// Note that you need to call Router.Unwrap() before calling this method if this Router
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Router) Update() *RouterUpdateOne {
	return NewRouterClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Router entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Router) Unwrap() *Router {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Router is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Router) String() string {
	var builder strings.Builder
	builder.WriteString("Router(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("username=")
	builder.WriteString(r.Username)
	builder.WriteString(", ")
	builder.WriteString("privkey=")
	builder.WriteString(fmt.Sprintf("%v", r.Privkey))
	builder.WriteString(", ")
	builder.WriteString("cert=")
	builder.WriteString(fmt.Sprintf("%v", r.Cert))
	builder.WriteString(", ")
	builder.WriteString("commands=")
	builder.WriteString(fmt.Sprintf("%v", r.Commands))
	builder.WriteByte(')')
	return builder.String()
}

// Routers is a parsable slice of Router.
type Routers []*Router
