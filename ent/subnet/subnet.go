// Code generated by ent, DO NOT EDIT.

package subnet

const (
	// Label holds the string label denoting the subnet type in the database.
	Label = "subnet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCidr holds the string denoting the cidr field in the database.
	FieldCidr = "cidr"
	// FieldMask holds the string denoting the mask field in the database.
	FieldMask = "mask"
	// EdgeHosts holds the string denoting the hosts edge name in mutations.
	EdgeHosts = "hosts"
	// Table holds the table name of the subnet in the database.
	Table = "subnets"
	// HostsTable is the table that holds the hosts relation/edge. The primary key declared below.
	HostsTable = "subnet_hosts"
	// HostsInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	HostsInverseTable = "devices"
)

// Columns holds all SQL columns for subnet fields.
var Columns = []string{
	FieldID,
	FieldCidr,
	FieldMask,
}

var (
	// HostsPrimaryKey and HostsColumn2 are the table columns denoting the
	// primary key for the hosts relation (M2M).
	HostsPrimaryKey = []string{"subnet_id", "device_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
