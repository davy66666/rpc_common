package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"
)

func LevelUpgradeLogInsert(data *types.LevelUpgradeLog) (int64, error) {

	query, _, _ := meta.Dialect.Insert("level_upgrade_log").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.LogDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
