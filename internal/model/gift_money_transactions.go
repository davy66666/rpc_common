package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"
)

func GiftMoneyTransactionsInsert(data *types.GiftMoneyTransaction) error {

	query, _, _ := meta.Dialect.Insert("gift_money_transactions").Rows(data).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}
