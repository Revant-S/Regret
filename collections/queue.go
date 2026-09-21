package collections

type Queue[T any] interface {
	Enqueue(v T)
	Front() T
	Dequeue()
	IsEmpty() bool
	Size() int
}
