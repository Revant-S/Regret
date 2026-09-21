package impl

import "cmp"

type color bool

const (
	red   color = true
	black color = false
)

type setNode[T cmp.Ordered] struct {
	val   T
	left  *setNode[T]
	right *setNode[T]
	color color
}

type Set[T cmp.Ordered] struct {
	root *setNode[T]
	size int
}

func NewSet[T cmp.Ordered]() *Set[T] {
	return &Set[T]{
		size: 0,
	}
}

func (s *Set[T]) Size() int {
	return s.size
}

func (s *Set[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Set[T]) Clear() {
	s.root = nil
	s.size = 0
}

// Contains checks if the value exists in the set O(log N)
func (s *Set[T]) Contains(val T) bool {
	curr := s.root
	for curr != nil {
		if val < curr.val {
			curr = curr.left
		} else if val > curr.val {
			curr = curr.right
		} else {
			return true
		}
	}
	return false
}

// Values returns all elements in sorted order O(N)
func (s *Set[T]) Values() []T {
	result := make([]T, 0, s.size)
	var inOrder func(n *setNode[T])
	inOrder = func(n *setNode[T]) {
		if n == nil {
			return
		}
		inOrder(n.left)
		result = append(result, n.val)
		inOrder(n.right)
	}
	inOrder(s.root)
	return result
}

// Insert adds a value to the set O(log N)
func (s *Set[T]) Insert(val T) {
	if s.Contains(val) {
		return
	}
	s.root = s.insertNode(s.root, val)
	s.root.color = black
	s.size++
}

func (s *Set[T]) Remove(val T) {
	if !s.Contains(val) {
		return
	}
	if !isRed(s.root.left) && !isRed(s.root.right) {
		s.root.color = red
	}
	s.root = s.removeNode(s.root, val)
	if s.root != nil {
		s.root.color = black
	}
	s.size--
}

func isRed[T cmp.Ordered](n *setNode[T]) bool {
	if n == nil {
		return false
	}
	return n.color == red
}

func rotateLeft[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	return x
}

func rotateRight[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	return x
}

func flipColors[T cmp.Ordered](h *setNode[T]) {
	h.color = !h.color
	if h.left != nil {
		h.left.color = !h.left.color
	}
	if h.right != nil {
		h.right.color = !h.right.color
	}
}

func fixUp[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	if isRed(h.right) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	return h
}

func (s *Set[T]) insertNode(h *setNode[T], val T) *setNode[T] {
	if h == nil {
		return &setNode[T]{val: val, color: red}
	}
	if val < h.val {
		h.left = s.insertNode(h.left, val)
	} else if val > h.val {
		h.right = s.insertNode(h.right, val)
	}
	return fixUp(h)
}

func moveRedLeft[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

func moveRedRight[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

func deleteMin[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	if h.left == nil {
		return nil
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return fixUp(h)
}

func minNode[T cmp.Ordered](h *setNode[T]) *setNode[T] {
	for h.left != nil {
		h = h.left
	}
	return h
}

func (s *Set[T]) removeNode(h *setNode[T], val T) *setNode[T] {
	if val < h.val {
		if h.left != nil && !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = s.removeNode(h.left, val)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if val == h.val && h.right == nil {
			return nil
		}
		if h.right != nil && !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		if val == h.val {
			minRight := minNode(h.right)
			h.val = minRight.val
			h.right = deleteMin(h.right)
		} else {
			h.right = s.removeNode(h.right, val)
		}
	}
	return fixUp(h)
}
