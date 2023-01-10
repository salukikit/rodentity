// Code generated by ent, DO NOT EDIT.

package rodent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/salukikit/rodentity/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldID, id))
}

// Xid applies equality check predicate on the "xid" field. It's identical to XidEQ.
func Xid(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldXid, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldType, v))
}

// Codename applies equality check predicate on the "codename" field. It's identical to CodenameEQ.
func Codename(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldCodename, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldKey, v))
}

// Usercontext applies equality check predicate on the "usercontext" field. It's identical to UsercontextEQ.
func Usercontext(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldUsercontext, v))
}

// Beacontime applies equality check predicate on the "beacontime" field. It's identical to BeacontimeEQ.
func Beacontime(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldBeacontime, v))
}

// Burned applies equality check predicate on the "burned" field. It's identical to BurnedEQ.
func Burned(v bool) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldBurned, v))
}

// Alive applies equality check predicate on the "alive" field. It's identical to AliveEQ.
func Alive(v bool) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldAlive, v))
}

// Joined applies equality check predicate on the "joined" field. It's identical to JoinedEQ.
func Joined(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldJoined, v))
}

// Lastseen applies equality check predicate on the "lastseen" field. It's identical to LastseenEQ.
func Lastseen(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldLastseen, v))
}

// XidEQ applies the EQ predicate on the "xid" field.
func XidEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldXid, v))
}

// XidNEQ applies the NEQ predicate on the "xid" field.
func XidNEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldXid, v))
}

// XidIn applies the In predicate on the "xid" field.
func XidIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldXid, vs...))
}

// XidNotIn applies the NotIn predicate on the "xid" field.
func XidNotIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldXid, vs...))
}

// XidGT applies the GT predicate on the "xid" field.
func XidGT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldXid, v))
}

// XidGTE applies the GTE predicate on the "xid" field.
func XidGTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldXid, v))
}

// XidLT applies the LT predicate on the "xid" field.
func XidLT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldXid, v))
}

// XidLTE applies the LTE predicate on the "xid" field.
func XidLTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldXid, v))
}

// XidContains applies the Contains predicate on the "xid" field.
func XidContains(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContains(FieldXid, v))
}

// XidHasPrefix applies the HasPrefix predicate on the "xid" field.
func XidHasPrefix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasPrefix(FieldXid, v))
}

// XidHasSuffix applies the HasSuffix predicate on the "xid" field.
func XidHasSuffix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasSuffix(FieldXid, v))
}

// XidEqualFold applies the EqualFold predicate on the "xid" field.
func XidEqualFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEqualFold(FieldXid, v))
}

// XidContainsFold applies the ContainsFold predicate on the "xid" field.
func XidContainsFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContainsFold(FieldXid, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContainsFold(FieldType, v))
}

// CodenameEQ applies the EQ predicate on the "codename" field.
func CodenameEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldCodename, v))
}

// CodenameNEQ applies the NEQ predicate on the "codename" field.
func CodenameNEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldCodename, v))
}

// CodenameIn applies the In predicate on the "codename" field.
func CodenameIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldCodename, vs...))
}

// CodenameNotIn applies the NotIn predicate on the "codename" field.
func CodenameNotIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldCodename, vs...))
}

// CodenameGT applies the GT predicate on the "codename" field.
func CodenameGT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldCodename, v))
}

// CodenameGTE applies the GTE predicate on the "codename" field.
func CodenameGTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldCodename, v))
}

// CodenameLT applies the LT predicate on the "codename" field.
func CodenameLT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldCodename, v))
}

// CodenameLTE applies the LTE predicate on the "codename" field.
func CodenameLTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldCodename, v))
}

// CodenameContains applies the Contains predicate on the "codename" field.
func CodenameContains(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContains(FieldCodename, v))
}

// CodenameHasPrefix applies the HasPrefix predicate on the "codename" field.
func CodenameHasPrefix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasPrefix(FieldCodename, v))
}

// CodenameHasSuffix applies the HasSuffix predicate on the "codename" field.
func CodenameHasSuffix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasSuffix(FieldCodename, v))
}

// CodenameEqualFold applies the EqualFold predicate on the "codename" field.
func CodenameEqualFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEqualFold(FieldCodename, v))
}

// CodenameContainsFold applies the ContainsFold predicate on the "codename" field.
func CodenameContainsFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContainsFold(FieldCodename, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContainsFold(FieldKey, v))
}

// UsercontextEQ applies the EQ predicate on the "usercontext" field.
func UsercontextEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldUsercontext, v))
}

// UsercontextNEQ applies the NEQ predicate on the "usercontext" field.
func UsercontextNEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldUsercontext, v))
}

// UsercontextIn applies the In predicate on the "usercontext" field.
func UsercontextIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldUsercontext, vs...))
}

// UsercontextNotIn applies the NotIn predicate on the "usercontext" field.
func UsercontextNotIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldUsercontext, vs...))
}

// UsercontextGT applies the GT predicate on the "usercontext" field.
func UsercontextGT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldUsercontext, v))
}

// UsercontextGTE applies the GTE predicate on the "usercontext" field.
func UsercontextGTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldUsercontext, v))
}

// UsercontextLT applies the LT predicate on the "usercontext" field.
func UsercontextLT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldUsercontext, v))
}

// UsercontextLTE applies the LTE predicate on the "usercontext" field.
func UsercontextLTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldUsercontext, v))
}

// UsercontextContains applies the Contains predicate on the "usercontext" field.
func UsercontextContains(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContains(FieldUsercontext, v))
}

// UsercontextHasPrefix applies the HasPrefix predicate on the "usercontext" field.
func UsercontextHasPrefix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasPrefix(FieldUsercontext, v))
}

// UsercontextHasSuffix applies the HasSuffix predicate on the "usercontext" field.
func UsercontextHasSuffix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasSuffix(FieldUsercontext, v))
}

// UsercontextIsNil applies the IsNil predicate on the "usercontext" field.
func UsercontextIsNil() predicate.Rodent {
	return predicate.Rodent(sql.FieldIsNull(FieldUsercontext))
}

// UsercontextNotNil applies the NotNil predicate on the "usercontext" field.
func UsercontextNotNil() predicate.Rodent {
	return predicate.Rodent(sql.FieldNotNull(FieldUsercontext))
}

// UsercontextEqualFold applies the EqualFold predicate on the "usercontext" field.
func UsercontextEqualFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEqualFold(FieldUsercontext, v))
}

// UsercontextContainsFold applies the ContainsFold predicate on the "usercontext" field.
func UsercontextContainsFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContainsFold(FieldUsercontext, v))
}

// BeacontimeEQ applies the EQ predicate on the "beacontime" field.
func BeacontimeEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldBeacontime, v))
}

// BeacontimeNEQ applies the NEQ predicate on the "beacontime" field.
func BeacontimeNEQ(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldBeacontime, v))
}

// BeacontimeIn applies the In predicate on the "beacontime" field.
func BeacontimeIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldBeacontime, vs...))
}

// BeacontimeNotIn applies the NotIn predicate on the "beacontime" field.
func BeacontimeNotIn(vs ...string) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldBeacontime, vs...))
}

// BeacontimeGT applies the GT predicate on the "beacontime" field.
func BeacontimeGT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldBeacontime, v))
}

// BeacontimeGTE applies the GTE predicate on the "beacontime" field.
func BeacontimeGTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldBeacontime, v))
}

// BeacontimeLT applies the LT predicate on the "beacontime" field.
func BeacontimeLT(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldBeacontime, v))
}

// BeacontimeLTE applies the LTE predicate on the "beacontime" field.
func BeacontimeLTE(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldBeacontime, v))
}

// BeacontimeContains applies the Contains predicate on the "beacontime" field.
func BeacontimeContains(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContains(FieldBeacontime, v))
}

// BeacontimeHasPrefix applies the HasPrefix predicate on the "beacontime" field.
func BeacontimeHasPrefix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasPrefix(FieldBeacontime, v))
}

// BeacontimeHasSuffix applies the HasSuffix predicate on the "beacontime" field.
func BeacontimeHasSuffix(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldHasSuffix(FieldBeacontime, v))
}

// BeacontimeIsNil applies the IsNil predicate on the "beacontime" field.
func BeacontimeIsNil() predicate.Rodent {
	return predicate.Rodent(sql.FieldIsNull(FieldBeacontime))
}

// BeacontimeNotNil applies the NotNil predicate on the "beacontime" field.
func BeacontimeNotNil() predicate.Rodent {
	return predicate.Rodent(sql.FieldNotNull(FieldBeacontime))
}

// BeacontimeEqualFold applies the EqualFold predicate on the "beacontime" field.
func BeacontimeEqualFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldEqualFold(FieldBeacontime, v))
}

// BeacontimeContainsFold applies the ContainsFold predicate on the "beacontime" field.
func BeacontimeContainsFold(v string) predicate.Rodent {
	return predicate.Rodent(sql.FieldContainsFold(FieldBeacontime, v))
}

// BurnedEQ applies the EQ predicate on the "burned" field.
func BurnedEQ(v bool) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldBurned, v))
}

// BurnedNEQ applies the NEQ predicate on the "burned" field.
func BurnedNEQ(v bool) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldBurned, v))
}

// AliveEQ applies the EQ predicate on the "alive" field.
func AliveEQ(v bool) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldAlive, v))
}

// AliveNEQ applies the NEQ predicate on the "alive" field.
func AliveNEQ(v bool) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldAlive, v))
}

// JoinedEQ applies the EQ predicate on the "joined" field.
func JoinedEQ(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldJoined, v))
}

// JoinedNEQ applies the NEQ predicate on the "joined" field.
func JoinedNEQ(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldJoined, v))
}

// JoinedIn applies the In predicate on the "joined" field.
func JoinedIn(vs ...time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldJoined, vs...))
}

// JoinedNotIn applies the NotIn predicate on the "joined" field.
func JoinedNotIn(vs ...time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldJoined, vs...))
}

// JoinedGT applies the GT predicate on the "joined" field.
func JoinedGT(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldJoined, v))
}

// JoinedGTE applies the GTE predicate on the "joined" field.
func JoinedGTE(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldJoined, v))
}

// JoinedLT applies the LT predicate on the "joined" field.
func JoinedLT(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldJoined, v))
}

// JoinedLTE applies the LTE predicate on the "joined" field.
func JoinedLTE(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldJoined, v))
}

// LastseenEQ applies the EQ predicate on the "lastseen" field.
func LastseenEQ(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldEQ(FieldLastseen, v))
}

// LastseenNEQ applies the NEQ predicate on the "lastseen" field.
func LastseenNEQ(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldNEQ(FieldLastseen, v))
}

// LastseenIn applies the In predicate on the "lastseen" field.
func LastseenIn(vs ...time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldIn(FieldLastseen, vs...))
}

// LastseenNotIn applies the NotIn predicate on the "lastseen" field.
func LastseenNotIn(vs ...time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldNotIn(FieldLastseen, vs...))
}

// LastseenGT applies the GT predicate on the "lastseen" field.
func LastseenGT(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldGT(FieldLastseen, v))
}

// LastseenGTE applies the GTE predicate on the "lastseen" field.
func LastseenGTE(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldGTE(FieldLastseen, v))
}

// LastseenLT applies the LT predicate on the "lastseen" field.
func LastseenLT(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldLT(FieldLastseen, v))
}

// LastseenLTE applies the LTE predicate on the "lastseen" field.
func LastseenLTE(v time.Time) predicate.Rodent {
	return predicate.Rodent(sql.FieldLTE(FieldLastseen, v))
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeviceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoot applies the HasEdge predicate on the "loot" edge.
func HasLoot() predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LootTable, LootPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLootWith applies the HasEdge predicate on the "loot" edge with a given conditions (other predicates).
func HasLootWith(preds ...predicate.Loot) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LootInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LootTable, LootPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rodent) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rodent) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
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
func Not(p predicate.Rodent) predicate.Rodent {
	return predicate.Rodent(func(s *sql.Selector) {
		p(s.Not())
	})
}
