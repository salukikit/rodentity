// Code generated by ent, DO NOT EDIT.

package group

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeDomain holds the string denoting the domain edge name in mutations.
	EdgeDomain = "domain"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// DevicesTable is the table that holds the devices relation/edge. The primary key declared below.
	DevicesTable = "group_devices"
	// DevicesInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DevicesInverseTable = "devices"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "group_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// DomainTable is the table that holds the domain relation/edge.
	DomainTable = "groups"
	// DomainInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	DomainInverseTable = "domains"
	// DomainColumn is the table column denoting the domain relation/edge.
	DomainColumn = "domain_groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPermissions,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"domain_groups",
}

var (
	// DevicesPrimaryKey and DevicesColumn2 are the table columns denoting the
	// primary key for the devices relation (M2M).
	DevicesPrimaryKey = []string{"group_id", "device_id"}
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"group_id", "user_id"}
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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultPermissions holds the default value on creation for the "permissions" field.
	DefaultPermissions string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)
