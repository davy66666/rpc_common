package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func ActTransactionFindOne(ex g.Ex) (types.ActTransaction, error) {

	var data types.ActTransaction
	query, _, _ := meta.Dialect.From("act_transactions").Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.ActivityDb.Get(&data, query)

	return data, err
}

func ActTransactionUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("act_transactions").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.ActivityDb.Exec(query)

	return err
}

func ActTransactionInsert(data *types.ActTransaction) (int64, error) {

	query, _, _ := meta.Dialect.Insert("act_transactions").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.ActivityDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
