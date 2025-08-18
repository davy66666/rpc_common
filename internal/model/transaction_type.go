package model

import (
	"context"
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"
	g "github.com/doug-martin/goqu/v9"
	"github.com/jinzhu/copier"
	"strconv"
)

func TransactionTypeFindOne(ex g.Ex) (types.TransactionType, error) {

	var data types.TransactionType
	query, _, _ := meta.Dialect.From("transaction_types").Select(TransactionTypeFields...).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func EsTransaction(ctx context.Context, id int64) error {

	tran, err := TransactionTypeFindOne(g.Ex{"id": id})
	if err != nil {
		return err
	}

	var data types.TransactionTypeEs
	err = copier.Copy(&data, tran)
	if err != nil {
		return err
	}

	data.CreatedAt = tran.CreatedAt.Time.UTC().Format("2006-01-02 15:04:05")
	data.UpdatedAt = tran.CreatedAt.Time.UTC().Format("2006-01-02 15:04:05")
	// 写入
	_, err = meta.EsClient.Index().
		Index("transactions").
		Id(strconv.FormatInt(id, 10)).
		BodyJson(data).
		Refresh("wait_for"). //sync方式的写入
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
