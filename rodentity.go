package rodentity

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/salukikit/rodentity/ent"
)

var (
	Client *ent.Client
)

// Creates a simple in memory sqlite3 db,will not persist after closing and so good for testing.
func InitSQLite() (*ent.Client, error) {
	var err error
	dbtype := "sqlite3"
	Client = &ent.Client{}

	Client, err = ent.Open(dbtype, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dbtype, err)
	}

	// Run the auto migration tool.
	if err = Client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v\n", err)
	}

	return Client, nil

}

func InitMySQL(user, pass, host, port, db string) (*ent.Client, error) {
	var err error
	Client = &ent.Client{}
	dbtype := "mysql"

	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, pass, host, port, db)

	Client, err = ent.Open(dbtype, connstr)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dbtype, err)
	}

	// Run the auto migration tool.
	if err = Client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v\n", err)
	}

	return Client, nil

}

// Connects to a new PostGresDB and creates a new schema
func InitPostGres(user, pass, host, port, db string) (*ent.Client, error) {
	var err error
	Client = &ent.Client{}

	connstr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, db, pass)

	dbtype := "postgres"

	Client, err = ent.Open(dbtype, connstr)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dbtype, err)
	}

	// Run the auto migration tool.
	if err = Client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v\n", err)
	}

	return Client, nil

}
