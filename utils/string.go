package utils

import (
	"crypto/rand"
	"math/big"
)

func RandString(len uint) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	ret := make([]rune, len)
	var randInt *big.Int
	for i := range ret {
		randInt, _ = rand.Int(rand.Reader, big.NewInt(26+26+10))
		ret[i] = letters[randInt.Int64()]
	}
	return string(ret)
}
