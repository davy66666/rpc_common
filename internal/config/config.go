package config

import (
	"ks_api_service/common/utils/dbx"
	"ks_api_service/common/utils/elasticx"
	"ks_api_service/common/utils/rabbitmqc"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/zrpc"
)

type Mysql struct {
	MasterDB   dbx.DBConf // 使用通用的 DBConf 类型
	SlaveDB    dbx.DBConf // 使用通用的 DBConf 类型
	Dbm1DB     dbx.DBConf // 使用通用的 DBConf 类型
	LogDB      dbx.DBConf // 使用通用的 DBConf 类型
	ActivityDB dbx.DBConf // 使用通用的 DBConf 类型
}
type Config struct {
	zrpc.RpcServerConf
	ElasticConf  elasticx.ElasticConf
	RedisConf    dbx.RedisConf
	Mysql        Mysql
	RabbitMqConf rabbitmqc.RabbitMqConf
}

func (c *Config) Parse(path, key string) {

	// 读取本地配置，常用于本地开发
	if strings.Contains(path, "yaml") {
		conf.MustLoad(path, c)
	} else {
		c.parseRemote(path, key)
	}
}

func (c *Config) parseRemote(path, key string) {
	// 创建 etcd subscriber
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{path}, // etcd 地址
		Key:   key,            // 配置key
	})

	// 创建 configurator
	cc := configurator.MustNewConfigCenter[*Config](configurator.Config{
		Type: "yaml", // 配置值类型：json,yaml,toml
	}, ss)

	// 获取配置
	// 注意: 配置如果发生变更，调用的结果永远获取到最新的配置
	v, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}

	// 更新配置对象
	*c = *v // 更新传入的指针对象

	// 启动一个 goroutine 来异步监听配置变更
	threading.GoSafe(func() {
		// 添加监听器
		cc.AddListener(func() {
			// 获取最新的配置
			v, err = cc.GetConfig()
			if err != nil {
				panic(err) // 根据需要进行错误处理
			}
			*c = *v // 更新传入的指针对象
		})

		// 使用 select {} 语句保持 goroutine 不退出，持续监听配置变化
		select {} // 阻塞该 goroutine，保持监听器持续运行
	})
}
