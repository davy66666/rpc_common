package idx

import (
	"github.com/gofrs/uuid"
)

// GenUUID 生成UUID
func GenUUID() string {
	v7, _ := uuid.NewV7()
	return v7.String()
}
