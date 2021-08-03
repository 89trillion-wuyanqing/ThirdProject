package utils

import (
	"crypto/rand"
	"math/big"
)

type RandomCode struct {
}

var defaultLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 随机返回8位字符串
func (this *RandomCode) RandomString() string {

	b := make([]rune, 8)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(defaultLetters))))
		b[i] = defaultLetters[n.Int64()]
	}

	return string(b)
}
