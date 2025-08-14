package svc

import (
	g "github.com/doug-martin/goqu/v9"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/davy66666/rpc_service/common/utils/dbx"
	"github.com/davy66666/rpc_service/common/utils/rabbitmqc"
	"github.com/davy66666/rpc_service/internal/config"
	"github.com/davy66666/rpc_service/internal/model"
)

type ServiceContext struct {
	Config   config.Config
	Rabbitmq *amqp.Connection
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.DisableStat()
	mt := new(model.MetaTable)
	mt.SqlxDb = dbx.MustSqlxDB(c.Mysql.MasterDB)
	mt.LogDb = dbx.MustSqlxDB(c.Mysql.LogDB)
	mt.ActivityDb = dbx.MustSqlxDB(c.Mysql.ActivityDB)
	mt.Dialect = g.Dialect("mysql")
	mt.PGDialect = g.Dialect("postgres")
	mt.Rds = dbx.MustRedis(c.RedisConf)
	conn := rabbitmqc.MustProducer(c.RabbitMqConf)
	//mt.Dbm1Db = dbx.MustSqlxPostgres(c.Mysql.Dbm1DB)
	//mt.EsClient = elasticx.MustElastic(c.ElasticConf)
	model.Constructor(mt)
	return &ServiceContext{
		Config:   c,
		Rabbitmq: conn,
	}
}

func (c *ServiceContext) Close() {

}
