package main

import (
	"fmt"
	"testing"
)

type ints []int

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

func TestSort(t *testing.T) {
	data := []struct {
		test *ints
		size int
	}{
		{reversedInts(1), 1},
		{reversedInts(10), 10},
		{reversedInts(100), 100},
		{reversedInts(1000), 1000},
		{reversedInts(10000), 10000},
	}

	for i := range data {
		if isSorted(data[i].test) {
			t.Fatalf("TestSort: %v\n", *data[i].test)
		}
	}
}

func BenchmarkInsertionSortable(b0 *testing.B) {
	data := []struct {
		test *ints
		size int
	}{
		{reversedInts(1), 1},
		{reversedInts(10), 10},
		{reversedInts(100), 100},
		{reversedInts(1000), 1000},
		{reversedInts(10000), 10000},
	}

	for i := range data {
		b0.Run(
			fmt.Sprintf("InsertionSortable on size 10^%d", i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					insertionSortable(copyInts(data[i].test), 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkQuickSortable(b0 *testing.B) {
	data := []struct {
		test *ints
		size int
	}{
		{reversedInts(1), 1},
		{reversedInts(10), 10},
		{reversedInts(100), 100},
		{reversedInts(1000), 1000},
		{reversedInts(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("QuickSortable on size 10^%d", i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					quickSortable(copyInts(data[i].test), 0, data[i].size-1)
				}
			},
		)
	}
}
