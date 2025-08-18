package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func PayLevelFindOne(ex g.Ex) (types.PayLevel, error) {

	var data types.PayLevel
	query, _, _ := meta.Dialect.From("pay_levels").Select(PayLevelFields...).Where(ex).Order(g.C("id").Desc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func PayLevelUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("pay_levels").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}

func PayLevelInsert(data *types.PayLevel) (int64, error) {

	query, _, _ := meta.Dialect.Insert("pay_levels").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.SqlxDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
