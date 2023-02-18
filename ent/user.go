// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Givenname holds the value of the "givenname" field.
	Givenname string `json:"givenname,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Owned holds the value of the "owned" field.
	Owned bool `json:"owned,omitempty"`
	// Age holds the value of the "age" field.
	Age string `json:"age,omitempty"`
	// Homedir holds the value of the "homedir" field.
	Homedir string `json:"homedir,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	domain_users *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Devices holds the value of the devices edge.
	Devices []*Device `json:"devices,omitempty"`
	// Rodents holds the value of the rodents edge.
	Rodents []*Rodent `json:"rodents,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Domain holds the value of the domain edge.
	Domain *Domain `json:"domain,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DevicesOrErr() ([]*Device, error) {
	if e.loadedTypes[0] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// RodentsOrErr returns the Rodents value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RodentsOrErr() ([]*Rodent, error) {
	if e.loadedTypes[1] {
		return e.Rodents, nil
	}
	return nil, &NotLoadedError{edge: "rodents"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) DomainOrErr() (*Domain, error) {
	if e.loadedTypes[3] {
		if e.Domain == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: domain.Label}
		}
		return e.Domain, nil
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldOwned, user.FieldEnabled:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldGivenname, user.FieldEmail, user.FieldAge, user.FieldHomedir:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // domain_users
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldGivenname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field givenname", values[i])
			} else if value.Valid {
				u.Givenname = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldOwned:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field owned", values[i])
			} else if value.Valid {
				u.Owned = value.Bool
			}
		case user.FieldAge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				u.Age = value.String
			}
		case user.FieldHomedir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field homedir", values[i])
			} else if value.Valid {
				u.Homedir = value.String
			}
		case user.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				u.Enabled = value.Bool
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field domain_users", value)
			} else if value.Valid {
				u.domain_users = new(int)
				*u.domain_users = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDevices queries the "devices" edge of the User entity.
func (u *User) QueryDevices() *DeviceQuery {
	return NewUserClient(u.config).QueryDevices(u)
}

// QueryRodents queries the "rodents" edge of the User entity.
func (u *User) QueryRodents() *RodentQuery {
	return NewUserClient(u.config).QueryRodents(u)
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return NewUserClient(u.config).QueryGroups(u)
}

// QueryDomain queries the "domain" edge of the User entity.
func (u *User) QueryDomain() *DomainQuery {
	return NewUserClient(u.config).QueryDomain(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("givenname=")
	builder.WriteString(u.Givenname)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("owned=")
	builder.WriteString(fmt.Sprintf("%v", u.Owned))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(u.Age)
	builder.WriteString(", ")
	builder.WriteString("homedir=")
	builder.WriteString(u.Homedir)
	builder.WriteString(", ")
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", u.Enabled))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User