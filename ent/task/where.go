// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/salukikit/rodentity/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// Xid applies equality check predicate on the "xid" field. It's identical to XidEQ.
func Xid(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldXid, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldType, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldData, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v []byte) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldResult, v))
}

// Executed applies equality check predicate on the "Executed" field. It's identical to ExecutedEQ.
func Executed(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldExecuted, v))
}

// Looted applies equality check predicate on the "looted" field. It's identical to LootedEQ.
func Looted(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldLooted, v))
}

// Requestedat applies equality check predicate on the "requestedat" field. It's identical to RequestedatEQ.
func Requestedat(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldRequestedat, v))
}

// Completedat applies equality check predicate on the "completedat" field. It's identical to CompletedatEQ.
func Completedat(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCompletedat, v))
}

// XidEQ applies the EQ predicate on the "xid" field.
func XidEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldXid, v))
}

// XidNEQ applies the NEQ predicate on the "xid" field.
func XidNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldXid, v))
}

// XidIn applies the In predicate on the "xid" field.
func XidIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldXid, vs...))
}

// XidNotIn applies the NotIn predicate on the "xid" field.
func XidNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldXid, vs...))
}

// XidGT applies the GT predicate on the "xid" field.
func XidGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldXid, v))
}

// XidGTE applies the GTE predicate on the "xid" field.
func XidGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldXid, v))
}

// XidLT applies the LT predicate on the "xid" field.
func XidLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldXid, v))
}

// XidLTE applies the LTE predicate on the "xid" field.
func XidLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldXid, v))
}

// XidContains applies the Contains predicate on the "xid" field.
func XidContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldXid, v))
}

// XidHasPrefix applies the HasPrefix predicate on the "xid" field.
func XidHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldXid, v))
}

// XidHasSuffix applies the HasSuffix predicate on the "xid" field.
func XidHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldXid, v))
}

// XidEqualFold applies the EqualFold predicate on the "xid" field.
func XidEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldXid, v))
}

// XidContainsFold applies the ContainsFold predicate on the "xid" field.
func XidContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldXid, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldType, v))
}

// ArgsIsNil applies the IsNil predicate on the "args" field.
func ArgsIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldArgs))
}

// ArgsNotNil applies the NotNil predicate on the "args" field.
func ArgsNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldArgs))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldData, v))
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldData))
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldData))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v []byte) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldResult, v))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v []byte) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldResult, v))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...[]byte) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldResult, vs...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...[]byte) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldResult, vs...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v []byte) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldResult, v))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v []byte) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldResult, v))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v []byte) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldResult, v))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v []byte) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldResult, v))
}

// ResultIsNil applies the IsNil predicate on the "result" field.
func ResultIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldResult))
}

// ResultNotNil applies the NotNil predicate on the "result" field.
func ResultNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldResult))
}

// ExecutedEQ applies the EQ predicate on the "Executed" field.
func ExecutedEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldExecuted, v))
}

// ExecutedNEQ applies the NEQ predicate on the "Executed" field.
func ExecutedNEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldExecuted, v))
}

// LootedEQ applies the EQ predicate on the "looted" field.
func LootedEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldLooted, v))
}

// LootedNEQ applies the NEQ predicate on the "looted" field.
func LootedNEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldLooted, v))
}

// RequestedatEQ applies the EQ predicate on the "requestedat" field.
func RequestedatEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldRequestedat, v))
}

// RequestedatNEQ applies the NEQ predicate on the "requestedat" field.
func RequestedatNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldRequestedat, v))
}

// RequestedatIn applies the In predicate on the "requestedat" field.
func RequestedatIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldRequestedat, vs...))
}

// RequestedatNotIn applies the NotIn predicate on the "requestedat" field.
func RequestedatNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldRequestedat, vs...))
}

// RequestedatGT applies the GT predicate on the "requestedat" field.
func RequestedatGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldRequestedat, v))
}

// RequestedatGTE applies the GTE predicate on the "requestedat" field.
func RequestedatGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldRequestedat, v))
}

// RequestedatLT applies the LT predicate on the "requestedat" field.
func RequestedatLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldRequestedat, v))
}

// RequestedatLTE applies the LTE predicate on the "requestedat" field.
func RequestedatLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldRequestedat, v))
}

// CompletedatEQ applies the EQ predicate on the "completedat" field.
func CompletedatEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCompletedat, v))
}

// CompletedatNEQ applies the NEQ predicate on the "completedat" field.
func CompletedatNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCompletedat, v))
}

// CompletedatIn applies the In predicate on the "completedat" field.
func CompletedatIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCompletedat, vs...))
}

// CompletedatNotIn applies the NotIn predicate on the "completedat" field.
func CompletedatNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCompletedat, vs...))
}

// CompletedatGT applies the GT predicate on the "completedat" field.
func CompletedatGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCompletedat, v))
}

// CompletedatGTE applies the GTE predicate on the "completedat" field.
func CompletedatGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCompletedat, v))
}

// CompletedatLT applies the LT predicate on the "completedat" field.
func CompletedatLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCompletedat, v))
}

// CompletedatLTE applies the LTE predicate on the "completedat" field.
func CompletedatLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCompletedat, v))
}

// CompletedatIsNil applies the IsNil predicate on the "completedat" field.
func CompletedatIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldCompletedat))
}

// CompletedatNotNil applies the NotNil predicate on the "completedat" field.
func CompletedatNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldCompletedat))
}

// TTPsIsNil applies the IsNil predicate on the "TTPs" field.
func TTPsIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldTTPs))
}

// TTPsNotNil applies the NotNil predicate on the "TTPs" field.
func TTPsNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldTTPs))
}

// HasRodent applies the HasEdge predicate on the "rodent" edge.
func HasRodent() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RodentTable, RodentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRodentWith applies the HasEdge predicate on the "rodent" edge with a given conditions (other predicates).
func HasRodentWith(preds ...predicate.Rodent) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RodentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RodentTable, RodentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoot applies the HasEdge predicate on the "loot" edge.
func HasLoot() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LootTable, LootColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLootWith applies the HasEdge predicate on the "loot" edge with a given conditions (other predicates).
func HasLootWith(preds ...predicate.Loot) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LootInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LootTable, LootColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}