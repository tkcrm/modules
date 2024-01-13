package utils

import (
	"crypto/rand"
	"math"
	mrand "math/rand"
	"unsafe"
)

var numbers = []byte("0123456789")
var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var symbols = []byte("[]/-()%#@")

// GenerateRandomString
func GenerateRandomString(length int, includeNumbers bool, includeSymbols bool) string {
	str := alphabet
	if includeNumbers {
		str = append(str, numbers...)
	}
	if includeSymbols {
		str = append(str, symbols...)
	}

	b := make([]byte, length)
	rand.Read(b)
	for i := 0; i < length; i++ {
		b[i] = str[b[i]%byte(len(str))]
	}

	return *(*string)(unsafe.Pointer(&b))
}

// GenerateRandomNumber - returns random number with specified number length
func GenerateRandomNumber(length uint) int {
	min := 1
	max := 9

	if length > 1 {
		min = int(math.Pow(10, float64(length-1)))
		max = min*10 - 1
	}

	return mrand.Intn(max-min) + min
}
