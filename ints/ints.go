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

// Iterator ...
type Iterator func(i int) bool

// Iterate ...
func Iterate(f Iterator, a, b int) {
	if a < b && f(a) {
		Iterate(f, a+1, b)
	}
}

// Mapper ...
type Mapper func(a int) int

// Map ...
func (A Ints) Map(f Mapper) Ints {
	var (
		n = len(A)
		B = make(Ints, 0, n)
		g = func(i int) bool {
			B = append(B, f(A[i]))
			return true
		}
	)
	Iterate(g, 0, n)
	return B
}

// Filterer ...
type Filterer func(a int) bool

// Filter ...
func (A Ints) Filter(f Filterer) Ints {
	var (
		n = len(A)
		B = make(Ints, 0, n)
		g = func(i int) bool {
			a := A[i]
			if f(a) {
				B = append(B, a)
			}
			return true
		}
	)
	Iterate(g, 0, n)
	return B
}

// Reducer ...
type Reducer func(a, b int) int

// Reduce ...
func (A Ints) Reduce(f Reducer) int {
	var (
		v int
		g = func(i int) bool {
			v = f(v, A[i])
			return true
		}
	)
	Iterate(g, 0, len(A))
	return v
}

//
func (A Ints) FPSort(a, b int) Ints {
	p := (b + a) / 2
	if a == p {
		return A
	}

	B, C := make(Ints, 0, p), make(Ints, 0, p)
	// f:=func(i int)bool{
	// 	if A[i]<A[]
	// }
	return B.Append(C)
}

// Append ...
func (A Ints) Append(B Ints) Ints {
	var (
		m, n = len(A), len(B)
		C    = make(Ints, 0, m+n)
		f    = func(i int) bool {
			if i < m {
				C = append(C, A[i])
			} else {
				C = append(C, B[i])
			}
			return true
		}
	)
	Iterate(f, 0, m+n)
	return C
}
