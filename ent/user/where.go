// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/salukikit/rodentity/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Givenname applies equality check predicate on the "givenname" field. It's identical to GivennameEQ.
func Givenname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGivenname, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Owned applies equality check predicate on the "owned" field. It's identical to OwnedEQ.
func Owned(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOwned, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// Homedir applies equality check predicate on the "homedir" field. It's identical to HomedirEQ.
func Homedir(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomedir, v))
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEnabled, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// GivennameEQ applies the EQ predicate on the "givenname" field.
func GivennameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGivenname, v))
}

// GivennameNEQ applies the NEQ predicate on the "givenname" field.
func GivennameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGivenname, v))
}

// GivennameIn applies the In predicate on the "givenname" field.
func GivennameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGivenname, vs...))
}

// GivennameNotIn applies the NotIn predicate on the "givenname" field.
func GivennameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGivenname, vs...))
}

// GivennameGT applies the GT predicate on the "givenname" field.
func GivennameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGivenname, v))
}

// GivennameGTE applies the GTE predicate on the "givenname" field.
func GivennameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGivenname, v))
}

// GivennameLT applies the LT predicate on the "givenname" field.
func GivennameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGivenname, v))
}

// GivennameLTE applies the LTE predicate on the "givenname" field.
func GivennameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGivenname, v))
}

// GivennameContains applies the Contains predicate on the "givenname" field.
func GivennameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGivenname, v))
}

// GivennameHasPrefix applies the HasPrefix predicate on the "givenname" field.
func GivennameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGivenname, v))
}

// GivennameHasSuffix applies the HasSuffix predicate on the "givenname" field.
func GivennameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGivenname, v))
}

// GivennameEqualFold applies the EqualFold predicate on the "givenname" field.
func GivennameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGivenname, v))
}

// GivennameContainsFold applies the ContainsFold predicate on the "givenname" field.
func GivennameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGivenname, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// OwnedEQ applies the EQ predicate on the "owned" field.
func OwnedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOwned, v))
}

// OwnedNEQ applies the NEQ predicate on the "owned" field.
func OwnedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOwned, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAge, v))
}

// AgeContains applies the Contains predicate on the "age" field.
func AgeContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAge, v))
}

// AgeHasPrefix applies the HasPrefix predicate on the "age" field.
func AgeHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAge, v))
}

// AgeHasSuffix applies the HasSuffix predicate on the "age" field.
func AgeHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAge, v))
}

// AgeIsNil applies the IsNil predicate on the "age" field.
func AgeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAge))
}

// AgeNotNil applies the NotNil predicate on the "age" field.
func AgeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAge))
}

// AgeEqualFold applies the EqualFold predicate on the "age" field.
func AgeEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAge, v))
}

// AgeContainsFold applies the ContainsFold predicate on the "age" field.
func AgeContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAge, v))
}

// HomedirEQ applies the EQ predicate on the "homedir" field.
func HomedirEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomedir, v))
}

// HomedirNEQ applies the NEQ predicate on the "homedir" field.
func HomedirNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHomedir, v))
}

// HomedirIn applies the In predicate on the "homedir" field.
func HomedirIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHomedir, vs...))
}

// HomedirNotIn applies the NotIn predicate on the "homedir" field.
func HomedirNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHomedir, vs...))
}

// HomedirGT applies the GT predicate on the "homedir" field.
func HomedirGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHomedir, v))
}

// HomedirGTE applies the GTE predicate on the "homedir" field.
func HomedirGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHomedir, v))
}

// HomedirLT applies the LT predicate on the "homedir" field.
func HomedirLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHomedir, v))
}

// HomedirLTE applies the LTE predicate on the "homedir" field.
func HomedirLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHomedir, v))
}

// HomedirContains applies the Contains predicate on the "homedir" field.
func HomedirContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHomedir, v))
}

// HomedirHasPrefix applies the HasPrefix predicate on the "homedir" field.
func HomedirHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHomedir, v))
}

// HomedirHasSuffix applies the HasSuffix predicate on the "homedir" field.
func HomedirHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHomedir, v))
}

// HomedirEqualFold applies the EqualFold predicate on the "homedir" field.
func HomedirEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHomedir, v))
}

// HomedirContainsFold applies the ContainsFold predicate on the "homedir" field.
func HomedirContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHomedir, v))
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEnabled, v))
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEnabled, v))
}

// HasDevices applies the HasEdge predicate on the "devices" edge.
func HasDevices() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DevicesTable, DevicesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDevicesWith applies the HasEdge predicate on the "devices" edge with a given conditions (other predicates).
func HasDevicesWith(preds ...predicate.Device) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DevicesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DevicesTable, DevicesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRodents applies the HasEdge predicate on the "rodents" edge.
func HasRodents() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RodentsTable, RodentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRodentsWith applies the HasEdge predicate on the "rodents" edge with a given conditions (other predicates).
func HasRodentsWith(preds ...predicate.Rodent) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func HasGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func HasDomain() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DomainTable, DomainColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDomainWith applies the HasEdge predicate on the "domain" edge with a given conditions (other predicates).
func HasDomainWith(preds ...predicate.Domain) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
