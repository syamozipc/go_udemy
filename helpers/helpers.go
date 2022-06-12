package helpers

import (
	"math/rand"
	"time"
)

// lesson16で使用
// custom package
type SomeType struct {
	TypeName string
	TypeNumber int
}

// lesson17で使用
// channelについて
func RandomNumber(n int) int{
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}