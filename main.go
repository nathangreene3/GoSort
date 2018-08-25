package main

import (
	"fmt"
	"math/rand"
)

type sortable interface {
	length() int
	less(i, j int) bool
	swap(i, j int)
}

type person struct {
	first string
	last  string
}

type people []person

func main() {
	A := randomSlice(25)
	B := copySlice(A)
	fmt.Println("A:", A)
	fmt.Println(mergeSort(A, 0, len(A)-1))
	fmt.Println("B:", B)
	fmt.Println(quickSort(B, 0, len(B)-1))
	ppl := people{
		person{first: "e", last: "E"},
		person{first: "b", last: "B"},
		person{first: "c", last: "C"},
		person{first: "a", last: "A"},
		person{first: "d", last: "D"},
	}
	fmt.Println(experBubbleSort(ppl))
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

func pivot(A []int, a, b int) int {
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
		j := i
		for 0 < j {
			if a < A[j-1] {
				A[j] = A[j-1]
			}
			j--
		}
		A[j] = a
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

func experBubbleSort(A sortable) sortable {
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
