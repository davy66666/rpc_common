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
