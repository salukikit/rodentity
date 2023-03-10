// Code generated by ent, DO NOT EDIT.

package router

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the router type in the database.
	Label = "router"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRname holds the string denoting the rname field in the database.
	FieldRname = "rname"
	// FieldPrivkey holds the string denoting the privkey field in the database.
	FieldPrivkey = "privkey"
	// FieldCert holds the string denoting the cert field in the database.
	FieldCert = "cert"
	// FieldCommands holds the string denoting the commands field in the database.
	FieldCommands = "commands"
	// FieldInterfaces holds the string denoting the interfaces field in the database.
	FieldInterfaces = "interfaces"
	// EdgeRodents holds the string denoting the rodents edge name in mutations.
	EdgeRodents = "rodents"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the router in the database.
	Table = "routers"
	// RodentsTable is the table that holds the rodents relation/edge. The primary key declared below.
	RodentsTable = "router_rodents"
	// RodentsInverseTable is the table name for the Rodent entity.
	// It exists in this package in order to avoid circular dependency with the "rodent" package.
	RodentsInverseTable = "rodents"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "routers"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_routers"
)

// Columns holds all SQL columns for router fields.
var Columns = []string{
	FieldID,
	FieldRname,
	FieldPrivkey,
	FieldCert,
	FieldCommands,
	FieldInterfaces,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "routers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_routers",
}

var (
	// RodentsPrimaryKey and RodentsColumn2 are the table columns denoting the
	// primary key for the rodents relation (M2M).
	RodentsPrimaryKey = []string{"router_id", "rodent_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)
