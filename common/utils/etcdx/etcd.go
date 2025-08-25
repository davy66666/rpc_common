package etcdx

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConf struct {
	Endpoints []string
	Username  string
	Password  string
}

func MustEtcdClient(c EtcdConf) *clientv3.Client {
	// 配置 etcd 客户端
	conf := clientv3.Config{
		Endpoints:   c.Endpoints, // etcd 节点地址列表
		DialTimeout: 5 * time.Second,
	}

	// 如果提供了用户名和密码，则启用认证
	if c.Username != "" && c.Password != "" {
		conf.Username = c.Username
		conf.Password = c.Password
	}

	// 创建 etcd 客户端
	client, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}

	return client
}
