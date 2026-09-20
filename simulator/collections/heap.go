package collections

type HeapType int

const (
	MinHeap HeapType = 0
	MaxHeap HeapType = 1
)

type Heap[T any] struct {
	pqArray  T[]
	size     int
	heapType HeapType
}

func NewHeap[D any](T HeapType) *Heap[D] {
	return &Heap[D]{
		heapType: T,
	}
}
