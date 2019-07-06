package heap

import (
	"github.com/nathangreene3/sort"
)

// Interface defines the context for implementing a heap.
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

// Heapify sets up an Interface as a min heap.
func Heapify(h Interface) {
	n := h.Len() - 1
	for i := n / 2; 0 <= i; i-- {
		siftDown(h, i, n)
	}
}

// Push an item onto the heap.
func Push(h Interface, x interface{}) {
	h.Push(x)
	siftUp(h, h.Len()-1)
}

// Pop the top of the heap. The item will have the lowest value of the heap.
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	siftDown(h, 0, n-1)
	return h.Pop()
}

// siftUp corrects the heap from i up.
func siftUp(h Interface, i int) {
	if 0 < i {
		p := int(uint(i-1) >> 1) // Parent index: (i-1)/2
		if 0 < h.Compare(p, i) {
			h.Swap(p, i)
			siftUp(h, p)
		}
	}
}

// siftDown corrects the heap from i down to n.
func siftDown(h Interface, i, n int) {
	j := int(uint(i)<<1) + 1 // Left child index
	k := j + 1               // Right child index
	if j <= n {
		if k <= n && 0 < h.Compare(j, k) {
			j = k
		}

		if 0 < h.Compare(i, j) {
			h.Swap(i, j) // Swap with largest child
			siftDown(h, j, n)
		}
	}
}
