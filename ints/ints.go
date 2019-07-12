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

// ------------------------------
// Functional Programming methods
// ------------------------------

// Iterator performs some action given an index i. Returns true if iteration is to continue and false if iteration is to halt early.
type Iterator func(i int) bool

// Iterate over range [a,b) and while f(a) is true.
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

// FPQuicksort ...
func (A Ints) FPQuicksort() Ints {
	if n := len(A); 1 < n {
		var (
			p    = n >> 1 // n/2
			B, C = make(Ints, 0, n), make(Ints, 0, n)
			f    = func(i int) bool {
				if A[i] <= A[p] {
					B = append(B, A[i])
				} else {
					C = append(C, A[i])
				}
				return true
			}
		)

		Iterate(f, 0, n)
		return merge(B.FPQuicksort(), C.FPQuicksort())
	}

	return A
}

// merge A and B into a new, sorted Ints. A and B must be sorted.
func merge(A, B Ints) Ints {
	var (
		a, b int
		m, n = len(A), len(B)
		C    = make(Ints, 0, m+n)
	)

	if m == 0 {
		if n == 0 {
			return Ints{}
		}
		return B
	}

	if n == 0 {
		return A
	}

	f := func(i int) bool {
		switch {
		case a < m:
			if b < n && B[b] < A[a] {
				C = append(C, B[b])
				b++
			} else {
				C = append(C, A[a])
				a++
			}
			return true
		case b < n:
			if a < m && A[a] < B[b] {
				C = append(C, A[a])
				a++
			} else {
				C = append(C, B[b])
				b++
			}
			return true
		default:
			return false
		}
	}
	Iterate(f, 0, m+n)
	return C
}
