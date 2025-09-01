package dbx

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/davy66666/rpc_common/utils/loggerx"

	sqlxx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
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
	Host         string
	Port         int
	Username     string
	Password     string
	Database     string
	MaxIdleConn  int    // 空闲连接池中的最大连接数
	MaxOpenConn  int    // 数据库的最大打开连接数
	AuthDatabase string `json:",default=admin"` // 认证数据库
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
	if len(skip) == 0 {
		// WHERE 条件检查
		err = db.Use(&CheckTenantPlugin{RequiredField: "site"})
		if err != nil {
			panic(err)
		}
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

func MustMongodb(ctx context.Context, c MongoConf) (*mongo.Database, error) {
	// 如果有用户名和密码，则构建带有认证信息的 URI
	var uri string
	if c.Username != "" && c.Password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?%s",
			url.QueryEscape(c.Username),
			url.QueryEscape(c.Password),
			c.Host,
			c.Port,
			c.Database,
			url.Values{
				"authSource":    []string{c.AuthDatabase}, // 认证数据库（默认admin）
				"maxPoolSize":   []string{strconv.Itoa(c.MaxOpenConn)},
				"wtimeoutMS":    []string{"5000"},
				"socketTimeout": []string{"30000"},
			}.Encode(),
		)
	} else {
		// 如果没有用户名和密码，则构建一个无认证的 URI
		uri = fmt.Sprintf("mongodb://%s:%d/%s?%s",
			c.Host,
			c.Port,
			c.Database,
			url.Values{
				"maxPoolSize":   []string{strconv.Itoa(c.MaxOpenConn)},
				"wtimeoutMS":    []string{"5000"},
				"socketTimeout": []string{"30000"},
			}.Encode(),
		)
	}

	// 配置客户端选项（包括连接池、超时等）
	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(5 * time.Second).
		SetMaxPoolSize(100). // 连接池上限
		SetMinPoolSize(10). // 连接池下限
		SetRetryReads(true). // 自动重试读操作
		SetHeartbeatInterval(10 * time.Second) // 心跳检测

	// 建立 MongoDB 连接
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB connection failed: %w (URI: %s)", err, maskURI(uri))
	}

	// 健康检查（带超时控制）
	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err = client.Ping(pingCtx, nil); err != nil {
		_ = client.Disconnect(ctx) // 立即释放资源
		return nil, fmt.Errorf("MongoDB ping failed: %w", err)
	}

	// 返回 MongoDB 数据库对象
	return client.Database(c.Database), nil
}

// 安全脱敏URI（用于日志）
func maskURI(uri string) string {
	u, _ := url.Parse(uri)
	if u == nil {
		return "invalid_uri"
	}
	if u.User != nil {
		u.User = url.UserPassword(u.User.Username(), "******")
	}
	return u.String()
}
