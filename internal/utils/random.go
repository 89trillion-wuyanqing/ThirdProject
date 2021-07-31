package utils

import (
	"crypto/rand"
	"math/big"
)

type RandomCode struct {
}

var defaultLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString returns a random string with a fixed length
func (this *RandomCode) RandomString() string {

	b := make([]rune, 8)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(defaultLetters))))
		b[i] = defaultLetters[n.Int64()]
	}

	return string(b)
}
