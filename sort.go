package sort

// Sortable data can be sorted using GoSort methods.
type Sortable interface {
	Compare(i, j int) int               // Compares two indexed items returning -1, 0, or 1
	CompareTo(x interface{}, i int) int // Compare an item to an indexed item returning -1, 0, or 1
	Length() int                        // Number of items
	Swap(i, j int)                      // Swaps two items
}

// Sort sortable data.
func Sort(A Sortable) {
	quicksort(A, 0, A.Length()-1)
}

// Stable sort sortable data.
func Stable(A Sortable) {
	insertionsort(A, 0, A.Length()-1)
}

// insertionsort sorts sortable data on the range [a,b].
func insertionsort(A Sortable, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.Compare(j, j+1); j-- {
			A.Swap(j, j+1)
		}
	}
}

// quicksort sorts sortable data on the range [a,b].
func quicksort(A Sortable, a, b int) {
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
func pivot(A Sortable, a, b int) int {
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
func medianOfThree(A Sortable, a, b int) {
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
func IsSorted(A Sortable) bool {
	n := A.Length() - 1
	for i := 0; i < n; i++ {
		if 0 < A.Compare(i, i+1) {
			return false
		}
	}

	return true
}

// Search returns the index an item belongs in a sortable set. Does not guarentee the item's existence in the set.
func Search(x interface{}, A Sortable) int {
	var (
		i int              // Lower index
		j = A.Length() - 1 // Upper index
		k int              // Middle index
	)
	for i < j {
		k = int(uint(i+j) >> 1)
		if 0 < A.CompareTo(x, k) {
			i = k + 1
		} else {
			j = k
		}
	}

	return i
}
