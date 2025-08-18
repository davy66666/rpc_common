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

func ActivityUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("activitys").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.ActivityDb.Exec(query)

	return err
}

func ActivityInsert(data *types.Activity) (int64, error) {

	query, _, _ := meta.Dialect.Insert("activitys").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.ActivityDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
