package poolx

import (
	"github.com/panjf2000/ants/v2"
)

func MustPoolWithFunc(size int, fn func(interface{})) *ants.PoolWithFunc {
	p, err := ants.NewPoolWithFunc(size, fn)
	if err != nil {
		panic(err)
	}

	return p
}

func MustPool(size int) *ants.Pool {
	p, err := ants.NewPool(size)
	if err != nil {
		panic(err)
	}

	return p
}
