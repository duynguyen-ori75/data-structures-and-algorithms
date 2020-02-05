package bloom

import (
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
)

type BloomFilter struct {
	content 	[]bool
	length 		int
	hashFuncs  	[]hash.Hash32
}

func NewBloomFilter(length int) *BloomFilter {
	result := &BloomFilter{length: length, content: make([]bool, length)}
	result.hashFuncs = []hash.Hash32{adler32.New(), crc32.NewIEEE(), fnv.New32()}
	return result
}

func (bloom *BloomFilter) hashSum(funcIndex int, s string) int {
    bloom.hashFuncs[funcIndex].Write([]byte(s))
    defer bloom.hashFuncs[funcIndex].Reset()
    return int(bloom.hashFuncs[funcIndex].Sum32()) % bloom.length
}

func (bloom *BloomFilter) Add(s string) {
	for hashIndex := range bloom.hashFuncs {
		index := bloom.hashSum(hashIndex, s)
		bloom.content[index] = true
	}
}

func (bloom *BloomFilter) PossiblyHave(s string) bool {
	for hashIndex := range bloom.hashFuncs {
		index := bloom.hashSum(hashIndex, s)
		if !bloom.content[index] {
			return false
		}
	}
	return true
}

func (bloom *BloomFilter) Reset() {
	for index := 0; index < bloom.length; index ++ {
		bloom.content[index] = false
	}
}