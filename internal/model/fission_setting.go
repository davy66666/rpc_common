package model

import (
	"context"
	"fmt"
	json "github.com/bytedance/sonic"
	"ks_api_service/common/utils/strx"
	"ks_api_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func FissionSettingFindOne(ex g.Ex) (types.FissionSetting, error) {

	var data types.FissionSetting
	query, _, _ := meta.Dialect.From("fission_setting").Where(ex).Order(g.C("id").Asc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func FissionSettingInsert(data *types.FissionSetting) (int64, error) {

	query, _, _ := meta.Dialect.Insert("fission_setting").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.SqlxDb.Exec(query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func GetFissionSetting(ex g.Ex) ([]*types.FissionSetting, error) {

	var data []*types.FissionSetting
	query, _, _ := meta.Dialect.From("fission_setting").Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func RedisGetFissionSetting(ctx context.Context) ([]*types.FissionSetting, error) {

	var data []*types.FissionSetting
	result, err := meta.Rds.Get(ctx, SysKey).Result()
	if err != nil {
		return data, err
	}

	if result != "" {
		err = json.Unmarshal([]byte(result), &data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	query, _, _ := meta.Dialect.From("fission_setting").Where(g.Ex{"is_open": 1}).ToSQL()
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
