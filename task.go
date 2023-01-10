package rodentity

import (
	"context"
	"fmt"

	"github.com/salukikit/rodentity/ent"
	"github.com/salukikit/rodentity/ent/rodent"
)

func GetAllTasks() {

}

func UpdateTask() {

}

func GetRodentTasks(ctx context.Context, client *ent.Client, rodentid string) (*ent.Rodent, error) {
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

func QueryRodentOpenTasks(ctx context.Context, client *ent.Client, rodentid string) (*ent.Rodent, error) {
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
