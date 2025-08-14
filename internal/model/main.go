package model

import (
	g "github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
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
}

func Close() {

	_ = meta.SqlxDb.Close()
	_ = meta.Dbm1Db.Close()
}
