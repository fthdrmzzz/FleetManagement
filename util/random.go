package util

import (
	"math/rand"
	"strings"
	"time"
)

// Random utilities for db testing
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomBarcode() string {
	return RandomString(10)
}

func RandomWeight() int32 {
	return RandomInt(0, 100)
}
