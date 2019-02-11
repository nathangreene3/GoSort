package main

// sortable data can be sorted using GoSort methods.
type sortable interface {
	length() int        // Number of items to be sorted
	less(i, j int) bool // Less than condition for item comparison
	swap(i, j int)      // Swaps two items
}

// sortSortable is the general function to export that calls an internal sort
// function.
func sortSortable(A sortable, a, b int) {
	bmax := A.length() - 1
	if a < 0 {
		a = 0
	}
	if bmax < b {
		b = bmax
	}
	if b-a < 16 {
		insertionSortable(A, a, b)
	} else {
		quickSortable(A, a, b)
		// bubbleSortable(A, a, b)
		// insertionSortable(A, a, b)
		// mergeSortable(A, a, b)
	}
}

// bubbleSortable sorts sortable data on the range [a,b].
func bubbleSortable(A sortable, a, b int) {
	c := true
	for m := b; c; m-- {
		c = false
		for i := a; i < m; i++ {
			if A.less(i+1, i) {
				A.swap(i, i+1)
				c = true
			}
		}
	}
}

// insertionSortable sorts sortable data on the range [a,b].
func insertionSortable(A sortable, a, b int) sortable {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && A.less(j+1, j); j-- {
			A.swap(j, j+1)
		}
	}
	return A
}

// quickSortable sorts sortable data on the range [a,b].
func quickSortable(A sortable, a, b int) sortable {
	if a < b {
		p := pivotSortable(A, a, b)
		A = quickSortable(A, a, p-1)
		A = quickSortable(A, p+1, b)
	}
	return A
}

// pivotSortable pivots sortable data on the range [a,b] by selecting the
// pivot index by the median-of-three method.
func pivotSortable(A sortable, a, b int) int {
	A = medianOfThreeSortable(A, a, b)
	p := a
	for i := a + 1; i <= b; i++ {
		if A.less(i, a) {
			p++
			A.swap(i, p)
		}
	}
	A.swap(a, p)
	return p
}

// medianOfThree sorts the values at indices a, b, and (a+b)/2.
func medianOfThreeSortable(A sortable, a, b int) sortable {
	c := (a + b) / 2
	if A.less(a, b) {
		A.swap(a, b)
	}
	if A.less(c, a) {
		A.swap(a, c)
	}
	if A.less(b, c) {
		A.swap(b, c)
	}
	return A
}

func mergeSortable(A sortable, a, b int) {
	if a < b {
		m := (a + b) / 2
		mergeSortable(A, a, m)
		mergeSortable(A, m+1, b)
		mergeable(A, a, m, m+1, b)
	}
}

func mergeable(A sortable, a0, b0, a1, b1 int) {
	// a := a0
	// b := b1
	// c := 0
	// R := make([]int, len(A))
	// for a0 <= b0 && a1 <= b1 {
	// 	if A[a0] < A[a1] {
	// 		R[c] = A[a0]
	// 		a0++
	// 	} else {
	// 		R[c] = A[a1]
	// 		a1++
	// 	}
	// 	c++
	// }
	// if a0 <= b0 {
	// 	for i := a0; i <= b0; i++ {
	// 		R[c] = A[i]
	// 		c++
	// 	}
	// } else {
	// 	for i := a1; i <= b1; i++ {
	// 		R[c] = A[i]
	// 		c++
	// 	}
	// }
	// c = 0
	// for i := a; i <= b; i++ {
	// 	A[i] = R[c]
	// 	c++
	// }
}