package hashing

import (
	"hash/crc32"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const nameLength = 10

var hashFunc = crc32.NewIEEE()

func randomString() string {
	b := make([]byte, nameLength)
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

func insertString(slice []string, index int, newElement string) []string {
	return append(slice[:index], append([]string{newElement}, slice[index:]...)...)
}

func insertBool(slice []bool, index int, newElement bool) []bool {
	return append(slice[:index], append([]bool{newElement}, slice[index:]...)...)
}

func insertNode(slice []*Node, index int, newElement *Node) []*Node {
	return append(slice[:index], append([]*Node{newElement}, slice[index:]...)...)
}
