package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandSeq(n int) string {
	var sb strings.Builder
	k := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandSeq(10)
}

func RandomMoney() int64 {
	return RandInt(1, 1000)
}

func RandomCurrencies() string {
	currencies := []string{"USD", "EUR", "IDR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
