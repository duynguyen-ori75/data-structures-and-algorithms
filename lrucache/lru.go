package lrucache

// Tested in LeetCode

import (
	"container/list"
)

type Pair struct {
	key, value int
}

type LRUCache struct {
	capacity int
	data     *list.List
	elemPos  map[int]*list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{capacity: capacity, data: list.New(), elemPos: make(map[int]*list.Element)}
}

func (cache *LRUCache) evict() {
	pair := cache.data.Remove(cache.data.Front())
	delete(cache.elemPos, pair.(Pair).key)
}

func (cache *LRUCache) Get(key int) int {
	element, ok := cache.elemPos[key]
	if !ok {
		return -1
	}
	cache.Put(element.Value.(Pair).key, element.Value.(Pair).value)
	return element.Value.(Pair).value
}

func (cache *LRUCache) Put(key int, value int) {
	element, ok := cache.elemPos[key]
	if !ok && cache.data.Len() == cache.capacity {
		cache.evict()
	}
	if ok {
		cache.data.Remove(element)
		delete(cache.elemPos, key)
	}
	cache.elemPos[key] = cache.data.PushBack(Pair{key: key, value: value})
}
