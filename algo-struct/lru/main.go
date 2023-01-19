package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element // map<string>*listElem
	lastUsed *list.List               // list<*listElem>
}

type listElem struct {
	key   string
	value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    map[string]*list.Element{},
		lastUsed: list.New(),
	}
}

func (l *LRUCache) Get(key string) interface{} {
	elem, ok := l.cache[key]
	if !ok {
		return nil
	}

	l.lastUsed.MoveToFront(elem)
	return elem.Value.(*listElem).value
}

func (l *LRUCache) Put(key string, value interface{}) {
	elem, ok := l.cache[key]
	if !ok {
		newElem := l.lastUsed.PushFront(&listElem{
			key:   key,
			value: value,
		})
		l.cache[key] = newElem

		if l.lastUsed.Len() > l.capacity {
			l.removeOutOfCapacityElem()
		}
	} else {
		elem.Value.(*listElem).value = value
		l.lastUsed.MoveToFront(elem)
	}
}

func (l *LRUCache) removeOutOfCapacityElem() {
	fmt.Printf("** Removed elem\n")

	toRemove := l.lastUsed.Back()
	delete(l.cache, toRemove.Value.(*listElem).key)
	l.lastUsed.Remove(toRemove)
}

func main() {
	lru := NewLRUCache(3)
	lru.Put("name", "Lukasz")
	lru.Put("age", 46)
	lru.Put("name", "Konrad")
	lru.Put("mail", "test@test.test")
	lru.Get("age")
	lru.Put("salary", 10000)

	v := lru.Get("name")
	fmt.Printf("v=%#v\n", v)
}
