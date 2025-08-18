package model

import (
	"database/sql"
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func FissionRewardFindOne(ex g.Ex) (types.FissionReward, error) {

	var data types.FissionReward
	query, _, _ := meta.Dialect.From("fission_reward").Select(FissionRewardFields...).Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func FissionRewardInsert(data *types.FissionReward) (int64, error) {

	query, _, _ := meta.Dialect.Insert("fission_reward").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.SqlxDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func GetFissionReward(ex g.Ex) ([]*types.FissionReward, error) {

	var data []*types.FissionReward
	query, _, _ := meta.Dialect.From("fission_reward").Select(FissionRewardFields...).Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func FissionRewardCount(ex g.Ex) (int, error) {

	n := 0
	query, _, _ := meta.Dialect.From("fission_reward").Select(g.COUNT("id")).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&n, query)
	if err == sql.ErrNoRows {
		return n, nil
	}

	return n, err
}
