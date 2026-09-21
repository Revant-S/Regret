package std

import "cmp"

type Set[T cmp.Ordered] interface {
	Insert(val T)
	Remove(val T)
	Contains(val T) bool
	Size() int
	IsEmpty() bool
	Clear()
	Values() []T // Returns elements in sorted order (in-order traversal)
}
