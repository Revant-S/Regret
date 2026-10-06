package expiringmap

import (
	"sync"
	"time"
)

type entry[V any] struct {
	value    V
	deadline time.Time
}

type shard[K comparable, V any] struct {
	mu    sync.Mutex
	items map[K]entry[V]
}

func newShard[K comparable, V any]() shard[K, V] {
	return shard[K, V]{items: make(map[K]entry[V])}
}

func (s *shard[K, V]) set(key K, value V, deadline time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = entry[V]{value: value, deadline: deadline}
}

func (s *shard[K, V]) take(key K) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	delete(s.items, key)
	return e.value, true
}

func (s *shard[K, V]) len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}

func (s *shard[K, V]) removeExpired(now time.Time) map[K]V {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out map[K]V
	for k, e := range s.items {
		if now.After(e.deadline) {
			if out == nil {
				out = make(map[K]V)
			}
			out[k] = e.value
			delete(s.items, k)
		}
	}
	return out
}
