package collections

type Stack[T any] interface {
	IsEmpty() bool
	Push(val T)
	Pop()
	Size() int
	Top() T
	Clear()
}
