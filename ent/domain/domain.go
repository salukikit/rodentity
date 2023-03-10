// Code generated by ent, DO NOT EDIT.

package domain

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAD holds the string denoting the ad field in the database.
	FieldAD = "ad"
	// FieldOwned holds the string denoting the owned field in the database.
	FieldOwned = "owned"
	// FieldCloud holds the string denoting the cloud field in the database.
	FieldCloud = "cloud"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeChilddomains holds the string denoting the childdomains edge name in mutations.
	EdgeChilddomains = "childdomains"
	// EdgeParentdomain holds the string denoting the parentdomain edge name in mutations.
	EdgeParentdomain = "parentdomain"
	// Table holds the table name of the domain in the database.
	Table = "domains"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "devices"
	// DevicesInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DevicesInverseTable = "devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "domain_devices"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "domain_users"
	// GroupsTable is the table that holds the groups relation/edge.
	GroupsTable = "groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "domain_groups"
	// ChilddomainsTable is the table that holds the childdomains relation/edge.
	ChilddomainsTable = "domains"
	// ChilddomainsColumn is the table column denoting the childdomains relation/edge.
	ChilddomainsColumn = "domain_childdomains"
	// ParentdomainTable is the table that holds the parentdomain relation/edge.
	ParentdomainTable = "domains"
	// ParentdomainColumn is the table column denoting the parentdomain relation/edge.
	ParentdomainColumn = "domain_childdomains"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAD,
	FieldOwned,
	FieldCloud,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "domains"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"domain_childdomains",
}

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
	// DefaultAD holds the default value on creation for the "AD" field.
	DefaultAD bool
	// DefaultOwned holds the default value on creation for the "owned" field.
	DefaultOwned bool
	// DefaultCloud holds the default value on creation for the "cloud" field.
	DefaultCloud string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)
