package rodentity

import (
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/salukikit/rodentity/ent"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/task"
)

func GetRodentbyXid(ctx context.Context, client *ent.Client, rodentid string) (*ent.Rodent, error) {
	r, err := client.Rodent.
		Query().
		Where(rodent.Xid(rodentid)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying rodent: %w", err)
	}

	return r, nil
}

func GetAllRodentTasks(ctx context.Context, client *ent.Client, rodentid string) ([]*ent.Task, error) {
	rtasks, err := client.Rodent.Query().Where(rodent.Xid(rodentid)).QueryTasks().All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying tasks for rodent %s: %w\n", rodentid, err)
	}

	return rtasks, nil
}

func QueryRodentOpenTasks(ctx context.Context, client *ent.Client, rodentid string) ([]*ent.Task, error) {
	rtasks, err := client.Rodent.Query().Where(rodent.Xid(rodentid)).QueryTasks().Where(task.Executed(false)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying open tasks for rodent %s: %w\n", rodentid, err)
	}

	return rtasks, nil
}
