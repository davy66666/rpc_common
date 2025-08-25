package poolx

import (
	"context"
	"fmt"
	"testing"
)

func TestMustPool(t *testing.T) {
	p := MustPool(10)
	err := p.Submit(func() {
		fmt.Println("hello")
	})
	if err != nil {
		t.Error(err)
	}
	p.Release()
}
func TestMustPoolWithFunc(t *testing.T) {
	p := MustPoolWithFunc(10, Print)
	for i := 0; i < 10; i++ {
		err := p.Invoke(i)
		if err != nil {
			t.Error(err)
		}
	}

	p.Release()
}

func Print(in interface{}) {
	fmt.Println(in)
}

func TestMustPoolWithFunc2(t *testing.T) {
	s := NewService(context.Background(), "hello")

	p := MustPoolWithFunc(10, s.Run)
	for i := 0; i < 10; i++ {
		err := p.Invoke("world")
		if err != nil {
			t.Error(err)
		}
	}

	p.Release()
}

type Service struct {
	ctx   context.Context
	Param any
}

func NewService(ctx context.Context, param any) (s *Service) {
	return &Service{ctx, param}
}

func (s *Service) Run(in any) {
	fmt.Printf("%v, %v\n", s.Param, in)
}
