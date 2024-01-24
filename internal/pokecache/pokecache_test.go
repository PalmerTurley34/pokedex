package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("example data"),
		},
		{
			key: "https://example.com/test",
			val: []byte("testing data"),
		},
		{
			key: "httpls://example.com/test/path",
			val: []byte("more test data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("values do not match: %v vs %v", val, c.val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const reapInterval = 50 * time.Millisecond
	const waitTime = reapInterval * 2
	const testKey = "https://example.com"
	cache := NewCache(reapInterval)
	cache.Add(testKey, []byte("test data"))
	_, ok := cache.Get(testKey)
	if !ok {
		t.Errorf("expected to find key")
	}
	time.Sleep(waitTime)
	_, ok = cache.Get(testKey)
	if ok {
		t.Errorf("expected key to be reaped")
	}

}
