package model

import (
	"fmt"
	"ks_api_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func TeamUserFindOne(ex g.Ex) (types.TeamUser, error) {

	var data types.TeamUser
	query, _, _ := meta.Dialect.From("team_users").Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func TeamUserUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("team_users").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}

func TeamUserInsert(data *types.TeamUser) (int64, error) {

	query, _, _ := meta.Dialect.Insert("team_users").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.SqlxDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
