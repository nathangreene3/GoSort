package main

import (
	"math/rand"
	"time"
)

// ints implements sortable interface to demonstrate sorting functionality.
type ints []int

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func (A *ints) length() int {
	return len(*A)
}

func (A *ints) compare(i, j int) int {
	switch {
	case (*A)[i] < (*A)[j]:
		return -1
	case (*A)[j] < (*A)[i]:
		return 1
	default:
		return 0
	}
}

func (A *ints) compareTo(x interface{}, i int) int {
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

func (A *ints) swap(i, j int) {
	t := (*A)[i]
	(*A)[i] = (*A)[j]
	(*A)[j] = t
}

func reversedInts(n int) *ints {
	A := make(ints, 0, n)
	for 0 < n {
		n--
		A = append(A, n)
	}

	return &A
}

func copyInts(A *ints) *ints {
	n := A.length()
	B := make(ints, 0, n)
	for i := 0; i < n; i++ {
		B = append(B, (*A)[i])
	}

	return &B
}
