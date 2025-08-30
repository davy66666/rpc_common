package idx

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func HashAndEncodeUnique(num int64) string {
	data := fmt.Sprintf("%d-%d-%s", num, time.Now().UnixNano(), uuid.New().String()) // 加 UUID
	hash := md5.Sum([]byte(data))
	return base64.RawURLEncoding.EncodeToString(hash[:8]) // 取 8 字节
}
