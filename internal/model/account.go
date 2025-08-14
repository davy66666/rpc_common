package model

import (
	"fmt"

	g "github.com/doug-martin/goqu/v9"
)

func AccountUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("accounts").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}
