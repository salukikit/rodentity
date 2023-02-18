package rodentity

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/salukikit/rodentity/ent"
	_ "github.com/salukikit/rodentity/ent/runtime"
)

var (
	Client *ent.Client
)

// Creates a simple in memory sqlite3 db, will not persist after closing and so good for testing.
func InitSQLiteMem() (*ent.Client, error) {
	var err error
	Client = &ent.Client{}

	Client, err = ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dialect.SQLite, err)
	}

	// Run the auto migration tool.
	if err = Client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v\n", err)
	}

	return Client, nil

}

func InitSQLite(db string) (*ent.Client, error) {
	var err error
	Client = &ent.Client{}

	Client, err = ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dialect.SQLite, err)
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

	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, pass, host, port, db)

	Client, err = ent.Open(dialect.MySQL, connstr)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dialect.MySQL, err)
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

	Client, err = ent.Open(dialect.Postgres, connstr)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %v\n", dialect.Postgres, err)
	}

	// Run the auto migration tool.
	if err = Client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v\n", err)
	}

	return Client, nil

}
