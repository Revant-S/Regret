package expiringmap

import (
	"hash/maphash"
	"time"
)

/*
 Design:-
- Should support concurrency
- The entries are sharded
- All shards are managed by the data structure
- the package should provide only an API for the pool manager
- This is a generic data structure
- One can set a ttl and pass the call back function on the expiry

- Any Map should contain multiple shards
- Shards are essentially the map with mutex

- Number of shards is managed internally
*/

const shardCount uint64 = 32

// ExpiryMap is a concurrent map whose entries expire a fixed time after they
// are set. It is safe for use by multiple goroutines.
//
// Each entry leaves the map exactly once: either a caller removes it with
// Take, or it expires and onExpiry is called with its key and value. Never
// both
//
// Expiry is not instant. Expired entries are removed by a background sweep,
// so an entry can stay in the map for up to one sweep interval past its
// deadline. Until it is swept, Take still returns it.
//
// onExpiry runs on the sweeper's goroutine with no locks held, so it may
// call Set or Take on the same map. It should return quickly, since a slow
// callback delays expiry of every other entry.
//
// The zero value is not usable; create an ExpiryMap with NewExpiryMap.
// An ExpiryMap must not be copied after first use.
type ExpiryMap[K comparable, V any] struct {
	seed     maphash.Seed
	shards   []shard[K, V]
	ttl      time.Duration
	onExpiry func(K, V)
}

func (em *ExpiryMap[K, V]) getShardFor(key K) *shard[K, V] {
	h := maphash.Comparable(em.seed, key)
	return &em.shards[h&(shardCount-1)]
}
func NewExpiryMap[K comparable, V any](ttl time.Duration, onExpiry func(K, V)) *ExpiryMap[K, V] {
	em := &ExpiryMap[K, V]{
		seed:     maphash.MakeSeed(),
		ttl:      ttl,
		onExpiry: onExpiry,
		shards:   make([]shard[K, V], shardCount),
	}
	for i := range em.shards {
		em.shards[i] = newShard[K, V]()
	}
	return em
}

func (em *ExpiryMap[K, V]) Set(key K, value V) {
	em.getShardFor(key).set(key, value, time.Now().Add(em.ttl))
}

func (em *ExpiryMap[K, V]) Take(key K) (V, bool) {
	return em.getShardFor(key).take(key)
}

func (em *ExpiryMap[K, V]) sweep(now time.Time) {
	for i := range em.shards {
		for k, v := range em.shards[i].removeExpired(now) {
			if em.onExpiry != nil {
				em.onExpiry(k, v)
			}
		}
	}
}
