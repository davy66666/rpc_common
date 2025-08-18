package model

import (
	"github.com/davy66666/rpc_service/internal/types"
	g "github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"

	"reflect"
	"strings"
)

var (
	UserFields                   []interface{}
	TransactionTypeFields        []interface{}
	TransactionFields            []interface{}
	BetAmountFields              []interface{}
	UserGameBetAmountFields      []interface{}
	ClearBetAmountLogFields      []interface{}
	PayLevelFields               []interface{}
	ActivityFields               []interface{}
	FissionSettingFields         []interface{}
	FissionExclusiveRewardFields []interface{}
	TeamUserFields               []interface{}
	ActTransactionFields         []interface{}
	ActivityReportFields         []interface{}
	BetAmountLogFields           []interface{}
	FissionRewardFields          []interface{}
	GiftMoneyTransactionFields   []interface{}
	LevelUpgradeLogFields        []interface{}
)

type MetaTable struct {
	SqlxDb     *sqlx.DB
	Dbm1Db     *sqlx.DB
	LogDb      *sqlx.DB
	ActivityDb *sqlx.DB
	Dialect    g.DialectWrapper
	PGDialect  g.DialectWrapper
	EsClient   *elastic.Client
	Rds        redis.Cmdable
}

var (
	meta *MetaTable
)

func Constructor(mt *MetaTable) {
	meta = mt
	InitTableField()
}

func InitTableField() {
	UserFields = GetSelectFieldList(&types.User{})
	TransactionTypeFields = GetSelectFieldList(&types.TransactionType{})
	TransactionFields = GetSelectFieldList(&types.Transaction{})
	BetAmountFields = GetSelectFieldList(&types.BetAmount{})
	UserGameBetAmountFields = GetSelectFieldList(&types.UserGameBetAmount{})
	ClearBetAmountLogFields = GetSelectFieldList(&types.ClearBetAmountLog{})
	PayLevelFields = GetSelectFieldList(&types.PayLevel{})
	ActivityFields = GetSelectFieldList(&types.Activity{})
	FissionSettingFields = GetSelectFieldList(&types.FissionSetting{})
	FissionExclusiveRewardFields = GetSelectFieldList(&types.FissionExclusiveReward{})
	TeamUserFields = GetSelectFieldList(&types.TeamUser{})
	ActTransactionFields = GetSelectFieldList(&types.ActTransaction{})
	ActivityReportFields = GetSelectFieldList(&types.ActivityReport{})
	BetAmountLogFields = GetSelectFieldList(&types.BetAmountLog{})
	FissionRewardFields = GetSelectFieldList(&types.FissionReward{})
	GiftMoneyTransactionFields = GetSelectFieldList(&types.GiftMoneyTransaction{})
	LevelUpgradeLogFields = GetSelectFieldList(&types.LevelUpgradeLog{})
}
func Close() {

	_ = meta.SqlxDb.Close()
	_ = meta.Dbm1Db.Close()
}

// 自动提取结构体字段名（优先使用 db 标签）
func GetSelectFieldList(obj interface{}) []interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fields []interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			fields = append(fields, dbTag)
		} else {
			fields = append(fields, strings.ToLower(field.Name))
		}
	}
	return fields
}
