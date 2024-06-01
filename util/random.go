package util

import (
	"math/rand"
	"strings"
)

const (
	alpha = "abcdefghijklmnopqrstuvwxyz"
)

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alpha)
	for i := 0; i < n; i++ {
		c := alpha[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "INR", "YEN", "ERU", "AUD"}
	return currencies[rand.Intn(len(currencies))]
}
