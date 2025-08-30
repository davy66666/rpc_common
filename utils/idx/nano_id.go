package idx

import "github.com/matoous/go-nanoid/v2"

// GenNanoId 获取NanoId
// 该ID替代UUID,占用的字节数更小,生成速度更快
func GenNanoId() string {
	id, _ := gonanoid.New()
	return id
}
