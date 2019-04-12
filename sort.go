package main

// sortable data can be sorted using GoSort methods.
type sortable interface {
	compare(i, j int) int               // Compares two indexed items returning -1, 0, or 1
	compareTo(x interface{}, i int) int // Compare an item to an indexed item returning -1, 0, or 1
	length() int                        // Number of items
	swap(i, j int)                      // Swaps two items
}

// sort sortable data.
func sort(A sortable) {
	quicksort(A, 0, A.length()-1)
}

// stable sort sortable data.
func stable(A sortable) {
	insertionsort(A, 0, A.length()-1)
}

// insertionsort sorts sortable data on the range [a,b].
func insertionsort(A sortable, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.compare(j, j+1); j-- {
			A.swap(j, j+1)
		}
	}
}

// quicksort sorts sortable data on the range [a,b].
func quicksort(A sortable, a, b int) {
	if a < b {
		p := pivot(A, a, b)
		quicksort(A, a, p-1)
		quicksort(A, p+1, b)
	}
}

// pivot pivots sortable data on the range [a,b] by selecting the
// pivot index by the median-of-three method.
func pivot(A sortable, a, b int) int {
	medianOfThree(A, a, b)
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
func medianOfThree(A sortable, a, b int) {
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

// search returns the index an item belongs in a sortable set. Does not guarentee the item's existence in the set.
func search(x interface{}, A sortable, a, b int) int {
	var c int
	for a < b {
		c = int(uint(a+b) >> 1)
		if 0 < A.compareTo(x, c) {
			a = c + 1
		} else {
			b = c
		}
	}

	return a
}

// reverse sort a range o
func reverse(A sortable, a, b int) {
	quicksort(A, a, b)
	for a < b {
		A.swap(a, b)
		a++
		b--
	}
}
