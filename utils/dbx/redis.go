package dbx

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/davy66666/rpc_common/value"
	"strings"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConf struct {
	Host         string
	Port         int
	Password     string
	ClientName   string
	TLS          bool   `json:"TLS,optional"`
	PoolSize     int64  // 单进程最大连接数（降低至4）
	MinIdleConns int64  // 减少空闲连接占用
	MaxIdleConns int64  // 避免堆积过多空闲连接
	Type         string // node 或者 cluster
}

type RedisDeleteConf struct {
	Prefix    string
	ScanSize  int64
	BatchSize int
	MaxQPS    int
	Progress  func(deleted int, totalEstimate int64)
	UseUnlink bool // 使用UNLINK代替DEL
}

func MustRedis(c RedisConf) redis.Cmdable {
	switch c.Type {
	case "node":
		return mustClientRedis(c)
	default:
		return mustClusterRedis(c)
	}
}

func mustRingRedis(c RedisConf) redis.Cmdable {
	var addr = map[string]string{}
	opt := &redis.RingOptions{
		Addrs:        addr,
		PoolSize:     4,  // 单进程最大连接数（降低至4）
		MinIdleConns: 1,  // 减少空闲连接占用
		MaxIdleConns: 50, // 避免堆积过多空闲连接
	}
	if c.Password != "" {
		opt.Password = c.Password
	}
	if c.TLS {
		opt.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证（仅测试环境用）
		}
	}
	if c.PoolSize != 0 {
		opt.PoolSize = int(c.PoolSize) // 最大连接数
	}
	if c.MinIdleConns != 0 {
		opt.MinIdleConns = int(c.MinIdleConns) // 最小空闲连接
	}
	if c.MaxIdleConns != 0 {
		opt.MaxIdleConns = int(c.MaxIdleConns) // 最大空闲连接
	}
	ring := redis.NewRing(opt)
	err := ring.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return ring
}
func mustClusterRedis(c RedisConf) redis.Cmdable {
	opt := &redis.ClusterOptions{
		Addrs: []string{fmt.Sprintf("%s:%d", c.Host, c.Port)},
	}
	if c.Password != "" {
		opt.Password = c.Password
	}
	if c.ClientName != "" {
		opt.ClientName = c.ClientName
	}
	if c.TLS {
		opt.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证（仅测试环境用）
		}
	}
	if c.PoolSize != 0 {
		opt.PoolSize = int(c.PoolSize) // 最大连接数
	}
	if c.MinIdleConns != 0 {
		opt.MinIdleConns = int(c.MinIdleConns) // 最小空闲连接
	}
	if c.MaxIdleConns != 0 {
		opt.MaxIdleConns = int(c.MaxIdleConns) // 最大空闲连接
	}
	client := redis.NewClusterClient(opt)
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	client.ReloadState(context.Background())
	return client
}
func mustClientRedis(c RedisConf) redis.Cmdable {
	opt := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
	}
	if c.Password != "" {
		opt.Password = c.Password
	}
	if c.TLS {
		opt.TLSConfig = &tls.Config{}
	}
	client := redis.NewClient(opt)
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return client
}

func GenKey(prefix, key string, a ...any) string {

	var nKey string
	prefix = prefix + ":"
	switch key {
	case
		values.RedisKeyNodeKeyPrefix,     // 节点 ID 的 Redis 键前缀
		values.RedisKeyLockUniqueRequest, // 请求幂等锁
		values.RedisKeyLockUniqueRoundID, // 局数唯一ID
		values.RedisKeyUniqueUserId,      // 用户唯一ID
		values.RedisKeyUniqueMsgId,       // 消息唯一ID
		values.RedisKeyUniqueSysId,       // 系统唯一ID
		values.RedisKeyUniqueAd:          // 广告唯一事件
		nKey = key
	default:
		if strings.HasPrefix(key, prefix) {
			nKey = key
		} else {
			nKey = prefix + key
		}
	}

	if len(a) == 0 {
		return nKey
	}

	return fmt.Sprintf(nKey, a...)
}

func DeleteRedisKeysByPrefix(ctx context.Context, rdb redis.Cmdable, conf RedisDeleteConf) error {
	start := time.Now()
	var totalDeleted int64

	// 判断是否为 ClusterClient
	if cluster, ok := rdb.(*redis.ClusterClient); ok {
		return DeleteRedisKeysByPrefixClusterSafe(ctx, cluster, conf)
	}

	// 单节点处理
	var cursor uint64
	for {
		keys, nextCursor, err := rdb.Scan(ctx, cursor, conf.Prefix+"*", conf.ScanSize).Result()
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		if err = deleteInBatches(ctx, rdb, keys, &totalDeleted, conf); err != nil {
			return err
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	fmt.Printf("✅ Deleted %d keys in %v\n", totalDeleted, time.Since(start))
	return nil
}

// DeleteRedisKeysByPrefixClusterSafe 集群版本，遍历所有 master 节点安全删除
func DeleteRedisKeysByPrefixClusterSafe(ctx context.Context, cluster *redis.ClusterClient, conf RedisDeleteConf) error {
	var totalDeleted int64
	start := time.Now()

	err := cluster.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
		var cursor uint64
		for {
			keys, nextCursor, err := client.Scan(ctx, cursor, conf.Prefix+"*", conf.ScanSize).Result()
			if err != nil {
				return fmt.Errorf("cluster scan failed: %w", err)
			}

			if err := deleteInBatchesCluster(ctx, client, keys, &totalDeleted, conf); err != nil {
				return err
			}

			if nextCursor == 0 {
				break
			}
			cursor = nextCursor
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("✅ [Cluster] Deleted %d keys in %v\n", totalDeleted, time.Since(start))
	return nil
}

// 通用删除批处理逻辑
func deleteInBatches(ctx context.Context, rdb redis.Cmdable, keys []string, totalDeleted *int64, conf RedisDeleteConf) error {
	for i := 0; i < len(keys); i += conf.BatchSize {
		end := i + conf.BatchSize
		if end > len(keys) {
			end = len(keys)
		}
		batch := keys[i:end]

		if len(batch) == 0 {
			continue
		}

		var cmd *redis.IntCmd
		if conf.UseUnlink {
			cmd = rdb.Unlink(ctx, batch...)
		} else {
			cmd = rdb.Del(ctx, batch...)
		}

		if _, err := cmd.Result(); err != nil {
			return fmt.Errorf("delete failed: %w", err)
		}

		atomic.AddInt64(totalDeleted, int64(len(batch)))
		if conf.Progress != nil {
			conf.Progress(int(atomic.LoadInt64(totalDeleted)), -1)
		}

		if conf.MaxQPS > 0 {
			time.Sleep(time.Duration(float64(time.Second) * float64(len(batch)) / float64(conf.MaxQPS)))
		}
	}
	return nil
}

// 集群模式删除批处理逻辑
func deleteInBatchesCluster(ctx context.Context, rdb redis.Cmdable, keys []string, totalDeleted *int64, conf RedisDeleteConf) error {
	for i := 0; i < len(keys); i += conf.BatchSize {
		end := i + conf.BatchSize
		if end > len(keys) {
			end = len(keys)
		}
		batch := keys[i:end]

		if len(batch) == 0 {
			continue
		}

		cmder, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			for _, key := range batch {
				if conf.UseUnlink {
					pipe.Unlink(ctx, key)
				} else {
					pipe.Del(ctx, key)
				}
			}
			return nil
		})
		if err != nil {
			return err
		}

		for _, cmd := range cmder {
			if err = cmd.Err(); err != nil {
				return fmt.Errorf("delete failed: %w", err)
			}
		}

		atomic.AddInt64(totalDeleted, int64(len(batch)))
		if conf.Progress != nil {
			conf.Progress(int(atomic.LoadInt64(totalDeleted)), -1)
		}

		if conf.MaxQPS > 0 {
			time.Sleep(time.Duration(float64(time.Second) * float64(len(batch)) / float64(conf.MaxQPS)))
		}
	}
	return nil
}
