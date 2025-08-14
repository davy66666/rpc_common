package model

import (
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"ks_api_service/internal/types"
)

func GiftMoneyTransactionsFindOne(ex g.Ex) (types.Transaction, error) {

	var data types.Transaction
	query, _, _ := meta.Dialect.From("gift_money_transactions").Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func GiftMoneyTransactionsFindAll(ex g.Ex) ([]*types.Transaction, error) {

	var data []*types.Transaction
	query, _, _ := meta.Dialect.From("gift_money_transactions").Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func GiftMoneyTransactionsInsert(data *types.Transaction) error {

	query, _, _ := meta.Dialect.Insert("gift_money_transactions").Rows(data).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}
