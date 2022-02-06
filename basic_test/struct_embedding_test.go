package basic

import (
	"sync"
	"testing"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func TestStructEm(t *testing.T) {
	go t.Log(Lookup("a"))
	go t.Log(Lookup("b"))
}