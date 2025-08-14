package model

import (
	"context"
	"fmt"
	json "github.com/bytedance/sonic"
	g "github.com/doug-martin/goqu/v9"
	"ks_api_service/common/utils/strx"
	"ks_api_service/internal/types"
	"strconv"
)

func UserFindOne(ex g.Ex) (types.User, error) {

	var data types.User
	query, _, _ := meta.Dialect.From("users").Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func GetUserById(ctx context.Context, id int64) (types.User, error) {

	key := fmt.Sprintf("user_%d", id)
	var data types.User
	result, err := meta.Rds.Get(ctx, key).Result()
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

	query, _, _ := meta.Dialect.From("users").Where(g.Ex{"id": id}).Limit(1).ToSQL()
	fmt.Println(query)
	err = meta.SqlxDb.Get(&data, query)
	if err != nil {
		return data, err
	}

	_, err = meta.Rds.Set(ctx, key, strx.Any2Str(data), -1).Result()
	if err != nil {
		return data, err
	}

	return data, err
}

func UserUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("users").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}

func EsUser(ctx context.Context, id int64) error {

	var data types.User
	query, _, _ := meta.Dialect.From("users").Where(g.Ex{"id": id}).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)
	if err != nil {
		return err
	}

	// 写入
	_, err = meta.EsClient.Index().
		Index("users").
		Id(strconv.FormatInt(id, 10)).
		BodyJson(data).
		Refresh("wait_for"). //sync方式的写入
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
