package queue

type node[T any] struct {
	value T
	next  *node[T]
}

// Queue  FIFO Data structure
// Zero Value is an empty queue
type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	len  int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	n := &node[T]{value: v}
	if q.tail == nil {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.len++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	n := q.head
	q.head = n.next
	if q.head == nil {
		q.tail = nil
	}
	n.next = nil
	q.len--
	return n.value, true
}

// Peek returns the value at the front of the queue without removing it.
func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	return q.head.value, true
}

func (q *Queue[T]) Len() int {
	return q.len
}
