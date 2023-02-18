// Code generated by ent, DO NOT EDIT.

package project

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeOperators holds the string denoting the operators edge name in mutations.
	EdgeOperators = "operators"
	// EdgeRodents holds the string denoting the rodents edge name in mutations.
	EdgeRodents = "rodents"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// OperatorsTable is the table that holds the operators relation/edge. The primary key declared below.
	OperatorsTable = "project_operators"
	// OperatorsInverseTable is the table name for the Operator entity.
	// It exists in this package in order to avoid circular dependency with the "operator" package.
	OperatorsInverseTable = "operators"
	// RodentsTable is the table that holds the rodents relation/edge.
	RodentsTable = "rodents"
	// RodentsInverseTable is the table name for the Rodent entity.
	// It exists in this package in order to avoid circular dependency with the "rodent" package.
	RodentsInverseTable = "rodents"
	// RodentsColumn is the table column denoting the rodents relation/edge.
	RodentsColumn = "project_rodents"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
}

var (
	// OperatorsPrimaryKey and OperatorsColumn2 are the table columns denoting the
	// primary key for the operators relation (M2M).
	OperatorsPrimaryKey = []string{"project_id", "operator_id"}
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