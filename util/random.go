package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}

	return sb.String()
}

func GenerateOwner() string {
	return RandomString(5)
}

func GenerateBalance() int64 {
	return RandomInt(0, 1000)
}

func GenerateCurrency() string {
	currs := []string{
		EUR, USD, CAD,
	}
	return currs[rand.Intn(len(currs))]
}