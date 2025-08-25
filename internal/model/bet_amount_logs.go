package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func BetAmountLogUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("bet_amount_logs").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.LogDb.Exec(query)

	return err
}

func BetAmountLogInsert(data *types.BetAmountLog) (int64, error) {

	query, _, _ := meta.Dialect.Insert("bet_amount_logs").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.LogDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
