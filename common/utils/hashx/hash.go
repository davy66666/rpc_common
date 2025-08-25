package hashx

import (
	"hash/crc32"
)

// GetShardByNum 计算哈希取模数字类型
func GetShardByNum(field, numShards int64) int64 {
	shard := field % numShards
	return shard // 返回分片名
}

// GetShardByStr 计算哈希取模-String
func GetShardByStr(field string, numShards int64) int64 {
	hash := crc32.ChecksumIEEE([]byte(field))
	shard := hash % uint32(numShards)
	return int64(shard) // 返回分片名
}
