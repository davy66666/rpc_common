package idx

import (
	"testing"
)

func TestGetNanoId(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenNanoId())
	}
}

func TestGenShortId(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenShortId())
	}
}

func TestGenUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenUUID())
	}
}

func TestGenTid(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenTid())
	}
}
