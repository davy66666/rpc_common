package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func ActivityReportFindOne(ex g.Ex) (types.ActivityReport, error) {

	var data types.ActivityReport
	query, _, _ := meta.Dialect.From("activity_reports").Select(ActivityReportFields...).Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.ActivityDb.Get(&data, query)

	return data, err
}

func ActivityReportUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("activity_reports").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.ActivityDb.Exec(query)

	return err
}

func ActivityReportInsert(data *types.ActivityReport) (int64, error) {

	query, _, _ := meta.Dialect.Insert("activity_reports").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.ActivityDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
