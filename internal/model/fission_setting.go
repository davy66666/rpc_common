package model

import (
	"context"
	"errors"
	"fmt"
	json "github.com/bytedance/sonic"
	"github.com/davy66666/rpc_service/common/utils/strx"
	"github.com/davy66666/rpc_service/internal/types"
	"github.com/redis/go-redis/v9"

	g "github.com/doug-martin/goqu/v9"
)

func RedisGetFissionSetting(ctx context.Context) ([]*types.FissionSetting, error) {

	var data []*types.FissionSetting
	result, err := meta.Rds.Get(ctx, SysKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return data, err
	}

	if result != "" {
		err = json.Unmarshal([]byte(result), &data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	query, _, _ := meta.Dialect.From("fission_setting").Select(FissionSettingFields...).Where(g.Ex{"is_open": 1}).ToSQL()
	fmt.Println(query)
	err = meta.SqlxDb.Select(&data, query)
	if err != nil {
		return data, err
	}

	_, err = meta.Rds.Set(ctx, SysKey, strx.Any2Str(data), -1).Result()
	if err != nil {
		return data, err
	}

	return data, nil
}
