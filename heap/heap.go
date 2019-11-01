package heap

import "github.com/nathangreene3/sort"

// Heapable defines the context for implementing a heap.
type Heapable interface {
	sort.Sortable
	Push(x interface{})
	Pop() interface{}
}

// Heapify sets up an Heapable as a max heap.
func Heapify(h Heapable) {
	n := h.Len() - 1
	for i := n >> 1; 0 <= i; i-- {
		siftDown(h, i, n)
	}
}

// Push an item onto the heap.
func Push(h Heapable, x interface{}) {
	h.Push(x)
	siftUp(h, h.Len()-1)
}

// Pop the top of the heap. The item will have the lowest value of the heap.
func Pop(h Heapable) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	siftDown(h, 0, n-1)
	return h.Pop()
}

// siftUp corrects the heap from i up.
func siftUp(h Heapable, i int) {
	if 0 < i {
		p := int(uint(i-1) >> 1) // Parent index: (i-1)/2
		if 0 < h.Compare(p, i) {
			h.Swap(p, i)
			siftUp(h, p)
		}
	}
}

// siftDown corrects the heap from i down to n.
func siftDown(h Heapable, i, n int) {
	// j is initially left child, k is right child
	if j := int(uint(i)<<1) + 1; j <= n {
		k := j + 1
		if k <= n && h.Compare(j, k) < 0 {
			j = k // right child is larger
		}

		if h.Compare(i, j) < 0 {
			h.Swap(i, j) // Swap with largest child
			siftDown(h, j, n)
		}
	}
}

// Sort the heap.
func Sort(h Heapable) {
	Heapify(h)
	for n := h.Len() - 1; 0 < n; {
		h.Swap(0, n)
		n--
		siftDown(h, 0, n)
	}
}
