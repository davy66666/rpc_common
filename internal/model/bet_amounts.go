package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func BetAmountFindOne(ex g.Ex) (types.BetAmount, error) {

	var data types.BetAmount
	query, _, _ := meta.Dialect.From("bet_amounts").Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func BetAmountUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("bet_amounts").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}

func BetAmountInsert(data *types.BetAmount) (int64, error) {

	query, _, _ := meta.Dialect.Insert("bet_amounts").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.SqlxDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
