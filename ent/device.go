// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/domain"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// Os holds the value of the "os" field.
	Os string `json:"os,omitempty"`
	// Arch holds the value of the "arch" field.
	Arch string `json:"arch,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// NetInterfaces holds the value of the "net_interfaces" field.
	NetInterfaces []string `json:"net_interfaces,omitempty"`
	// Machinepass holds the value of the "machinepass" field.
	Machinepass string `json:"machinepass,omitempty"`
	// Certificates holds the value of the "certificates" field.
	Certificates []string `json:"certificates,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges          DeviceEdges `json:"edges"`
	domain_devices *xid.ID
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Rodents holds the value of the rodents edge.
	Rodents []*Rodent `json:"rodents,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Domain holds the value of the domain edge.
	Domain *Domain `json:"domain,omitempty"`
	// Subnets holds the value of the subnets edge.
	Subnets []*Subnet `json:"subnets,omitempty"`
	// Services holds the value of the services edge.
	Services []*Services `json:"services,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// RodentsOrErr returns the Rodents value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) RodentsOrErr() ([]*Rodent, error) {
	if e.loadedTypes[1] {
		return e.Rodents, nil
	}
	return nil, &NotLoadedError{edge: "rodents"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) DomainOrErr() (*Domain, error) {
	if e.loadedTypes[3] {
		if e.Domain == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: domain.Label}
		}
		return e.Domain, nil
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// SubnetsOrErr returns the Subnets value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) SubnetsOrErr() ([]*Subnet, error) {
	if e.loadedTypes[4] {
		return e.Subnets, nil
	}
	return nil, &NotLoadedError{edge: "subnets"}
}

// ServicesOrErr returns the Services value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) ServicesOrErr() ([]*Services, error) {
	if e.loadedTypes[5] {
		return e.Services, nil
	}
	return nil, &NotLoadedError{edge: "services"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldNetInterfaces, device.FieldCertificates:
			values[i] = new([]byte)
		case device.FieldHostname, device.FieldOs, device.FieldArch, device.FieldVersion, device.FieldMachinepass:
			values[i] = new(sql.NullString)
		case device.FieldID:
			values[i] = new(xid.ID)
		case device.ForeignKeys[0]: // domain_devices
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Device", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case device.FieldHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[i])
			} else if value.Valid {
				d.Hostname = value.String
			}
		case device.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				d.Os = value.String
			}
		case device.FieldArch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arch", values[i])
			} else if value.Valid {
				d.Arch = value.String
			}
		case device.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				d.Version = value.String
			}
		case device.FieldNetInterfaces:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field net_interfaces", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.NetInterfaces); err != nil {
					return fmt.Errorf("unmarshal field net_interfaces: %w", err)
				}
			}
		case device.FieldMachinepass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field machinepass", values[i])
			} else if value.Valid {
				d.Machinepass = value.String
			}
		case device.FieldCertificates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field certificates", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.Certificates); err != nil {
					return fmt.Errorf("unmarshal field certificates: %w", err)
				}
			}
		case device.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field domain_devices", values[i])
			} else if value.Valid {
				d.domain_devices = new(xid.ID)
				*d.domain_devices = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Device entity.
func (d *Device) QueryUsers() *UserQuery {
	return NewDeviceClient(d.config).QueryUsers(d)
}

// QueryRodents queries the "rodents" edge of the Device entity.
func (d *Device) QueryRodents() *RodentQuery {
	return NewDeviceClient(d.config).QueryRodents(d)
}

// QueryGroups queries the "groups" edge of the Device entity.
func (d *Device) QueryGroups() *GroupQuery {
	return NewDeviceClient(d.config).QueryGroups(d)
}

// QueryDomain queries the "domain" edge of the Device entity.
func (d *Device) QueryDomain() *DomainQuery {
	return NewDeviceClient(d.config).QueryDomain(d)
}

// QuerySubnets queries the "subnets" edge of the Device entity.
func (d *Device) QuerySubnets() *SubnetQuery {
	return NewDeviceClient(d.config).QuerySubnets(d)
}

// QueryServices queries the "services" edge of the Device entity.
func (d *Device) QueryServices() *ServicesQuery {
	return NewDeviceClient(d.config).QueryServices(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return NewDeviceClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("hostname=")
	builder.WriteString(d.Hostname)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(d.Os)
	builder.WriteString(", ")
	builder.WriteString("arch=")
	builder.WriteString(d.Arch)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(d.Version)
	builder.WriteString(", ")
	builder.WriteString("net_interfaces=")
	builder.WriteString(fmt.Sprintf("%v", d.NetInterfaces))
	builder.WriteString(", ")
	builder.WriteString("machinepass=")
	builder.WriteString(d.Machinepass)
	builder.WriteString(", ")
	builder.WriteString("certificates=")
	builder.WriteString(fmt.Sprintf("%v", d.Certificates))
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device
