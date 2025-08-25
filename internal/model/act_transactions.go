package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"
)

func ActTransactionInsert(data *types.ActTransaction) (int64, error) {

	query, _, _ := meta.Dialect.Insert("act_transactions").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.ActivityDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
