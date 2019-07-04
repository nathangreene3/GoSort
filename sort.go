package sort

// Interface defines the contract for sorting items.
type Interface interface {
	Compare(i, j int) int
	CompareAt(i int, x interface{}) int
	Len() int
	Swap(i, j int)
}

// Sort data.
func Sort(A Interface) {
	quicksort(A, 0, A.Len()-1)
}

// Stable sort data.
func Stable(A Interface) {
	insertionsort(A, 0, A.Len()-1)
}

// insertionsort data on the range [a,b].
func insertionsort(A Interface, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.Compare(j, j+1); j-- {
			A.Swap(j, j+1)
		}
	}
}

// quicksort data on the range [a,b].
func quicksort(A Interface, a, b int) {
	if a < b {
		if b-a < 9 {
			insertionsort(A, a, b)
		} else {
			p := pivot(A, a, b)
			quicksort(A, a, p-1)
			quicksort(A, p+1, b)
		}
	}
}

// pivot pivots data on the range [a,b] by selecting the pivot index
// by the median-of-three method. The pivot index is returned.
func pivot(A Interface, a, b int) int {
	medianOfThree(A, a, b)
	p := a
	for i := a + 1; i <= b; i++ {
		if A.Compare(i, a) < 0 {
			p++
			A.Swap(i, p)
		}
	}

	A.Swap(a, p)
	return p
}

// medianOfThree sets the a-th value to be the median, and the
// (a+b)/2-th and b-th value to be in order.
func medianOfThree(A Interface, a, b int) {
	// Example: [9,7,5] becomes [7,5,9].

	c := int(uint(a+b) >> 1) // (a+b)/2
	if A.Compare(a, b) < 0 {
		A.Swap(a, b)
	}

	if A.Compare(c, a) < 0 {
		A.Swap(a, c)
	}

	if A.Compare(b, c) < 0 {
		A.Swap(b, c)
	}
}

// medianOfFive sets the median of five values to the a-th index. The
// other four values may not be in order.
func medianOfFive(A Interface, a, b int) {
	// This doesn't seem to improve anything. It's worse than medianOfThree.
	c := int(uint(a+b) >> 1) // (a+b)/2
	if 4 < b-a {
		medianOfThree(A, c, b) // Put median of [c:b] to c
		medianOfThree(A, a, c) // Then put median of [a:c] to a
	} else {
		medianOfThree(A, c, b) // Put median of [c:b] to c
	}
}

// IsSorted determines if data is sorted.
func IsSorted(A Interface) bool {
	n := A.Len() - 1
	for i := 0; i < n; i++ {
		if 0 < A.Compare(i, i+1) {
			return false
		}
	}

	return true
}

// Search returns the index an item belongs in a list of items and
// whether or not it was found.
func Search(A Interface, x interface{}) (int, bool) {
	var (
		r int           // Comparison result
		a int           // Lower index
		b int           // Middle index
		c = A.Len() - 1 // Upper index
	)
	for a <= c {
		b = int(uint(a+c) >> 1) // (i + k) / 2
		r = A.CompareAt(b, x)
		switch {
		case r < 0:
			a = b + 1
		case 0 < r:
			c = b - 1
		default:
			return b, true
		}
	}

	return a, false
}
