package model

import (
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"ks_api_service/internal/types"
	"strings"
)

func GetMessageJob(id int64) (types.MessageJob, error) {
	var mj types.MessageJob
	//query, _, _ := l.svcCtx.Dialect.From("message_job").Where(g.Ex{"id": id}).Limit(1).ToSQL()
	query := fmt.Sprintf(`SELECT id,msg_id,msg_name,data,status  FROM message_job WHERE id = %d`, id)
	fmt.Println(query)
	err := meta.SqlxDb.Get(&mj, query)

	return mj, err
}

func InsertMessage(ms []*types.PGMessage) error {

	query, _, _ := meta.PGDialect.Insert("messages").Rows(ms).ToSQL()
	fmt.Println(query)
	_, err := meta.Dbm1Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func InsertReplyContent(rs []types.ReplyContent) error {

	query, _, _ := meta.PGDialect.Insert("reply_content").Rows(rs).ToSQL()
	fmt.Println(query)
	_, err := meta.Dbm1Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMessageJob(id, status int64) error {

	query, _, _ := meta.Dialect.Update("message_job").Set(g.Record{"status": status}).Where(g.Ex{"id": id}).ToSQL()
	fmt.Println(query)
	_, err := meta.Dbm1Db.Exec(query)
	if err != nil {
		return fmt.Errorf("更新失败: %w", err)
	}

	return nil
}

func InsertMessagePushBatch(mpb types.MessagePushBatch) (int64, error) {

	var id int64
	query, _, _ := meta.PGDialect.Insert("message_push_batches").Rows(mpb).ToSQL()
	query = query + " RETURNING id"
	fmt.Println(query)
	err := meta.Dbm1Db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateMessagePushBatch(id, successCount, failCount int64) error {

	query, _, _ := meta.PGDialect.Update("message_push_batches").Set(g.Record{
		"success_count": successCount,
		"fail_count":    failCount,
	}).Where(g.Ex{"id": id}).ToSQL()
	fmt.Println(query)
	_, err := meta.Dbm1Db.Exec(query)
	if err != nil {
		return fmt.Errorf("更新失败: %w", err)
	}

	return nil
}

func InsertMessagePushRecord(inserts []types.MessagePushRecord) error {

	query, _, _ := meta.Dialect.Insert("message_push_records").Rows(inserts).ToSQL()
	query = strings.Replace(query, "INSERT", "INSERT IGNORE", 1)
	fmt.Println(query)
	_, err := meta.Dbm1Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func FetchMessageType() ([]*types.MessageType, error) {

	var pl []*types.MessageType
	query, _, _ := meta.Dialect.Select("id", "en_name").From("message_types").ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&pl, query)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	return pl, nil
}

func FetchMessage(ex g.Ex, page, size int) ([]types.Message, error) {

	var ms []types.Message
	query, _, _ := meta.PGDialect.Select("id", "user_id", "username", "msg_id").Order(g.C("id").Asc()).From("messages").Limit(uint(size)).Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.Dbm1Db.Select(&ms, query)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	return ms, nil
}
