package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"
	"strconv"

	g "github.com/doug-martin/goqu/v9"
)

func GetFissionExclusiveRewardFirst(exclusiveRewardsType int, userLevel string, payLevel int64) (types.FissionExclusiveReward, error) {

	var data types.FissionExclusiveReward
	ds := meta.Dialect.From("fission_exclusive_rewards").Where(g.Ex{"type": exclusiveRewardsType}).Select(FissionExclusiveRewardFields...)
	// 添加 FIND_IN_SET 条件
	var level string
	if exclusiveRewardsType == 1 {
		level = userLevel
	} else {
		level = strconv.FormatInt(payLevel, 10)
	}
	query, _, _ := ds.Where(g.L("FIND_IN_SET(?, ext_id)", level)).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}
