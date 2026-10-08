# Requirements For The Data Structure

**Component:** `collections/delayqueue`
**Priority:** P1
**Reporter:** Routing team

---

## Background

When a send to a carrier fails with a retryable error (timeout, 429, 503) the router will try immediately. Under a carrier outage this turns one failure into a burst of retries, which makes the outage worse and pollutes the bandit's reward data with failures caused by our own retry storm.

We want retries to go out with exponential backoff: 1 s, 2 s, 4 s, up to 30 s. The same mechanism will be reused by the traffic simulator to schedule synthetic delivery receipts at realistic delays, so it has to be a general-purpose, generic component, not something retry-specific.

## Problem

We need an in-process queue where producers submit items with a delay, and worker goroutines block until the earliest item is due, then process it. Several request goroutines will submit retries concurrently, and a fixed pool of retry workers will consume them.

## API contract

The package must expose exactly this (unexported internals are up to you):

```go
package delayqueue

var ErrClosed = errors.New("delayqueue: closed")

type DelayQueue[T any] struct { /* unexported */ }

func New[T any]() *DelayQueue[T]

// Put schedules item to become available after delay.
// A delay <= 0 makes the item available immediately.
// Returns ErrClosed if the queue has been closed.
func (q *DelayQueue[T]) Put(item T, delay time.Duration) error

// Take blocks until an item is due and returns it, or returns ErrClosed.
func (q *DelayQueue[T]) Take() (T, error)

// TryTake returns a due item if one exists, without blocking.
func (q *DelayQueue[T]) TryTake() (T, bool)

// Len reports the number of items in the queue, due or not.
func (q *DelayQueue[T]) Len() int

// Close releases all blocked callers. Safe to call more than once.
func (q *DelayQueue[T]) Close()
```

## Functional requirements

1. Items are returned in order of ready time. Items with equal ready times may come out in any order.
2. `Take` must never return an item before its ready time.
3. Every item that was successfully `Put` is returned by `Take`/`TryTake` at most once.
4. A `Put` whose ready time is earlier than everything currently queued must be picked up by an already-blocked `Take` at its own ready time, not the previous earliest one.
5. After `Close`:
    - every goroutine blocked in `Take` returns promptly;
    - `Put` returns `ErrClosed`;
    - define and document what happens to items still queued, and what `Take` does with them.
6. `Close` may be called concurrently from multiple goroutines without panicking.

## Non-functional requirements

- Safe for concurrent use by any number of producers and consumers.
- No goroutine leaks: once `Close` returns, any goroutine the queue started has exited or will exit without further input.
- No busy-waiting or polling loops. An idle queue with blocked takers should use no CPU.
- A `Take` should wake within 10 ms of its item's ready time under normal load.
- `Put` and `Take` should be O(log n) in the number of queued items.
- Standard library only.
