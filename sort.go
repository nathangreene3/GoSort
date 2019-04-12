package main

// sortable data can be sorted using GoSort methods.
type sortable interface {
	compare(i, j int) int // Compares two items returning -1, 0, or 1
	length() int          // Number of items
	swap(i, j int)        // Swaps two items
}

// isSorted determines if a sortable object is sorted on the range [a,b].
func isSorted(A sortable) bool {
	n := A.length() - 1
	for a := 0; a < n; a++ {
		if 0 < A.compare(a, a+1) {
			return false
		}
	}

	return true
}

// sortSortable is the general function to export that calls an internal sort
// function.
func sortSortable(A sortable) sortable {
	return quickSortable(A, 0, A.length()-1)
}

// insertionSortable sorts sortable data on the range [a,b].
func insertionSortable(A sortable, a, b int) sortable {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.compare(j, j+1); j-- {
			A.swap(j, j+1)
		}
	}

	return A
}

// quickSortable sorts sortable data on the range [a,b].
func quickSortable(A sortable, a, b int) sortable {
	if a < b {
		if b-a < 16 {
			return insertionSortable(A, a, b)
		}

		p := pivotSortable(A, a, b)
		quickSortable(A, a, p-1)
		quickSortable(A, p+1, b)
	}

	return A
}

// pivotSortable pivots sortable data on the range [a,b] by selecting the
// pivot index by the median-of-three method.
func pivotSortable(A sortable, a, b int) int {
	A = medianOfThreeSortable(A, a, b)
	p := a
	for i := a + 1; i <= b; i++ {
		if A.compare(i, a) < 0 {
			p++
			A.swap(i, p)
		}
	}

	A.swap(a, p)
	return p
}

// medianOfThree sorts the values at indices a, b, and (a+b)/2.
func medianOfThreeSortable(A sortable, a, b int) sortable {
	c := int(uint(a+b) >> 1)
	if A.compare(a, b) < 0 {
		A.swap(a, b)
	}

	if A.compare(c, a) < 0 {
		A.swap(a, c)
	}

	if A.compare(b, c) < 0 {
		A.swap(b, c)
	}

	return A
}

// search returns the index an item belongs in a sortable set. Does not guarentee the item's existence in the set.
func search(x interface{}, A sortable, a, b int) int {
	return 0
}
