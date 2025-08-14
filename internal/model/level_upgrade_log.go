package model

import (
	"fmt"
	"ks_api_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func LevelUpgradeLogFindOne(ex g.Ex) (types.LevelUpgradeLog, error) {

	var data types.LevelUpgradeLog
	query, _, _ := meta.Dialect.From("clear_bet_amount_logs").Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.LogDb.Get(&data, query)

	return data, err
}

func LevelUpgradeLogInsert(data *types.LevelUpgradeLog) (int64, error) {

	query, _, _ := meta.Dialect.Insert("clear_bet_amount_logs").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.LogDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func GetLevelUpgradeLog(ex g.Ex) ([]*types.LevelUpgradeLog, error) {

	var data []*types.LevelUpgradeLog
	query, _, _ := meta.Dialect.From("clear_bet_amount_logs").Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.LogDb.Select(&data, query)

	return data, err
}
