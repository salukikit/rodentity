// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/subnet"
)

// Subnet is the model entity for the Subnet schema.
type Subnet struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Cidr holds the value of the "cidr" field.
	Cidr string `json:"cidr,omitempty"`
	// Mask holds the value of the "mask" field.
	Mask []byte `json:"mask,omitempty"`
	// OutboundTcpports holds the value of the "outbound_tcpports" field.
	OutboundTcpports []string `json:"outbound_tcpports,omitempty"`
	// OutboundUdpports holds the value of the "outbound_udpports" field.
	OutboundUdpports []string `json:"outbound_udpports,omitempty"`
	// Proxy holds the value of the "proxy" field.
	Proxy bool `json:"proxy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubnetQuery when eager-loading is set.
	Edges           SubnetEdges `json:"edges"`
	services_subnet *int
}

// SubnetEdges holds the relations/edges for other nodes in the graph.
type SubnetEdges struct {
	// Hosts holds the value of the hosts edge.
	Hosts []*Device `json:"hosts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HostsOrErr returns the Hosts value or an error if the edge
// was not loaded in eager-loading.
func (e SubnetEdges) HostsOrErr() ([]*Device, error) {
	if e.loadedTypes[0] {
		return e.Hosts, nil
	}
	return nil, &NotLoadedError{edge: "hosts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subnet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subnet.FieldMask, subnet.FieldOutboundTcpports, subnet.FieldOutboundUdpports:
			values[i] = new([]byte)
		case subnet.FieldProxy:
			values[i] = new(sql.NullBool)
		case subnet.FieldCidr:
			values[i] = new(sql.NullString)
		case subnet.FieldID:
			values[i] = new(xid.ID)
		case subnet.ForeignKeys[0]: // services_subnet
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subnet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subnet fields.
func (s *Subnet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subnet.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subnet.FieldCidr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cidr", values[i])
			} else if value.Valid {
				s.Cidr = value.String
			}
		case subnet.FieldMask:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field mask", values[i])
			} else if value != nil {
				s.Mask = *value
			}
		case subnet.FieldOutboundTcpports:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field outbound_tcpports", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.OutboundTcpports); err != nil {
					return fmt.Errorf("unmarshal field outbound_tcpports: %w", err)
				}
			}
		case subnet.FieldOutboundUdpports:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field outbound_udpports", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.OutboundUdpports); err != nil {
					return fmt.Errorf("unmarshal field outbound_udpports: %w", err)
				}
			}
		case subnet.FieldProxy:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field proxy", values[i])
			} else if value.Valid {
				s.Proxy = value.Bool
			}
		case subnet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field services_subnet", value)
			} else if value.Valid {
				s.services_subnet = new(int)
				*s.services_subnet = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryHosts queries the "hosts" edge of the Subnet entity.
func (s *Subnet) QueryHosts() *DeviceQuery {
	return NewSubnetClient(s.config).QueryHosts(s)
}

// Update returns a builder for updating this Subnet.
// Note that you need to call Subnet.Unwrap() before calling this method if this Subnet
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subnet) Update() *SubnetUpdateOne {
	return NewSubnetClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subnet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subnet) Unwrap() *Subnet {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subnet is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subnet) String() string {
	var builder strings.Builder
	builder.WriteString("Subnet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("cidr=")
	builder.WriteString(s.Cidr)
	builder.WriteString(", ")
	builder.WriteString("mask=")
	builder.WriteString(fmt.Sprintf("%v", s.Mask))
	builder.WriteString(", ")
	builder.WriteString("outbound_tcpports=")
	builder.WriteString(fmt.Sprintf("%v", s.OutboundTcpports))
	builder.WriteString(", ")
	builder.WriteString("outbound_udpports=")
	builder.WriteString(fmt.Sprintf("%v", s.OutboundUdpports))
	builder.WriteString(", ")
	builder.WriteString("proxy=")
	builder.WriteString(fmt.Sprintf("%v", s.Proxy))
	builder.WriteByte(')')
	return builder.String()
}

// Subnets is a parsable slice of Subnet.
type Subnets []*Subnet
