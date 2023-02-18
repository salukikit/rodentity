// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/group"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions string `json:"permissions,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges         GroupEdges `json:"edges"`
	domain_groups *int
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Devices holds the value of the devices edge.
	Devices []*Device `json:"devices,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Domain holds the value of the domain edge.
	Domain *Domain `json:"domain,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) DevicesOrErr() ([]*Device, error) {
	if e.loadedTypes[0] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) DomainOrErr() (*Domain, error) {
	if e.loadedTypes[2] {
		if e.Domain == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: domain.Label}
		}
		return e.Domain, nil
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			values[i] = new(sql.NullInt64)
		case group.FieldName, group.FieldDescription, group.FieldPermissions:
			values[i] = new(sql.NullString)
		case group.ForeignKeys[0]: // domain_groups
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Group", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gr.Description = value.String
			}
		case group.FieldPermissions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value.Valid {
				gr.Permissions = value.String
			}
		case group.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field domain_groups", value)
			} else if value.Valid {
				gr.domain_groups = new(int)
				*gr.domain_groups = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDevices queries the "devices" edge of the Group entity.
func (gr *Group) QueryDevices() *DeviceQuery {
	return NewGroupClient(gr.config).QueryDevices(gr)
}

// QueryUsers queries the "users" edge of the Group entity.
func (gr *Group) QueryUsers() *UserQuery {
	return NewGroupClient(gr.config).QueryUsers(gr)
}

// QueryDomain queries the "domain" edge of the Group entity.
func (gr *Group) QueryDomain() *DomainQuery {
	return NewGroupClient(gr.config).QueryDomain(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", ")
	builder.WriteString("permissions=")
	builder.WriteString(gr.Permissions)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group
