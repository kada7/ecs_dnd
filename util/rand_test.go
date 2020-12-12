package util

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandRange(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		_ = RandRange(r, -100, 100)
		//t.Log(n)
	}
}
