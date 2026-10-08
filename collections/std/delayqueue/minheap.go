package delayqueue

import "time"

type item[V any] struct {
	value   V
	readyAt time.Time
	seq     int64
}

type minHeap[V any] struct {
	items []*item[V]
}

func (mh *minHeap[V]) Len() int {
	return len(mh.items)
}

func (mh *minHeap[V]) Less(i, j int) bool {
	if !mh.items[i].readyAt.Equal(mh.items[j].readyAt) {
		return mh.items[i].readyAt.Before(mh.items[j].readyAt)
	}
	return mh.items[i].seq < mh.items[j].seq
}

func (mh *minHeap[V]) Swap(i, j int) {
	mh.items[i], mh.items[j] = mh.items[j], mh.items[i]
}

func (mh *minHeap[V]) Push(x any) {
	it := x.(*item[V])
	mh.items = append(mh.items, it)
}

func (mh *minHeap[V]) Pop() any {
	old := mh.items
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	mh.items = old[0 : n-1]
	return it
}
