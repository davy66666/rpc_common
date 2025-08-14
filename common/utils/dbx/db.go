package dbx

import (
	"fmt"
	"log"
	"time"

	"github.com/davy66666/rpc_service/common/utils/loggerx"

	sqlxx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type MysqlConf struct {
	DataSource string
	// 连接池参数
	MaxOpenConns    int   `json:",default=100"`
	MaxIdleConns    int   `json:",default=50"`
	ConnMaxLifetime int   `json:",default=3600"`
	SlowThreshold   int64 `json:",default=500"` // 慢查询阈值（毫秒）
}

type PgsqlConf struct {
	DataSource string
}
type MongoConf struct {
	Uri      string
	Database string
}

type DBConf struct {
	Host         string
	Port         int
	Resolver     []string `json:"Resolver,optional"`
	Username     string
	Password     string
	Database     string
	MaxIdleConns int // 空闲连接池中的最大连接数
	MaxOpenConns int // 数据库的最大打开连接数
}

func MustDB(c DBConf, skip ...any) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", c.Username, c.Password, c.Host, c.Port, c.Database)
	// 打开 MySQL 数据库连接。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                           loggerx.NewGormLogger(),
		IgnoreRelationshipsWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	if len(c.Resolver) > 0 {
		var replicas []gorm.Dialector
		for _, v := range c.Resolver {
			resolverDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", c.Username, c.Password, v, c.Port, c.Database)
			replicas = append(replicas, mysql.Open(resolverDsn))
		}
		// 配置读写分离
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsn)}, // 主库（写库）
			Replicas: replicas,                          // 从库（读库），可以配置多个
			Policy:   dbresolver.RandomPolicy{},         // 负载均衡策略
		}))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(c.MaxIdleConns) // 空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(c.MaxOpenConns) // 数据库的最大打开连接数
	// 注册插件
	err = db.Use(&ReplaceSelectStatementPlugin{})
	if err != nil {
		panic(err)
	}

	return db
}

func NewSilentMysql(c MysqlConf) sqlx.SqlConn {
	// 创建自定义驱动
	conn := sqlx.NewMysql(c.DataSource)

	// 获取底层驱动并设置日志
	rawConn, err := conn.RawDB()
	if err != nil {
		panic("初始化mysql失败")
	}
	if c.ConnMaxLifetime > 0 {
		rawConn.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))
	}
	if c.MaxIdleConns > 0 {
		rawConn.SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.MaxIdleConns > 0 {
		rawConn.SetMaxOpenConns(c.MaxOpenConns)
	}
	// 关键：设置空日志

	return conn
}

func MustSqlxDB(c DBConf) *sqlxx.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", c.Username, c.Password, c.Host, c.Port, c.Database)
	db, err := sqlxx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * 30)
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func MustSqlxPostgres(c DBConf) *sqlxx.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.Password, c.Database)
	db, err := sqlxx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * 30)
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
