package main

import (
	"fmt"
	"sync"
)

type KeyValue struct {
	data map[string]int32
	mu sync.RWMutex
}

func NewKeyValue() *KeyValue {
	return &KeyValue{
		data: make(map[string]int32),
	}
}

func (kv *KeyValue) Set(key string, val int32){
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[key] = val
}

func (kv *KeyValue) Get(key string) (int32, bool) {
	kv.mu.RLock()
	defer  kv.mu.RUnlock()

	val, ok := kv.data[key]
	return val, ok
}

func main() {
	kv := NewKeyValue()
	kv.Set("key", 100)
	val, ok := kv.Get("key")
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("key is not register")
	}
	val, ok = kv.Get("key2")
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("key is not register")
	}
}
