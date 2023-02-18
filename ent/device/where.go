// Code generated by ent, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/salukikit/rodentity/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldID, id))
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldHostname, v))
}

// Os applies equality check predicate on the "os" field. It's identical to OsEQ.
func Os(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldOs, v))
}

// Arch applies equality check predicate on the "arch" field. It's identical to ArchEQ.
func Arch(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldArch, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldVersion, v))
}

// Machinepass applies equality check predicate on the "machinepass" field. It's identical to MachinepassEQ.
func Machinepass(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldMachinepass, v))
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldHostname, v))
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldHostname, v))
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldHostname, vs...))
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldHostname, vs...))
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldHostname, v))
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldHostname, v))
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldHostname, v))
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldHostname, v))
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldHostname, v))
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldHostname, v))
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldHostname, v))
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldHostname, v))
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldHostname, v))
}

// OsEQ applies the EQ predicate on the "os" field.
func OsEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldOs, v))
}

// OsNEQ applies the NEQ predicate on the "os" field.
func OsNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldOs, v))
}

// OsIn applies the In predicate on the "os" field.
func OsIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldOs, vs...))
}

// OsNotIn applies the NotIn predicate on the "os" field.
func OsNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldOs, vs...))
}

// OsGT applies the GT predicate on the "os" field.
func OsGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldOs, v))
}

// OsGTE applies the GTE predicate on the "os" field.
func OsGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldOs, v))
}

// OsLT applies the LT predicate on the "os" field.
func OsLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldOs, v))
}

// OsLTE applies the LTE predicate on the "os" field.
func OsLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldOs, v))
}

// OsContains applies the Contains predicate on the "os" field.
func OsContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldOs, v))
}

// OsHasPrefix applies the HasPrefix predicate on the "os" field.
func OsHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldOs, v))
}

// OsHasSuffix applies the HasSuffix predicate on the "os" field.
func OsHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldOs, v))
}

// OsEqualFold applies the EqualFold predicate on the "os" field.
func OsEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldOs, v))
}

// OsContainsFold applies the ContainsFold predicate on the "os" field.
func OsContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldOs, v))
}

// ArchEQ applies the EQ predicate on the "arch" field.
func ArchEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldArch, v))
}

// ArchNEQ applies the NEQ predicate on the "arch" field.
func ArchNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldArch, v))
}

// ArchIn applies the In predicate on the "arch" field.
func ArchIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldArch, vs...))
}

// ArchNotIn applies the NotIn predicate on the "arch" field.
func ArchNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldArch, vs...))
}

// ArchGT applies the GT predicate on the "arch" field.
func ArchGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldArch, v))
}

// ArchGTE applies the GTE predicate on the "arch" field.
func ArchGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldArch, v))
}

// ArchLT applies the LT predicate on the "arch" field.
func ArchLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldArch, v))
}

// ArchLTE applies the LTE predicate on the "arch" field.
func ArchLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldArch, v))
}

// ArchContains applies the Contains predicate on the "arch" field.
func ArchContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldArch, v))
}

// ArchHasPrefix applies the HasPrefix predicate on the "arch" field.
func ArchHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldArch, v))
}

// ArchHasSuffix applies the HasSuffix predicate on the "arch" field.
func ArchHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldArch, v))
}

// ArchEqualFold applies the EqualFold predicate on the "arch" field.
func ArchEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldArch, v))
}

// ArchContainsFold applies the ContainsFold predicate on the "arch" field.
func ArchContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldArch, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldVersion, v))
}

// NetInterfacesIsNil applies the IsNil predicate on the "net_interfaces" field.
func NetInterfacesIsNil() predicate.Device {
	return predicate.Device(sql.FieldIsNull(FieldNetInterfaces))
}

// NetInterfacesNotNil applies the NotNil predicate on the "net_interfaces" field.
func NetInterfacesNotNil() predicate.Device {
	return predicate.Device(sql.FieldNotNull(FieldNetInterfaces))
}

// MachinepassEQ applies the EQ predicate on the "machinepass" field.
func MachinepassEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldMachinepass, v))
}

// MachinepassNEQ applies the NEQ predicate on the "machinepass" field.
func MachinepassNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldMachinepass, v))
}

// MachinepassIn applies the In predicate on the "machinepass" field.
func MachinepassIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldMachinepass, vs...))
}

// MachinepassNotIn applies the NotIn predicate on the "machinepass" field.
func MachinepassNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldMachinepass, vs...))
}

// MachinepassGT applies the GT predicate on the "machinepass" field.
func MachinepassGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldMachinepass, v))
}

// MachinepassGTE applies the GTE predicate on the "machinepass" field.
func MachinepassGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldMachinepass, v))
}

// MachinepassLT applies the LT predicate on the "machinepass" field.
func MachinepassLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldMachinepass, v))
}

// MachinepassLTE applies the LTE predicate on the "machinepass" field.
func MachinepassLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldMachinepass, v))
}

// MachinepassContains applies the Contains predicate on the "machinepass" field.
func MachinepassContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldMachinepass, v))
}

// MachinepassHasPrefix applies the HasPrefix predicate on the "machinepass" field.
func MachinepassHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldMachinepass, v))
}

// MachinepassHasSuffix applies the HasSuffix predicate on the "machinepass" field.
func MachinepassHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldMachinepass, v))
}

// MachinepassIsNil applies the IsNil predicate on the "machinepass" field.
func MachinepassIsNil() predicate.Device {
	return predicate.Device(sql.FieldIsNull(FieldMachinepass))
}

// MachinepassNotNil applies the NotNil predicate on the "machinepass" field.
func MachinepassNotNil() predicate.Device {
	return predicate.Device(sql.FieldNotNull(FieldMachinepass))
}

// MachinepassEqualFold applies the EqualFold predicate on the "machinepass" field.
func MachinepassEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldMachinepass, v))
}

// MachinepassContainsFold applies the ContainsFold predicate on the "machinepass" field.
func MachinepassContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldMachinepass, v))
}

// CertificatesIsNil applies the IsNil predicate on the "certificates" field.
func CertificatesIsNil() predicate.Device {
	return predicate.Device(sql.FieldIsNull(FieldCertificates))
}

// CertificatesNotNil applies the NotNil predicate on the "certificates" field.
func CertificatesNotNil() predicate.Device {
	return predicate.Device(sql.FieldNotNull(FieldCertificates))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRodents applies the HasEdge predicate on the "rodents" edge.
func HasRodents() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RodentsTable, RodentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRodentsWith applies the HasEdge predicate on the "rodents" edge with a given conditions (other predicates).
func HasRodentsWith(preds ...predicate.Rodent) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RodentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RodentsTable, RodentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDomain applies the HasEdge predicate on the "domain" edge.
func HasDomain() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DomainTable, DomainColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDomainWith applies the HasEdge predicate on the "domain" edge with a given conditions (other predicates).
func HasDomainWith(preds ...predicate.Domain) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DomainInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DomainTable, DomainColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubnets applies the HasEdge predicate on the "subnets" edge.
func HasSubnets() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubnetsTable, SubnetsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubnetsWith applies the HasEdge predicate on the "subnets" edge with a given conditions (other predicates).
func HasSubnetsWith(preds ...predicate.Subnet) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubnetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubnetsTable, SubnetsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasServices applies the HasEdge predicate on the "services" edge.
func HasServices() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ServicesTable, ServicesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServicesWith applies the HasEdge predicate on the "services" edge with a given conditions (other predicates).
func HasServicesWith(preds ...predicate.Services) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ServicesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ServicesTable, ServicesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		p(s.Not())
	})
}
