package question

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		cache := LRUConstructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		if got := cache.Get(1); got != 1 {
			t.Errorf("LRUCache() = %v, want %v", got, 1)
		}
		cache.Put(3, 3)
		if got := cache.Get(2); got != -1 {
			t.Errorf("LRUCache() = %v, want %v", got, -1)
		}
		cache.Put(4, 4)
		if got := cache.Get(1); got != -1 {
			t.Errorf("LRUCache() = %v, want %v", got, -1)
		}
		if got := cache.Get(3); got != 3 {
			t.Errorf("LRUCache() = %v, want %v", got, 4)
		}
		if got := cache.Get(4); got != 4 {
			t.Errorf("LRUCache() = %v, want %v", got, 4)
		}
	})
}
