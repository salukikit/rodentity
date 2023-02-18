// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hostname", Type: field.TypeString},
		{Name: "os", Type: field.TypeString, Default: "unknown"},
		{Name: "arch", Type: field.TypeString, Default: "unknown"},
		{Name: "version", Type: field.TypeString, Default: "unknown"},
		{Name: "localaddress", Type: field.TypeString, Default: "unknown"},
		{Name: "machinepass", Type: field.TypeString, Nullable: true},
		{Name: "certificates", Type: field.TypeString, Nullable: true},
		{Name: "domain_devices", Type: field.TypeInt, Nullable: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_domains_devices",
				Columns:    []*schema.Column{DevicesColumns[8]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DomainsColumns holds the columns for the "domains" table.
	DomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "ad", Type: field.TypeBool, Default: true},
		{Name: "owned", Type: field.TypeBool, Default: false},
		{Name: "cloud", Type: field.TypeString, Default: "unknown"},
		{Name: "domain_childdomains", Type: field.TypeInt, Nullable: true},
	}
	// DomainsTable holds the schema information for the "domains" table.
	DomainsTable = &schema.Table{
		Name:       "domains",
		Columns:    DomainsColumns,
		PrimaryKey: []*schema.Column{DomainsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "domains_domains_childdomains",
				Columns:    []*schema.Column{DomainsColumns[5]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: "unknown"},
		{Name: "permissions", Type: field.TypeString, Default: "unknown"},
		{Name: "domain_groups", Type: field.TypeInt, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_domains_groups",
				Columns:    []*schema.Column{GroupsColumns[4]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LootsColumns holds the columns for the "loots" table.
	LootsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "xid", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"cred", "key", "cert", "enum", "objective", "pii", "other"}},
		{Name: "location", Type: field.TypeString, Default: "unknown"},
		{Name: "data", Type: field.TypeBytes},
		{Name: "collectedon", Type: field.TypeTime},
		{Name: "rodent_loot", Type: field.TypeInt, Nullable: true},
		{Name: "task_loot", Type: field.TypeInt, Nullable: true},
	}
	// LootsTable holds the schema information for the "loots" table.
	LootsTable = &schema.Table{
		Name:       "loots",
		Columns:    LootsColumns,
		PrimaryKey: []*schema.Column{LootsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "loots_rodents_loot",
				Columns:    []*schema.Column{LootsColumns[6]},
				RefColumns: []*schema.Column{RodentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "loots_tasks_loot",
				Columns:    []*schema.Column{LootsColumns[7]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OperatorsColumns holds the columns for the "operators" table.
	OperatorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "privkey", Type: field.TypeBytes, Nullable: true},
		{Name: "cert", Type: field.TypeBytes, Nullable: true},
	}
	// OperatorsTable holds the schema information for the "operators" table.
	OperatorsTable = &schema.Table{
		Name:       "operators",
		Columns:    OperatorsColumns,
		PrimaryKey: []*schema.Column{OperatorsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// RodentsColumns holds the columns for the "rodents" table.
	RodentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "xid", Type: field.TypeString},
		{Name: "type", Type: field.TypeString, Default: "FancyRat"},
		{Name: "codename", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "usercontext", Type: field.TypeString, Nullable: true},
		{Name: "beacontime", Type: field.TypeString, Nullable: true},
		{Name: "burned", Type: field.TypeBool, Default: false},
		{Name: "alive", Type: field.TypeBool, Default: true},
		{Name: "joined", Type: field.TypeTime},
		{Name: "lastseen", Type: field.TypeTime},
		{Name: "device_rodents", Type: field.TypeInt, Nullable: true},
		{Name: "project_rodents", Type: field.TypeInt, Nullable: true},
		{Name: "router_rodents", Type: field.TypeInt, Nullable: true},
		{Name: "user_rodents", Type: field.TypeInt, Nullable: true},
	}
	// RodentsTable holds the schema information for the "rodents" table.
	RodentsTable = &schema.Table{
		Name:       "rodents",
		Columns:    RodentsColumns,
		PrimaryKey: []*schema.Column{RodentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rodents_devices_rodents",
				Columns:    []*schema.Column{RodentsColumns[11]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "rodents_projects_rodents",
				Columns:    []*schema.Column{RodentsColumns[12]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "rodents_routers_rodents",
				Columns:    []*schema.Column{RodentsColumns[13]},
				RefColumns: []*schema.Column{RoutersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "rodents_users_rodents",
				Columns:    []*schema.Column{RodentsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoutersColumns holds the columns for the "routers" table.
	RoutersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "privkey", Type: field.TypeBytes, Nullable: true},
		{Name: "cert", Type: field.TypeBytes, Nullable: true},
		{Name: "commands", Type: field.TypeJSON, Nullable: true},
	}
	// RoutersTable holds the schema information for the "routers" table.
	RoutersTable = &schema.Table{
		Name:       "routers",
		Columns:    RoutersColumns,
		PrimaryKey: []*schema.Column{RoutersColumns[0]},
	}
	// SubnetsColumns holds the columns for the "subnets" table.
	SubnetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cidr", Type: field.TypeString},
		{Name: "mask", Type: field.TypeBytes, Nullable: true},
	}
	// SubnetsTable holds the schema information for the "subnets" table.
	SubnetsTable = &schema.Table{
		Name:       "subnets",
		Columns:    SubnetsColumns,
		PrimaryKey: []*schema.Column{SubnetsColumns[0]},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "xid", Type: field.TypeString},
		{Name: "type", Type: field.TypeString, Default: "cmd"},
		{Name: "args", Type: field.TypeJSON, Nullable: true},
		{Name: "data", Type: field.TypeBytes, Nullable: true},
		{Name: "result", Type: field.TypeBytes, Nullable: true},
		{Name: "executed", Type: field.TypeBool, Default: false},
		{Name: "looted", Type: field.TypeBool, Default: false},
		{Name: "requestedat", Type: field.TypeTime},
		{Name: "completedat", Type: field.TypeTime, Nullable: true},
		{Name: "tt_ps", Type: field.TypeJSON, Nullable: true},
		{Name: "rodent_tasks", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_rodents_tasks",
				Columns:    []*schema.Column{TasksColumns[11]},
				RefColumns: []*schema.Column{RodentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "givenname", Type: field.TypeString, Default: "unknown"},
		{Name: "email", Type: field.TypeString, Default: "unknown"},
		{Name: "owned", Type: field.TypeBool, Default: false},
		{Name: "age", Type: field.TypeString, Nullable: true, Default: "unknown"},
		{Name: "homedir", Type: field.TypeString, Default: "unknown"},
		{Name: "enabled", Type: field.TypeBool, Default: true},
		{Name: "domain_users", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_domains_users",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DeviceUsersColumns holds the columns for the "device_users" table.
	DeviceUsersColumns = []*schema.Column{
		{Name: "device_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// DeviceUsersTable holds the schema information for the "device_users" table.
	DeviceUsersTable = &schema.Table{
		Name:       "device_users",
		Columns:    DeviceUsersColumns,
		PrimaryKey: []*schema.Column{DeviceUsersColumns[0], DeviceUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "device_users_device_id",
				Columns:    []*schema.Column{DeviceUsersColumns[0]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "device_users_user_id",
				Columns:    []*schema.Column{DeviceUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupDevicesColumns holds the columns for the "group_devices" table.
	GroupDevicesColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "device_id", Type: field.TypeInt},
	}
	// GroupDevicesTable holds the schema information for the "group_devices" table.
	GroupDevicesTable = &schema.Table{
		Name:       "group_devices",
		Columns:    GroupDevicesColumns,
		PrimaryKey: []*schema.Column{GroupDevicesColumns[0], GroupDevicesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_devices_group_id",
				Columns:    []*schema.Column{GroupDevicesColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_devices_device_id",
				Columns:    []*schema.Column{GroupDevicesColumns[1]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectOperatorsColumns holds the columns for the "project_operators" table.
	ProjectOperatorsColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeInt},
		{Name: "operator_id", Type: field.TypeInt},
	}
	// ProjectOperatorsTable holds the schema information for the "project_operators" table.
	ProjectOperatorsTable = &schema.Table{
		Name:       "project_operators",
		Columns:    ProjectOperatorsColumns,
		PrimaryKey: []*schema.Column{ProjectOperatorsColumns[0], ProjectOperatorsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_operators_project_id",
				Columns:    []*schema.Column{ProjectOperatorsColumns[0]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_operators_operator_id",
				Columns:    []*schema.Column{ProjectOperatorsColumns[1]},
				RefColumns: []*schema.Column{OperatorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubnetHostsColumns holds the columns for the "subnet_hosts" table.
	SubnetHostsColumns = []*schema.Column{
		{Name: "subnet_id", Type: field.TypeInt},
		{Name: "device_id", Type: field.TypeInt},
	}
	// SubnetHostsTable holds the schema information for the "subnet_hosts" table.
	SubnetHostsTable = &schema.Table{
		Name:       "subnet_hosts",
		Columns:    SubnetHostsColumns,
		PrimaryKey: []*schema.Column{SubnetHostsColumns[0], SubnetHostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subnet_hosts_subnet_id",
				Columns:    []*schema.Column{SubnetHostsColumns[0]},
				RefColumns: []*schema.Column{SubnetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subnet_hosts_device_id",
				Columns:    []*schema.Column{SubnetHostsColumns[1]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DevicesTable,
		DomainsTable,
		GroupsTable,
		LootsTable,
		OperatorsTable,
		ProjectsTable,
		RodentsTable,
		RoutersTable,
		SubnetsTable,
		TasksTable,
		UsersTable,
		DeviceUsersTable,
		GroupDevicesTable,
		GroupUsersTable,
		ProjectOperatorsTable,
		SubnetHostsTable,
	}
)

func init() {
	DevicesTable.ForeignKeys[0].RefTable = DomainsTable
	DomainsTable.ForeignKeys[0].RefTable = DomainsTable
	GroupsTable.ForeignKeys[0].RefTable = DomainsTable
	LootsTable.ForeignKeys[0].RefTable = RodentsTable
	LootsTable.ForeignKeys[1].RefTable = TasksTable
	RodentsTable.ForeignKeys[0].RefTable = DevicesTable
	RodentsTable.ForeignKeys[1].RefTable = ProjectsTable
	RodentsTable.ForeignKeys[2].RefTable = RoutersTable
	RodentsTable.ForeignKeys[3].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = RodentsTable
	UsersTable.ForeignKeys[0].RefTable = DomainsTable
	DeviceUsersTable.ForeignKeys[0].RefTable = DevicesTable
	DeviceUsersTable.ForeignKeys[1].RefTable = UsersTable
	GroupDevicesTable.ForeignKeys[0].RefTable = GroupsTable
	GroupDevicesTable.ForeignKeys[1].RefTable = DevicesTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	ProjectOperatorsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectOperatorsTable.ForeignKeys[1].RefTable = OperatorsTable
	SubnetHostsTable.ForeignKeys[0].RefTable = SubnetsTable
	SubnetHostsTable.ForeignKeys[1].RefTable = DevicesTable
}
