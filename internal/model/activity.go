package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func ActivityFindOne(ex g.Ex) (types.Activity, error) {

	var data types.Activity
	query, _, _ := meta.Dialect.From("activity").Select(ActivityFields...).Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.ActivityDb.Get(&data, query)

	return data, err
}
