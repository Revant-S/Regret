package impl

type stackNode[T any] struct {
	value T
	prev  *stackNode[T]
}

type Stack[T any] struct {
	top  *stackNode[T]
	head *stackNode[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{size: 0}
}

func (st *Stack[T]) IsEmpty() bool {
	return st.size == 0
}

func (st *Stack[T]) Push(val T) {
	node := &stackNode[T]{value: val}
	if st.IsEmpty() {
		st.head = node
		st.top = node
	} else {
		node.prev = st.top
		st.top = node
	}
	st.size++
}

func (st *Stack[T]) Pop() {
	if st.IsEmpty() {
		panic("stack is empty")
	}
	st.top = st.top.prev
	st.size--
	if st.top == nil {
		st.head = nil
	}
}

func (st *Stack[T]) Size() int {
	return st.size
}

func (st *Stack[T]) Top() T {
	if st.IsEmpty() {
		panic("empty stack")
	}
	return st.top.value
}

func (st *Stack[T]) Clear() {
	st.head = nil
	st.top = nil
	st.size = 0
}
