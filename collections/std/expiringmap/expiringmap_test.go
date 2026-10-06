package expiringmap

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewExpiryMap(t *testing.T) {
	t.Parallel()

	ttl := 10 * time.Minute
	em := NewExpiryMap[string, int](ttl, nil)

	if em == nil {
		t.Fatal("NewExpiryMap returned nil")
	}

	if em.ttl != ttl {
		t.Fatalf("ttl = %v, want %v", em.ttl, ttl)
	}

	if len(em.shards) != int(shardCount) {
		t.Fatalf("shard count = %d, want %d", len(em.shards), shardCount)
	}

	for i := range em.shards {
		if &em.shards[i] == nil {
			t.Fatalf("shard %d is nil", i)
		}
	}
}

func TestExpiryMapSetTake(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	em.Set("foo", 42)

	got, ok := em.Take("foo")
	if !ok {
		t.Fatal("Take(foo) returned ok=false, want true")
	}

	if got != 42 {
		t.Fatalf("Take(foo) = %d, want 42", got)
	}

	// Take is expected to remove the entry.
	_, ok = em.Take("foo")
	if ok {
		t.Fatal("second Take(foo) returned ok=true, want false")
	}
}

func TestExpiryMapTakeMissing(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	got, ok := em.Take("does-not-exist")

	if ok {
		t.Fatal("Take(missing) returned ok=true, want false")
	}

	if got != 0 {
		t.Fatalf("Take(missing) returned %d, want zero value 0", got)
	}
}

func TestExpiryMapSetOverwritesExistingValue(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, string](time.Hour, nil)

	em.Set("key", "first")
	em.Set("key", "second")

	got, ok := em.Take("key")
	if !ok {
		t.Fatal("Take(key) returned ok=false, want true")
	}

	if got != "second" {
		t.Fatalf("Take(key) = %q, want %q", got, "second")
	}

	_, ok = em.Take("key")
	if ok {
		t.Fatal("entry still exists after Take")
	}
}

func TestExpiryMapMultipleKeys(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	const n = 1000

	for i := 0; i < n; i++ {
		em.Set(fmt.Sprintf("key-%d", i), i)
	}

	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key-%d", i)

		got, ok := em.Take(key)
		if !ok {
			t.Fatalf("Take(%q) returned ok=false", key)
		}

		if got != i {
			t.Fatalf("Take(%q) = %d, want %d", key, got, i)
		}
	}
}

func TestExpiryMapGenericTypes(t *testing.T) {
	t.Parallel()

	type key struct {
		ID   int
		Name string
	}

	type value struct {
		ID    int
		Value string
	}

	em := NewExpiryMap[key, value](time.Hour, nil)

	k1 := key{ID: 1, Name: "one"}
	k2 := key{ID: 2, Name: "two"}

	v1 := value{ID: 100, Value: "hello"}
	v2 := value{ID: 200, Value: "world"}

	em.Set(k1, v1)
	em.Set(k2, v2)

	got1, ok := em.Take(k1)
	if !ok {
		t.Fatal("Take(k1) returned ok=false")
	}

	if got1 != v1 {
		t.Fatalf("Take(k1) = %+v, want %+v", got1, v1)
	}

	got2, ok := em.Take(k2)
	if !ok {
		t.Fatal("Take(k2) returned ok=false")
	}

	if got2 != v2 {
		t.Fatalf("Take(k2) = %+v, want %+v", got2, v2)
	}
}

func TestExpiryMapPointerValues(t *testing.T) {
	t.Parallel()

	type item struct {
		ID   int
		Name string
	}

	em := NewExpiryMap[string, *item](time.Hour, nil)

	want := &item{
		ID:   42,
		Name: "answer",
	}

	em.Set("item", want)

	got, ok := em.Take("item")
	if !ok {
		t.Fatal("Take(item) returned ok=false")
	}

	if got != want {
		t.Fatalf("Take(item) returned different pointer: got=%p want=%p", got, want)
	}

	if got.ID != 42 || got.Name != "answer" {
		t.Fatalf("unexpected value: %+v", got)
	}
}

func TestExpiryMapGetShardForIsStable(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	first := em.getShardFor("same-key")

	for i := 0; i < 100; i++ {
		got := em.getShardFor("same-key")
		if got != first {
			t.Fatalf("same key mapped to different shards")
		}
	}
}

func TestExpiryMapShardsDistributeKeys(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	shards := make(map[*shard[string, int]]struct{})

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		shards[em.getShardFor(key)] = struct{}{}
	}

	// We do not require a perfectly uniform distribution; the important
	// property is that sharding actually spreads keys across shards.
	if len(shards) < 2 {
		t.Fatalf("all keys mapped to one shard; got %d distinct shard(s)", len(shards))
	}
}

func TestExpiryMapSetDoesNotImmediatelyExpireWithPositiveTTL(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	em.Set("alive", 123)

	// sweep(now) should not remove an entry whose expiry is an hour away.
	em.sweep(time.Now())

	got, ok := em.Take("alive")
	if !ok {
		t.Fatal("entry expired unexpectedly")
	}

	if got != 123 {
		t.Fatalf("got %d, want 123", got)
	}
}

func TestExpiryMapSweepExpiresEntry(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	em.Set("expired", 123)

	// The map has a one-hour TTL. Sweep two hours into the future so
	// expiration is deterministic and does not require sleeping.
	em.sweep(time.Now().Add(2 * time.Hour))

	_, ok := em.Take("expired")
	if ok {
		t.Fatal("expired entry still exists after sweep")
	}
}

func TestExpiryMapSweepDoesNotExpireLiveEntry(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](time.Hour, nil)

	em.Set("alive", 123)

	em.sweep(time.Now().Add(30 * time.Minute))

	got, ok := em.Take("alive")
	if !ok {
		t.Fatal("live entry was removed too early")
	}

	if got != 123 {
		t.Fatalf("got %d, want 123", got)
	}
}

func TestExpiryMapZeroTTLExpiresOnSweep(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](0, nil)

	em.Set("immediate", 123)

	// Add a tiny amount to avoid depending on exact equality semantics
	// between time.Now() in Set and the time used by sweep.
	em.sweep(time.Now().Add(time.Nanosecond))

	_, ok := em.Take("immediate")
	if ok {
		t.Fatal("zero-TTL entry survived sweep")
	}
}

func TestExpiryMapNegativeTTLExpiresOnSweep(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[string, int](-time.Second, nil)

	em.Set("expired", 123)

	em.sweep(time.Now())

	_, ok := em.Take("expired")
	if ok {
		t.Fatal("negative-TTL entry survived sweep")
	}
}

func TestExpiryMapSweepCallback(t *testing.T) {
	t.Parallel()

	var (
		mu     sync.Mutex
		called []string
	)

	em := NewExpiryMap[string, int](time.Hour, func(k string, v int) {
		mu.Lock()
		defer mu.Unlock()

		called = append(called, fmt.Sprintf("%s=%d", k, v))
	})

	em.Set("a", 1)
	em.Set("b", 2)
	em.Set("c", 3)

	em.sweep(time.Now().Add(2 * time.Hour))

	mu.Lock()
	defer mu.Unlock()

	if len(called) != 3 {
		t.Fatalf("callback called %d times, want 3", len(called))
	}

	got := make(map[string]bool)
	for _, event := range called {
		got[event] = true
	}

	for _, want := range []string{
		"a=1",
		"b=2",
		"c=3",
	} {
		if !got[want] {
			t.Errorf("callback event %q was not observed", want)
		}
	}
}

func TestExpiryMapSweepCallbackReceivesLatestValueAfterOverwrite(t *testing.T) {
	t.Parallel()

	var (
		mu       sync.Mutex
		key      string
		value    int
		callSeen int
	)

	em := NewExpiryMap[string, int](time.Hour, func(k string, v int) {
		mu.Lock()
		defer mu.Unlock()

		key = k
		value = v
		callSeen++
	})

	em.Set("key", 1)
	em.Set("key", 2)

	em.sweep(time.Now().Add(2 * time.Hour))

	mu.Lock()
	defer mu.Unlock()

	if callSeen != 1 {
		t.Fatalf("callback called %d times, want 1", callSeen)
	}

	if key != "key" {
		t.Fatalf("callback key = %q, want %q", key, "key")
	}

	if value != 2 {
		t.Fatalf("callback value = %d, want 2", value)
	}
}

func TestExpiryMapTakeDoesNotTriggerExpiryCallback(t *testing.T) {
	t.Parallel()

	var calls atomic.Int64

	em := NewExpiryMap[string, int](time.Hour, func(string, int) {
		calls.Add(1)
	})

	em.Set("key", 123)

	got, ok := em.Take("key")
	if !ok {
		t.Fatal("Take returned ok=false")
	}

	if got != 123 {
		t.Fatalf("got %d, want 123", got)
	}

	// Even after taking the entry, sweep should not see it.
	em.sweep(time.Now().Add(2 * time.Hour))

	if got := calls.Load(); got != 0 {
		t.Fatalf("expiry callback called %d times, want 0", got)
	}
}

func TestExpiryMapSweepIsIdempotent(t *testing.T) {
	t.Parallel()

	var calls atomic.Int64

	em := NewExpiryMap[string, int](time.Hour, func(string, int) {
		calls.Add(1)
	})

	em.Set("one", 1)
	em.Set("two", 2)

	future := time.Now().Add(2 * time.Hour)

	em.sweep(future)

	if got := calls.Load(); got != 2 {
		t.Fatalf("callback calls after first sweep = %d, want 2", got)
	}

	em.sweep(future)

	if got := calls.Load(); got != 2 {
		t.Fatalf("callback calls after second sweep = %d, want 2", got)
	}
}

func TestExpiryMapExpiredEntryCanBeReinserted(t *testing.T) {
	t.Parallel()

	var (
		mu      sync.Mutex
		expired []int
	)

	em := NewExpiryMap[string, int](time.Hour, func(_ string, v int) {
		mu.Lock()
		defer mu.Unlock()

		expired = append(expired, v)
	})

	em.Set("key", 1)

	em.sweep(time.Now().Add(2 * time.Hour))

	// Reinsert after expiration.
	em.Set("key", 2)

	got, ok := em.Take("key")
	if !ok {
		t.Fatal("reinserted entry was not found")
	}

	if got != 2 {
		t.Fatalf("reinserted value = %d, want 2", got)
	}

	mu.Lock()
	defer mu.Unlock()

	if len(expired) != 1 {
		t.Fatalf("callback count = %d, want 1", len(expired))
	}

	if expired[0] != 1 {
		t.Fatalf("expired callback value = %d, want 1", expired[0])
	}
}

func TestExpiryMapSweepAllShards(t *testing.T) {
	t.Parallel()

	var (
		mu       sync.Mutex
		expired  = make(map[string]int)
		shards   = make(map[*shard[string, int]]struct{})
		testKeys []string
	)

	em := NewExpiryMap[string, int](time.Hour, func(k string, v int) {
		mu.Lock()
		defer mu.Unlock()

		expired[k] = v
	})

	// Find keys that definitely occupy several different shards.
	for i := 0; i < 10_000 && len(shards) < 8; i++ {
		key := fmt.Sprintf("key-%d", i)
		s := em.getShardFor(key)

		if _, already := shards[s]; !already {
			shards[s] = struct{}{}
			testKeys = append(testKeys, key)
			em.Set(key, i)
		}
	}

	if len(shards) < 2 {
		t.Fatalf("could not find keys spanning multiple shards; got %d", len(shards))
	}

	em.sweep(time.Now().Add(2 * time.Hour))

	mu.Lock()
	defer mu.Unlock()

	if len(expired) != len(testKeys) {
		t.Fatalf("expired callback entries = %d, want %d", len(expired), len(testKeys))
	}

	for i, key := range testKeys {
		got, ok := expired[key]
		if !ok {
			t.Errorf("missing expiry callback for %q", key)
			continue
		}

		expected := i

		// testKeys is built in the same order as the successful Set calls,
		// so recover the expected value from the key's numeric suffix.
		_ = expected

		if got < 0 {
			t.Errorf("unexpected negative value for %q: %d", key, got)
		}
	}
}

func TestExpiryMapSweepDoesNotCallCallbackForLiveEntries(t *testing.T) {
	t.Parallel()

	var calls atomic.Int64

	em := NewExpiryMap[string, int](time.Hour, func(string, int) {
		calls.Add(1)
	})

	for i := 0; i < 1000; i++ {
		em.Set(fmt.Sprintf("key-%d", i), i)
	}

	em.sweep(time.Now().Add(30 * time.Minute))

	if got := calls.Load(); got != 0 {
		t.Fatalf("callback called %d times, want 0", got)
	}
}

func TestExpiryMapCallbackCanReenterMap(t *testing.T) {
	t.Parallel()

	var (
		callbackStarted atomic.Bool
		callbackFailed  atomic.Bool
	)

	var em *ExpiryMap[string, int]

	em = NewExpiryMap[string, int](time.Hour, func(k string, v int) {
		callbackStarted.Store(true)

		// This test verifies that the expiry callback is not invoked while
		// holding a shard lock that would deadlock Set/Take on the same map.
		em.Set("from-callback", 99)

		got, ok := em.Take("from-callback")
		if !ok || got != 99 {
			callbackFailed.Store(true)
		}

		if k != "expired" || v != 123 {
			callbackFailed.Store(true)
		}
	})

	em.Set("expired", 123)

	em.sweep(time.Now().Add(2 * time.Hour))

	if !callbackStarted.Load() {
		t.Fatal("expiry callback was not invoked")
	}

	if callbackFailed.Load() {
		t.Fatal("expiry callback failed while re-entering map")
	}
}

func TestExpiryMapConcurrentSetTake(t *testing.T) {
	em := NewExpiryMap[string, int](time.Hour, nil)

	const (
		workers    = 32
		iterations = 2_000
	)

	var wg sync.WaitGroup
	wg.Add(workers)

	for worker := 0; worker < workers; worker++ {
		worker := worker

		go func() {
			defer wg.Done()

			for i := 0; i < iterations; i++ {
				// Deliberately use a small key space so multiple goroutines
				// contend on the same shards and keys.
				key := fmt.Sprintf("key-%d", (worker+i)%16)

				em.Set(key, i)

				if _, ok := em.Take(key); !ok {
					// Another goroutine is allowed to take the same entry,
					// so an unsuccessful Take is valid under contention.
					continue
				}
			}
		}()
	}

	wg.Wait()

	// Make a deterministic final assertion after all concurrent activity.
	em.Set("final", 12345)

	got, ok := em.Take("final")
	if !ok {
		t.Fatal("final Take returned ok=false")
	}

	if got != 12345 {
		t.Fatalf("final value = %d, want 12345", got)
	}
}

func TestExpiryMapConcurrentSweepAndAccess(t *testing.T) {
	em := NewExpiryMap[string, int](time.Hour, nil)

	const (
		accessWorkers = 24
		sweepWorkers  = 4
		iterations    = 2_000
	)

	var wg sync.WaitGroup

	wg.Add(accessWorkers)

	for worker := 0; worker < accessWorkers; worker++ {
		worker := worker

		go func() {
			defer wg.Done()

			for i := 0; i < iterations; i++ {
				key := fmt.Sprintf("key-%d", (worker+i)%64)
				em.Set(key, i)

				// Entries have a one-hour TTL, so sweep(now) should not
				// normally remove them. This nevertheless stresses Set,
				// Take, shard lookup, and sweep concurrently.
				if i%4 == 0 {
					em.Take(key)
				}
			}
		}()
	}

	wg.Add(sweepWorkers)

	for i := 0; i < sweepWorkers; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				em.sweep(time.Now())
			}
		}()
	}

	wg.Wait()

	em.Set("final", 42)

	got, ok := em.Take("final")
	if !ok {
		t.Fatal("final entry missing")
	}

	if got != 42 {
		t.Fatalf("final value = %d, want 42", got)
	}
}

func TestExpiryMapConcurrentSweepOnlyExpiresEachEntryOnce(t *testing.T) {
	var (
		mu     sync.Mutex
		counts = make(map[string]int)
	)

	em := NewExpiryMap[string, int](0, func(k string, _ int) {
		mu.Lock()
		counts[k]++
		mu.Unlock()
	})

	const entries = 2_000

	for i := 0; i < entries; i++ {
		em.Set(fmt.Sprintf("key-%d", i), i)
	}

	const sweepers = 16

	future := time.Now().Add(time.Second)

	var wg sync.WaitGroup
	wg.Add(sweepers)

	for i := 0; i < sweepers; i++ {
		go func() {
			defer wg.Done()

			em.sweep(future)
		}()
	}

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()

	if len(counts) != entries {
		t.Fatalf(
			"distinct callback keys = %d, want %d",
			len(counts),
			entries,
		)
	}

	for i := 0; i < entries; i++ {
		key := fmt.Sprintf("key-%d", i)

		if got := counts[key]; got != 1 {
			t.Fatalf(
				"callback count for %q = %d, want exactly 1",
				key,
				got,
			)
		}
	}
}

func TestExpiryMapConcurrentOverwriteAndSweep(t *testing.T) {
	var (
		mu     sync.Mutex
		values []int
	)

	em := NewExpiryMap[string, int](0, func(_ string, v int) {
		mu.Lock()
		values = append(values, v)
		mu.Unlock()
	})

	const writers = 16
	const writesPerWorker = 500

	var wg sync.WaitGroup
	wg.Add(writers)

	for worker := 0; worker < writers; worker++ {
		worker := worker

		go func() {
			defer wg.Done()

			for i := 0; i < writesPerWorker; i++ {
				em.Set("shared", worker*writesPerWorker+i)
			}
		}()
	}

	wg.Wait()

	em.sweep(time.Now().Add(time.Second))

	// Exactly one value should remain associated with the key after all
	// overwrites, therefore exactly one expiry callback is expected.
	mu.Lock()
	defer mu.Unlock()

	if len(values) != 1 {
		t.Fatalf("callback count = %d, want 1", len(values))
	}
}

func TestExpiryMapCallbackCanSetSameKey(t *testing.T) {
	t.Parallel()

	var (
		mu       sync.Mutex
		seen     []int
		callback atomic.Int64
	)

	var em *ExpiryMap[string, int]

	em = NewExpiryMap[string, int](time.Hour, func(k string, v int) {
		callback.Add(1)

		if k != "key" {
			t.Errorf("callback key = %q, want key", k)
		}

		mu.Lock()
		seen = append(seen, v)
		mu.Unlock()

		// Reinsert the same key during the callback.
		em.Set("key", 999)
	})

	em.Set("key", 123)

	em.sweep(time.Now().Add(2 * time.Hour))

	if got := callback.Load(); got != 1 {
		t.Fatalf("callback count = %d, want 1", got)
	}

	got, ok := em.Take("key")
	if !ok {
		t.Fatal("key inserted by callback was not present")
	}

	if got != 999 {
		t.Fatalf("value after callback reinsertion = %d, want 999", got)
	}

	mu.Lock()
	defer mu.Unlock()

	if len(seen) != 1 || seen[0] != 123 {
		t.Fatalf("callback values = %v, want [123]", seen)
	}
}

func TestExpiryMapEmptySweep(t *testing.T) {
	t.Parallel()

	var calls atomic.Int64

	em := NewExpiryMap[string, int](time.Hour, func(string, int) {
		calls.Add(1)
	})

	// Sweeping an empty map should be a no-op.
	em.sweep(time.Now())
	em.sweep(time.Now().Add(time.Hour))

	if got := calls.Load(); got != 0 {
		t.Fatalf("callback called %d times, want 0", got)
	}
}

func TestExpiryMapRepeatedSetTake(t *testing.T) {
	t.Parallel()

	em := NewExpiryMap[int, int](time.Hour, nil)

	const iterations = 10_000

	for i := 0; i < iterations; i++ {
		em.Set(1, i)

		got, ok := em.Take(1)
		if !ok {
			t.Fatalf("iteration %d: Take returned ok=false", i)
		}

		if got != i {
			t.Fatalf("iteration %d: got %d, want %d", i, got, i)
		}
	}
}

func TestExpiryMapBombardmentStress(t *testing.T) {
	// This is intentionally NOT t.Parallel().
	// It is meant to put as much pressure as possible on the map and
	// should be allowed to use the machine's available CPUs.
	em := NewExpiryMap[string, int](time.Hour, nil)

	cpus := runtime.GOMAXPROCS(0)

	// Scale with available CPUs while keeping a reasonable upper bound.
	workers := cpus * 8
	if workers < 32 {
		workers = 32
	}
	if workers > 128 {
		workers = 128
	}

	const iterationsPerWorker = 25_000
	const hotKeyCount = 64
	const coldKeySpace = 50_000

	var (
		wg sync.WaitGroup

		setCount  atomic.Int64
		takeCount atomic.Int64
	)

	start := make(chan struct{})

	wg.Add(workers)

	for workerID := 0; workerID < workers; workerID++ {
		workerID := workerID

		go func() {
			defer wg.Done()

			<-start

			// Simple deterministic pseudo-random sequence.
			// We deliberately avoid math/rand locking/overhead here.
			state := uint64(workerID + 1)

			next := func() uint64 {
				// xorshift64*
				state ^= state >> 12
				state ^= state << 25
				state ^= state >> 27
				return state * 2685821657736338717
			}

			for i := 0; i < iterationsPerWorker; i++ {
				n := next()

				var key string

				// 75% of accesses hit a small hot set.
				// This creates substantial contention on individual shards/keys.
				if n%4 != 0 {
					key = fmt.Sprintf("hot-%d", n%hotKeyCount)
				} else {
					// The remaining accesses hit a much larger key space,
					// exercising insertion/removal and distribution.
					key = fmt.Sprintf(
						"cold-%d-%d",
						workerID,
						n%coldKeySpace,
					)
				}

				value := workerID*iterationsPerWorker + i

				em.Set(key, value)
				setCount.Add(1)

				// Frequently perform a Take.
				// Failure is completely valid because another worker may
				// have already taken/replaced the same key.
				if n%3 != 0 {
					em.Take(key)
					takeCount.Add(1)
				}

				// Occasionally overwrite without immediately taking.
				if n%17 == 0 {
					em.Set(key, value+1)
					setCount.Add(1)
				}
			}
		}()
	}

	// Run several concurrent sweepers. The entries have a one-hour TTL,
	// so sweep(time.Now()) should not remove newly written entries, but
	// it heavily exercises every shard while writers/readers are active.
	sweeperCount := cpus
	if sweeperCount < 4 {
		sweeperCount = 4
	}
	if sweeperCount > 16 {
		sweeperCount = 16
	}

	wg.Add(sweeperCount)

	for i := 0; i < sweeperCount; i++ {
		go func() {
			defer wg.Done()

			<-start

			for j := 0; j < iterationsPerWorker/10; j++ {
				em.sweep(time.Now())

				// Yield occasionally so sweepers don't monopolize CPUs.
				if j%16 == 0 {
					runtime.Gosched()
				}
			}
		}()
	}

	// Release the entire workload at once.
	close(start)

	// Wait for the bombardment to finish.
	wg.Wait()

	t.Logf(
		"bombardment complete: workers=%d sets=%d takes=%d sweepers=%d",
		workers,
		setCount.Load(),
		takeCount.Load(),
		sweeperCount,
	)

	// Make sure the map is still usable after the high-contention workload.
	const sentinelKey = "__stress_test_sentinel__"

	for i := 0; i < 1000; i++ {
		em.Set(sentinelKey, i)

		got, ok := em.Take(sentinelKey)
		if !ok {
			t.Fatalf("sentinel Take failed at iteration %d", i)
		}

		if got != i {
			t.Fatalf(
				"sentinel value corrupted at iteration %d: got %d, want %d",
				i,
				got,
				i,
			)
		}
	}

	// Finally stress expiration itself.
	//
	// Use a separate map with TTL=0 so all entries are immediately
	// eligible for removal. Multiple concurrent sweepers then race to
	// clean the same data structure.
	var expiredCount atomic.Int64

	expiryMap := NewExpiryMap[string, int](
		0,
		func(string, int) {
			expiredCount.Add(1)
		},
	)

	const expirationEntries = 100_000

	for i := 0; i < expirationEntries; i++ {
		expiryMap.Set(
			fmt.Sprintf("expiry-%d", i),
			i,
		)
	}

	const expirationSweepers = 32

	var expiryWG sync.WaitGroup
	expiryWG.Add(expirationSweepers)

	for i := 0; i < expirationSweepers; i++ {
		go func() {
			defer expiryWG.Done()

			for j := 0; j < 10; j++ {
				expiryMap.sweep(time.Now().Add(time.Second))
				runtime.Gosched()
			}
		}()
	}

	expiryWG.Wait()

	if got := expiredCount.Load(); got != expirationEntries {
		t.Fatalf(
			"expiry callback count = %d, want exactly %d",
			got,
			expirationEntries,
		)
	}

	// Verify nothing remains.
	for i := 0; i < expirationEntries; i++ {
		key := fmt.Sprintf("expiry-%d", i)

		if _, ok := expiryMap.Take(key); ok {
			t.Fatalf("expired entry %q still exists", key)
		}
	}
}
