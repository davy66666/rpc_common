package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func TeamUserFindOne(ex g.Ex) (types.TeamUser, error) {

	var data types.TeamUser
	query, _, _ := meta.Dialect.From("team_users").Select(TeamUserFields...).Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.ActivityDb.Get(&data, query)

	return data, err
}

func TeamUserUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("team_users").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.ActivityDb.Exec(query)

	return err
}
