package utils

import (
	"math/rand"
	"time"
)

func Random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
