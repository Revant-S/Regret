package delayqueue

import (
	"errors"
	"sync"
)

var ErrClosed = errors.New("delayQueue: closed")

//
//type DelayQueue[T any] struct { /* unexported */ }
//
//func New[T any]() *DelayQueue[T]
//
//// Put schedules item to become available after delay.
//// A delay <= 0 makes the item available immediately.
//// Returns ErrClosed if the queue has been closed.
//func (q *DelayQueue[T]) Put(item T, delay time.Duration) error
//
//// Take blocks until an item is due and returns it, or returns ErrClosed.
//func (q *DelayQueue[T]) Take() (T, error)
//
//// TryTake returns a due item if one exists, without blocking.
//func (q *DelayQueue[T]) TryTake() (T, bool)
//
//// Len reports the number of items in the queue, due or not.
//func (q *DelayQueue[T]) Len() int
//
//// Close releases all blocked callers. Safe to call more than once.
//func (q *DelayQueue[T]) Close()

//We need an in-process queue where producers submit items with a delay,
//and worker goroutines block until the earliest item is due, then process it.
//Several request goroutines will submit retries concurrently, and a fixed pool of retry workers will consume them.

//type minHeap[V any] []*item[V]

type DelayQueue[V any] struct {
	heap minHeap[V]
	mu   sync.Mutex
	cond *sync.Cond
}

func New[V any]() *DelayQueue[V] {
	dq := DelayQueue[V]{
		heap: minHeap[V]{},
	}
	dq.cond = sync.NewCond(&dq.mu)
	return &dq
}

func (dq *DelayQueue[V]) Put(item item[V]) {
	
}
