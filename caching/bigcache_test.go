package caching

import (
	"fmt"
	"testing"
	"time"

	"github.com/allegro/bigcache"
)

func TestBigCache(t *testing.T) {
	cfg := bigcache.Config{
		Shards:             1024,
		LifeWindow:         20 * time.Second,
		CleanWindow:        5 * time.Second,
		OnRemoveWithReason: func(key string, entry []byte, reason bigcache.RemoveReason) {
			fmt.Printf("Removing %s %v\n", key, reason)
		},
	}

	bc, _ := bigcache.NewBigCache(cfg)

	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("key %d", i)
		v := fmt.Sprintf("value %d", i)
		bc.Set(k, []byte(v))
		time.Sleep(100 * time.Millisecond)
	}
}
