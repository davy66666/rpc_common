package idx

import (
	"github.com/teris-io/shortid"
)

func GenShortId() string {
	id, _ := shortid.Generate()
	return id
}
