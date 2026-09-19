package collections

type queueNode[T any] struct {
	value interface{}
	next  *queueNode[T]
}
type Queue[T any] struct {
	head *queueNode[T]
	tail *queueNode[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		size: 0,
	}
}
func (q *Queue[T]) Enqueue(v T) {
	node := &queueNode[T]{value: v}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue[T]) Front() T {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.head.value
}
func (q *Queue[T]) Dequeue() {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	q.head = q.head.next
	q.size--
	if q.head == nil {
		q.tail = nil
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}
