package main

import (
	"math"
	"math/rand"
	"sort"
)

type intSlice []int

// ----------------------------------------------------------------------
// []int sorting methods
// ----------------------------------------------------------------------

// countingSort: TODO
func countingSort(A []int) []int {
	// Find range of A to set up m
	m := make(map[int]int)
	min := A[0]
	max := A[0]
	for i := range A {
		if A[i] < min {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	for i := min; i <= max; i++ {
		m[i] = 0
	}

	// Define m to be the frequency of the range of A
	for i := range A {
		m[A[i]]++
	}

	// Redefine m to be the cumulative frequency of the range of A
	for i := min + 1; i <= max; i++ {
		m[i] += m[i-1]
	}

	B := make([]int, len(A))
	i := 0
	for k, v := range m {
		if 0 < v {
			B[i] = k
			i++
		}
	}
	return B
}

// mergeSort sorts integers on the range [a,b].
func mergeSort(A []int, a, b int) []int {
	if a < b {
		m := (a + b) / 2
		mergeSort(A, a, m)
		mergeSort(A, m+1, b)
		merge(A, a, m, m+1, b)
	}
	return A
}

// merge merges two ranges [a0,b0] and [a1,b1] into a single range within
// the set of integers.
func merge(A []int, a0, b0, a1, b1 int) []int {
	a := a0
	b := b1
	c := 0
	R := make([]int, len(A))
	for a0 <= b0 && a1 <= b1 {
		if A[a0] < A[a1] {
			R[c] = A[a0]
			a0++
		} else {
			R[c] = A[a1]
			a1++
		}
		c++
	}
	if a0 <= b0 {
		for i := a0; i <= b0; i++ {
			R[c] = A[i]
			c++
		}
	} else {
		for i := a1; i <= b1; i++ {
			R[c] = A[i]
			c++
		}
	}
	c = 0
	for i := a; i <= b; i++ {
		A[i] = R[c]
		c++
	}
	return A
}

// quickSort sorts integers on the range [a,b].
func quickSort(A []int, a, b int) []int {
	if a < b {
		p := pivot(A, a, b)
		quickSort(A, a, p-1)
		quickSort(A, p+1, b)
	}
	return A
}

// pivot sifts values less than a pivot value (selected by
// median-of-three) to the lower indices and the values larger than the
// pivot value to the higher indices on the range [a,b].
func pivot(A []int, a, b int) int {
	medianOfThree(A, a, b)
	v := A[a]
	p := a
	for i := a + 1; i <= b; i++ {
		if A[i] < v {
			p++
			A[p], A[i] = A[i], A[p]
		}
	}
	A[a], A[p] = A[p], A[a]
	return p
}

// medianOfThree sorts the values at indices a, b, and (a+b)/2.
func medianOfThree(A []int, a, b int) []int {
	c := (a + b) / 2
	if A[a] < A[b] {
		A[a], A[b] = A[b], A[a]
	}
	if A[c] < A[a] {
		A[a], A[c] = A[c], A[a]
	}
	if A[b] < A[c] {
		A[b], A[c] = A[c], A[b]
	}
	return A
}

// bubbleSort sorts integers.
func bubbleSort(A []int) []int {
	c := true
	m := len(A)
	for c {
		c = false
		for i := 0; i+1 < m; i++ {
			if A[i+1] < A[i] {
				A[i], A[i+1] = A[i+1], A[i]
				c = true
			}
		}
		m--
	}
	return A
}

// insertionSort sorts integers.
func insertionSort(A []int) []int {
	var a, j int
	for i := 1; i < len(A); i++ {
		a = A[i]
		j = i - 1
		for 0 <= j && a < A[j] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = a
	}
	return A
}

// newSlice returns [0,1,...,n-1].
func newSlice(n int) []int {
	A := make([]int, 0, n)
	for i := 0; i < n; i++ {
		A = append(A, i)
	}
	return A
}

// reversedSlice returns [n-1,...,1,0].
func reversedSlice(n int) []int {
	A := make([]int, 0, n)
	for 0 < n {
		A = append(A, n-1)
		n--
	}
	return A
}

// randomSlice returns a random permutation of [0,1,...,n-1].
func randomSlice(n int) []int {
	A := newSlice(n)
	rand.Shuffle(n, func(i, j int) { A[i], A[j] = A[j], A[i] })
	return A
}

// copySlice returns a deep copy of a set of integers.
func copySlice(A []int) []int {
	return append([]int{}, A...)
}

// ----------------------------------------------------------------------
// intSlice implements sortable interface
// ----------------------------------------------------------------------

// quickSortable sorts an intSlice on the range [a,b].
func (A intSlice) quickSortable(a, b int) {
	if a < b {
		p := A.pivotSortable(a, b)
		A.quickSortable(a, p-1)
		A.quickSortable(p+1, b)
	}
}

// pivotSortable sifts values less than a pivot value (selected by
// median-of-three) to the lower indices and the values larger than the
// pivot value to the higher indices on the range [a,b].
func (A intSlice) pivotSortable(a, b int) int {
	A.medianOfThreeSortable(a, b)
	v := A[a]
	p := a
	for i := a + 1; i <= b; i++ {
		if A[i] < v {
			p++
			A.swap(i, p)
		}
	}
	A.swap(a, p)
	return p
}

// medianOfThree sorts the values at indicies
func (A intSlice) medianOfThreeSortable(a, b int) {
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
}

// length returns the number of integers in a slice.
func (A intSlice) length() int {
	return len(A)
}
func (A intSlice) Len() int {
	return len(A)
}

// less returns the comparison of two integers.
func (A intSlice) less(i, j int) bool {
	return A[i] < A[j]
}

func (A intSlice) Less(i, j int) bool {
	return A[i] < A[j]
}

// swap swaps two values at two given indices.
func (A intSlice) swap(i, j int) {
	A[i], A[j] = A[j], A[i]
}
func (A intSlice) Swap(i, j int) {
	A[i], A[j] = A[j], A[i]
}

// at returns a value at an index.
func (A intSlice) at(i int) interface{} {
	return A[i]
}

// newIntSlice returns [0,1,...,n-1].
func newIntSlice(n int) intSlice {
	A := make(intSlice, 0, n)
	for i := 0; i < n; i++ {
		A = append(A, i)
	}
	return A
}

func copyIntSlice(A intSlice) intSlice {
	B := make(intSlice, 0, A.length())
	for i := 0; i < A.length(); i++ {
		B = append(B, A[i])
	}
	return B
}

// reversedSlice returns [n-1,...,1,0].
func reversedIntSlice(n int) intSlice {
	A := make(intSlice, 0, n)
	for 0 < n {
		A = append(A, n-1)
		n--
	}
	return A
}

// randomSlice returns a random permutation of [0,1,...,n-1].
func randomIntSlice(n int) intSlice {
	A := newIntSlice(n)
	for i := 0; i < n; i++ {
		rand.Shuffle(n, func(i, j int) { A.swap(i, j) })
	}
	return A
}

func mostUnsortedIntSlice(n int) intSlice {
	A := newIntSlice(n)
	return append(A[n/2:], A[:n/2]...)
}

func nextSlice(p []int) []int {
	n := len(p)
	k := -1
	for i := n - 2; 0 <= i; i-- {
		if p[i] < p[i+1] {
			k = i
			break
		}
	}
	if k == -1 {
		return newSlice(n)
	}
	j := -1
	for i := n - 1; k < i; i-- { // 0 <= k < n-1
		if p[k] < p[i] {
			j = i
			break
		}
	}
	q := copySlice(p)
	q[k], q[j] = q[j], q[k]
	a, b := k+1, n-1
	for a < b {
		q[a], q[b] = q[b], q[a]
		a++
		b--
	}
	return q
}

func mostUnsortedSlice(n int) []int {
	A := newSlice(n)
	return append(A[n/2:], A[:n/2]...)
}

func isSorted(A []int) bool {
	for i := len(A) - 1; 0 < i; i-- {
		if A[i] < A[i-1] {
			return false
		}
	}
	return true
}

func avgIndexError(A []int) float64 {
	n := len(A)
	B := copySlice(A) // Sorted copy
	sort.Ints(B)
	var e float64   // Index error
	var j, jmax int // Index range A[i] should be in B
	for i := 0; i < n; i++ {
		j = sort.Search(n, func(index int) bool { return A[i] <= B[index] })
		jmax = j
		for jmax+1 < n && B[j] == B[jmax+1] {
			jmax++
		}
		e += math.Abs(float64(i) - (float64(j)+float64(jmax))/2.0)
	}
	return e / float64(n)
}
