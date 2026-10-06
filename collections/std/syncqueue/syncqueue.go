package syncqueue

import (
	"Regret/collections/std/queue"
	"sync"
)

// SyncQueue it is a concurrency safe FIFO data structure
// It uses sync.Mutex to lock every method called on it
type SyncQueue[T any] struct {
	mu sync.RWMutex
	q  queue.Queue[T]
}

func (sq *SyncQueue[T]) Enqueue(value T) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.q.Enqueue(value)
}

func (sq *SyncQueue[T]) Dequeue() (T, bool) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	if sq.q.Len() == 0 {
		var zero T
		return zero, false
	}
	val, ok := sq.q.Dequeue()
	return val, ok
}

func (sq *SyncQueue[T]) Front() (T, bool) {
	sq.mu.RLock()
	defer sq.mu.RUnlock()
	if sq.q.Len() == 0 {
		var zero T
		return zero, false
	}
	val, ok := sq.q.Dequeue()
	return val, ok
}
