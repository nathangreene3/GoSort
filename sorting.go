package main

import "math/rand"

// sortable data can be sorted using GoSort methods.
type sortable interface {
	length() int        // Number of items to be sorted
	less(i, j int) bool // Less than condition for item comparison
	swap(i, j int)      // Swaps two items
}

type intSlice []int

// person is a first and last name collection.
type person struct {
	first string
	last  string
}

// people is a collection of persons.
type people []person
type pplByFirst people
type pplByLast people

func charAt(i int, key string) (s string) {
	return string(key[i%len(key)])
}

func (ppl people) length() (length int) {
	return len(ppl)
}

func (ppl people) less(i, j int) (less bool) {
	if ppl[i].last < ppl[j].last {
		less = true
	}
	return less
}

func (ppl people) swap(i, j int) {
	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func mergeSort(A []int, a, b int) []int {
	if a < b {
		m := (a + b) / 2
		mergeSort(A, a, m)
		mergeSort(A, m+1, b)
		merge(A, a, m, m+1, b)
	}
	return A
}

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

func quickSort(A []int, a, b int) []int {
	if a < b {
		p := pivot(A, a, b)
		quickSort(A, a, p-1)
		quickSort(A, p+1, b)
	}
	return A
}

func pivot(A []int, a, b int) (p int) {
	medianOfThree(A, a, b)
	v := A[a]
	p = a
	for i := a + 1; i <= b; i++ {
		if A[i] < v {
			p++
			A[p], A[i] = A[i], A[p]
		}
	}
	A[a], A[p] = A[p], A[a]
	return p
}

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

func insertionSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		a := A[i]
		j := i - 1
		for 0 <= j && a < A[j] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = a
	}
	return A
}

func reversedSlice(n int) (A []int) {
	A = make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = n - i
	}
	return A
}

func randomSlice(n int) (A []int) {
	A = make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i + 1
	}
	rand.Shuffle(n, func(i, j int) { A[i], A[j] = A[j], A[i] })
	return A
}

func copySlice(A []int) (B []int) {
	B = make([]int, len(A))
	for i, a := range A {
		B[i] = a
	}
	return B
}

func (A intSlice) quickSortable(a, b int) sortable {
	if a < b {
		p := A.pivotSortable(a, b)
		A.quickSortable(a, p-1)
		A.quickSortable(p+1, b)
	}
	return A
}

func (A intSlice) pivotSortable(a, b int) (p int) {
	A.medianOfThreeSortable(a, b)
	v := A[a]
	p = a
	for i := a + 1; i <= b; i++ {
		if A[i] < v {
			p++
			A.swap(i, p)
		}
	}
	A.swap(a, p)
	return p
}

func (A intSlice) medianOfThreeSortable(a, b int) sortable {
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

func bubbleSortable(A sortable) sortable {
	c := true
	m := A.length()
	for c {
		c = false
		for i := 0; i+1 < m; i++ {
			if A.less(i+1, i) {
				A.swap(i, i+1)
				c = true
			}
		}
		m--
	}
	return A
}

func (A intSlice) length() (length int) {
	return len(A)
}

func (A intSlice) less(i, j int) (less bool) {
	return A[i] < A[j]
}

func (A intSlice) swap(i, j int) {
	A[i], A[j] = A[j], A[i]
}

func (A intSlice) at(i int) (a int) {
	return A[i]
}
