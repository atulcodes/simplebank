package util

import (
	"math/rand"
	"strings"
	"time"
)

var source rand.Source

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source = rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	rng := rand.New(source)
	return min + rng.Int63n(max - min + 1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < k; i++ {
		ch := alphabet[rand.Intn(k)]
		sb.WriteByte(ch)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{INR, USD, CAD, EUR}
	return currencies[rand.Intn(len(currencies))]
}
