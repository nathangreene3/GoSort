package sort

// Interface defines the contract for sorting items.
type Interface interface {
	Compare(i, j int) int
	CompareAt(i int, x interface{}) int
	Len() int
	Swap(i, j int)
}

// Sort sortable data.
func Sort(A Interface) {
	quicksort(A, 0, A.Len()-1)
}

// Stable sort sortable data.
func Stable(A Interface) {
	insertionsort(A, 0, A.Len()-1)
}

// insertionsort sorts sortable data on the range [a,b].
func insertionsort(A Interface, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.Compare(j, j+1); j-- {
			A.Swap(j, j+1)
		}
	}
}

// quicksort sorts sortable data on the range [a,b].
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

// pivot pivots sortable data on the range [a,b] by selecting the
// pivot index by the median-of-three method.
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

// medianOfThree sorts the values at indices a, b, and (a+b)/2.
func medianOfThree(A Interface, a, b int) {
	c := int(uint(a+b) >> 1)
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

// IsSorted determines if a sortable object is sorted.
func IsSorted(A Interface) bool {
	n := A.Len() - 1
	for i := 0; i < n; i++ {
		if 0 < A.Compare(i, i+1) {
			return false
		}
	}

	return true
}

// Search returns the index an item belongs in a sortable set and whether or not it was found.
func Search(A Interface, x interface{}) (int, bool) {
	var (
		c int           // Comparison result
		i int           // Lower index
		j int           // Middle index
		k = A.Len() - 1 // Upper index
	)
	for i <= k {
		j = (i + k) / 2
		c = A.CompareAt(j, x)
		switch {
		case c < 0:
			i = j + 1
		case 0 < c:
			k = j - 1
		default:
			return j, true
		}
	}

	return i, false
}
