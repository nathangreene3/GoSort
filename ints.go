package sort

// ints implements sortable interface to demonstrate sorting functionality.
type ints []int

// Length returns the number of integers.
func (A *ints) Length() int {
	return len(*A)
}

// Compare two indexed integers.
func (A *ints) Compare(i, j int) int {
	switch {
	case (*A)[i] < (*A)[j]:
		return -1
	case (*A)[j] < (*A)[i]:
		return 1
	default:
		return 0
	}
}

// CompareTo compares an integer to an indexed integer.
func (A *ints) CompareTo(x interface{}, i int) int {
	y, ok := x.(int)
	switch {
	case !ok:
		panic("ints.compareTo: value not an int")
	case y < (*A)[i]:
		return -1
	case (*A)[i] < y:
		return 1
	default:
		return 0
	}
}

// Swap two indexed integers.
func (A *ints) Swap(i, j int) {
	t := (*A)[i]
	(*A)[i] = (*A)[j]
	(*A)[j] = t
}

// sortedInts returns [0, 1, ..., n-1].
func sortedInts(n int) *ints {
	A := make(ints, 0, n)
	for i := 0; i < n; i++ {
		A = append(A, i)
	}

	return &A
}

// reversedInts returns [n-1, n-2, ..., 0].
func reversedInts(n int) *ints {
	A := make(ints, 0, n)
	for 0 < n {
		n--
		A = append(A, n)
	}

	return &A
}

// copyInts returns a copy of a set of integers.
func copyInts(A *ints) *ints {
	n := A.Length()
	B := make(ints, 0, n)
	for i := 0; i < n; i++ {
		B = append(B, (*A)[i])
	}

	return &B
}
