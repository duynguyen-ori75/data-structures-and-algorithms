package hashing

import (
	"hash/crc32"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const nameLength = 10

var hashFunc = crc32.NewIEEE()

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func hashSum(s string) int {
	hashFunc.Write([]byte(s))
	defer hashFunc.Reset()
	return int(hashFunc.Sum32())
}
