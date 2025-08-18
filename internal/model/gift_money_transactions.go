package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"
	g "github.com/doug-martin/goqu/v9"
)

func GiftMoneyTransactionsFindOne(ex g.Ex) (types.GiftMoneyTransaction, error) {

	var data types.GiftMoneyTransaction
	query, _, _ := meta.Dialect.From("gift_money_transactions").Select(GiftMoneyTransactionFields...).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func GiftMoneyTransactionsFindAll(ex g.Ex) ([]*types.GiftMoneyTransaction, error) {

	var data []*types.GiftMoneyTransaction
	query, _, _ := meta.Dialect.From("gift_money_transactions").Select(GiftMoneyTransactionFields...).Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func GiftMoneyTransactionsInsert(data *types.GiftMoneyTransaction) error {

	query, _, _ := meta.Dialect.Insert("gift_money_transactions").Rows(data).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}
