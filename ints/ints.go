package ints

import "math/rand"

// Ints demonstrates sorting functionality. It serves as an example
// for using the sort library.
type Ints []int

// New returns the make function on Ints.
func New(len, cap int) Ints {
	return make(Ints, len, cap)
}

// Random returns a random ordering of Sorted.
func Random(n int) Ints {
	A := New(0, n)
	for i := 0; i < n; i++ {
		A = append(A, i)
		A.Swap(i, rand.Intn(i+1))
	}

	return A
}

// Reversed returns Sorted, but reversed.
func Reversed(n int) Ints {
	A := New(0, n)
	for 0 < n {
		n--
		A = append(A, n)
	}

	return A
}

// Sorted returns [0, 1, ..., n-1].
func Sorted(n int) Ints {
	A := New(0, n)
	for i := 0; i < n; i++ {
		A = append(A, i)
	}

	return A
}

// Compare two indexed integers.
func (A *Ints) Compare(i, j int) int {
	switch {
	case (*A)[i] < (*A)[j]:
		return -1
	case (*A)[j] < (*A)[i]:
		return 1
	default:
		return 0
	}
}

// CompareAt returns the comparison of A[i] to an integer.
func (A *Ints) CompareAt(i int, x interface{}) int {
	switch {
	case (*A)[i] < x.(int):
		return -1
	case x.(int) < (*A)[i]:
		return 1
	default:
		return 0
	}
}

// Copy returns a copy of a set of integers.
func (A *Ints) Copy() *Ints {
	B := make(Ints, A.Len())
	copy(B, *A)
	return &B
}

// Len returns the number of integers.
func (A *Ints) Len() int {
	return len(*A)
}

// Less returns A[i] < A[j].
func (A *Ints) Less(i, j int) bool {
	return (*A)[i] < (*A)[j]
}

// Swap two indexed integers.
func (A *Ints) Swap(i, j int) {
	(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
}

// Pop removes and returns an integer from a set of integers.
func (A *Ints) Pop() interface{} {
	n := len(*A) - 1
	a := (*A)[n]
	*A = (*A)[:n]
	return a
}

// Push inserts an integer into the last position.
func (A *Ints) Push(x interface{}) {
	*A = append(*A, x.(int))
}
