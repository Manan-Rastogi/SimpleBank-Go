package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}


// Generates a random string of n length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return strings.Title(RandomString(7) + " " + RandomString(5))
}

// Generates random money amount
func RandomMoney() int64 {
	return RandomInt(100, 100000)
}

// Generates random currency value
func RandomCurrency() string {
	currencies := []string{
		"INR",
		"USD",
		"CAD",
		"EUR",
	}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}