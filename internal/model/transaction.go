package model

import (
	"context"
	"fmt"
	json "github.com/bytedance/sonic"
	"github.com/davy66666/rpc_service/internal/types"
	g "github.com/doug-martin/goqu/v9"
	"github.com/olivere/elastic/v7"
)

func TransactionFindAll(ex g.Ex) ([]*types.Transaction, error) {

	var data []*types.Transaction
	query, _, _ := meta.Dialect.From("transactions").Select(TransactionFields...).Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func TransactionLastOne(ex g.Ex) (types.Transaction, error) {

	var data types.Transaction
	query, _, _ := meta.Dialect.From("transactions").Select(TransactionFields...).Where(ex).Order(g.C("id").Desc()).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func GetTransaction(ex g.Ex) ([]*types.Transaction, error) {

	var data []*types.Transaction
	query, _, _ := meta.Dialect.From("transactions").Select(TransactionFields...).Where(ex).Order(g.C("id").Desc()).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func TransactionInsert(data *types.Transaction) (int64, error) {

	query, _, _ := meta.Dialect.Insert("transactions").Rows(data).ToSQL()
	fmt.Println(query)
	res, err := meta.SqlxDb.Exec(query)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func TransListByTransFatherId(ctx context.Context, param types.GetTransListByTransFatherIdParams, page, size int) ([]*types.Transaction, int64, error) {
	var res []*types.Transaction
	var count int64

	boolQuery := elastic.NewBoolQuery()
	var terms []elastic.Query
	if param.TransFatherID != 0 {
		terms = append(terms, elastic.NewTermsQuery("trans_father_id", param.TransFatherID))
	}

	boolQuery.Must(terms...)
	boolQuery.Filter(elastic.NewRangeQuery("created_at").Gte(param.StartAt).Lte(param.EndAt))

	searchResult, err := meta.EsClient.Search().
		Index("transactions").
		From(page).
		Size(size).
		Sort("created_at", false).
		TrackTotalHits(true).
		Query(boolQuery).
		Do(ctx)
	if err != nil {
		return res, count, err
	}
	count = searchResult.Hits.TotalHits.Value
	for _, hit := range searchResult.Hits.Hits {
		var p types.Transaction
		err = json.Unmarshal(hit.Source, &p)
		if err != nil {
			return res, count, err
		}
		res = append(res, &p)
	}
	return res, count, nil
}

func TransAggByUserAndFatherId(ctx context.Context, userIds []int64, transFatherId int64, startTime, endTime string) (*types.TransAggResult, error) {
	boolQuery := elastic.NewBoolQuery()

	// 构建 must 查询条件
	boolQuery.Must(
		elastic.NewTermsQuery("user_id", userIds),
		elastic.NewTermQuery("trans_father_id", transFatherId),
		elastic.NewRangeQuery("created_at").Gte(startTime).Lte(endTime),
	)

	// 构建聚合
	aggUniqueUserCount := elastic.NewCardinalityAggregation().Field("user_id")
	aggSumAmount := elastic.NewSumAggregation().Field("amount")
	aggUserIds := elastic.NewTermsAggregation().Field("user_id").Size(100000)

	// 执行查询
	searchResult, err := meta.EsClient.Search().
		Index("transactions").
		Query(boolQuery).
		Size(0). // 不返回文档，只聚合
		Aggregation("unique_user_count", aggUniqueUserCount).
		Aggregation("sum_amount", aggSumAmount).
		Aggregation("unique_user_ids", aggUserIds).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	// 解析聚合结果
	result := &types.TransAggResult{}

	if agg, found := searchResult.Aggregations.Cardinality("unique_user_count"); found {
		result.UniqueUserCount = int64(*agg.Value)
	}
	if agg, found := searchResult.Aggregations.Sum("sum_amount"); found {
		result.SumAmount = *agg.Value
	}
	if agg, found := searchResult.Aggregations.Terms("unique_user_ids"); found {
		for _, bucket := range agg.Buckets {
			result.UserIDs = append(result.UserIDs, int64(bucket.Key.(float64)))
		}
	}

	return result, nil
}
