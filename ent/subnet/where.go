// Code generated by ent, DO NOT EDIT.

package subnet

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
	"github.com/salukikit/rodentity/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Subnet {
	return predicate.Subnet(sql.FieldLTE(FieldID, id))
}

// Cidr applies equality check predicate on the "cidr" field. It's identical to CidrEQ.
func Cidr(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldCidr, v))
}

// Mask applies equality check predicate on the "mask" field. It's identical to MaskEQ.
func Mask(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldMask, v))
}

// Proxy applies equality check predicate on the "proxy" field. It's identical to ProxyEQ.
func Proxy(v bool) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldProxy, v))
}

// CidrEQ applies the EQ predicate on the "cidr" field.
func CidrEQ(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldCidr, v))
}

// CidrNEQ applies the NEQ predicate on the "cidr" field.
func CidrNEQ(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldNEQ(FieldCidr, v))
}

// CidrIn applies the In predicate on the "cidr" field.
func CidrIn(vs ...string) predicate.Subnet {
	return predicate.Subnet(sql.FieldIn(FieldCidr, vs...))
}

// CidrNotIn applies the NotIn predicate on the "cidr" field.
func CidrNotIn(vs ...string) predicate.Subnet {
	return predicate.Subnet(sql.FieldNotIn(FieldCidr, vs...))
}

// CidrGT applies the GT predicate on the "cidr" field.
func CidrGT(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldGT(FieldCidr, v))
}

// CidrGTE applies the GTE predicate on the "cidr" field.
func CidrGTE(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldGTE(FieldCidr, v))
}

// CidrLT applies the LT predicate on the "cidr" field.
func CidrLT(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldLT(FieldCidr, v))
}

// CidrLTE applies the LTE predicate on the "cidr" field.
func CidrLTE(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldLTE(FieldCidr, v))
}

// CidrContains applies the Contains predicate on the "cidr" field.
func CidrContains(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldContains(FieldCidr, v))
}

// CidrHasPrefix applies the HasPrefix predicate on the "cidr" field.
func CidrHasPrefix(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldHasPrefix(FieldCidr, v))
}

// CidrHasSuffix applies the HasSuffix predicate on the "cidr" field.
func CidrHasSuffix(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldHasSuffix(FieldCidr, v))
}

// CidrEqualFold applies the EqualFold predicate on the "cidr" field.
func CidrEqualFold(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldEqualFold(FieldCidr, v))
}

// CidrContainsFold applies the ContainsFold predicate on the "cidr" field.
func CidrContainsFold(v string) predicate.Subnet {
	return predicate.Subnet(sql.FieldContainsFold(FieldCidr, v))
}

// MaskEQ applies the EQ predicate on the "mask" field.
func MaskEQ(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldMask, v))
}

// MaskNEQ applies the NEQ predicate on the "mask" field.
func MaskNEQ(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldNEQ(FieldMask, v))
}

// MaskIn applies the In predicate on the "mask" field.
func MaskIn(vs ...[]byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldIn(FieldMask, vs...))
}

// MaskNotIn applies the NotIn predicate on the "mask" field.
func MaskNotIn(vs ...[]byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldNotIn(FieldMask, vs...))
}

// MaskGT applies the GT predicate on the "mask" field.
func MaskGT(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldGT(FieldMask, v))
}

// MaskGTE applies the GTE predicate on the "mask" field.
func MaskGTE(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldGTE(FieldMask, v))
}

// MaskLT applies the LT predicate on the "mask" field.
func MaskLT(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldLT(FieldMask, v))
}

// MaskLTE applies the LTE predicate on the "mask" field.
func MaskLTE(v []byte) predicate.Subnet {
	return predicate.Subnet(sql.FieldLTE(FieldMask, v))
}

// MaskIsNil applies the IsNil predicate on the "mask" field.
func MaskIsNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldIsNull(FieldMask))
}

// MaskNotNil applies the NotNil predicate on the "mask" field.
func MaskNotNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldNotNull(FieldMask))
}

// OutboundTcpportsIsNil applies the IsNil predicate on the "outbound_tcpports" field.
func OutboundTcpportsIsNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldIsNull(FieldOutboundTcpports))
}

// OutboundTcpportsNotNil applies the NotNil predicate on the "outbound_tcpports" field.
func OutboundTcpportsNotNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldNotNull(FieldOutboundTcpports))
}

// OutboundUdpportsIsNil applies the IsNil predicate on the "outbound_udpports" field.
func OutboundUdpportsIsNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldIsNull(FieldOutboundUdpports))
}

// OutboundUdpportsNotNil applies the NotNil predicate on the "outbound_udpports" field.
func OutboundUdpportsNotNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldNotNull(FieldOutboundUdpports))
}

// ProxyEQ applies the EQ predicate on the "proxy" field.
func ProxyEQ(v bool) predicate.Subnet {
	return predicate.Subnet(sql.FieldEQ(FieldProxy, v))
}

// ProxyNEQ applies the NEQ predicate on the "proxy" field.
func ProxyNEQ(v bool) predicate.Subnet {
	return predicate.Subnet(sql.FieldNEQ(FieldProxy, v))
}

// ProxyIsNil applies the IsNil predicate on the "proxy" field.
func ProxyIsNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldIsNull(FieldProxy))
}

// ProxyNotNil applies the NotNil predicate on the "proxy" field.
func ProxyNotNil() predicate.Subnet {
	return predicate.Subnet(sql.FieldNotNull(FieldProxy))
}

// HasHosts applies the HasEdge predicate on the "hosts" edge.
func HasHosts() predicate.Subnet {
	return predicate.Subnet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, HostsTable, HostsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostsWith applies the HasEdge predicate on the "hosts" edge with a given conditions (other predicates).
func HasHostsWith(preds ...predicate.Device) predicate.Subnet {
	return predicate.Subnet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, HostsTable, HostsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subnet) predicate.Subnet {
	return predicate.Subnet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subnet) predicate.Subnet {
	return predicate.Subnet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subnet) predicate.Subnet {
	return predicate.Subnet(func(s *sql.Selector) {
		p(s.Not())
	})
}
