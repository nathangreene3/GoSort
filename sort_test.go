package sort

import (
	"fmt"
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	data := [5]*ints{}

	for i := range data {
		data[i] = reversedInts(int(math.Pow(10, float64(i))))
		Sort(data[i])
		if !IsSorted(data[i]) {
			t.Fatalf("TestSort: %v\n", *data[i])
		}
	}
}

func TestStable(t *testing.T) {
	data := [5]*ints{}

	for i := range data {
		data[i] = reversedInts(int(math.Pow(10, float64(i))))
		Stable(data[i])
		if !IsSorted(data[i]) {
			t.Fatalf("TestSort: %v\n", *data[i])
		}
	}
}

func TestSearch(t *testing.T) {
	var (
		data  = &ints{1, 2, 3, 4, 5}
		found bool
		j     int // Index returned from search
	)

	for i := range *data {
		if j, found = Search(data, i+1); i != j || !found {
			t.Fatalf("TestSearch: expected %d and %t, received %d and %t\n", i, true, j, found)
		}
	}

	if j, found = Search(data, 6); j != len(*data) || found {
		t.Fatalf("TestSearch: expected %d and %t, received %d and %t\n", 5, false, j, found)
	}

	if j, found = Search(data, 0); j != 0 || found {
		t.Fatalf("TestSearch: expected %d and %t, received %d and %t\n", 0, false, j, found)
	}
}

func BenchmarkInsertionsort(b0 *testing.B) {
	data := [5]*ints{}
	var n int // Largest index

	for i := range data {
		data[i] = reversedInts(int(math.Pow10(i)))
		n = data[i].Len() - 1

		b0.Run(
			fmt.Sprintf("InsertionSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					insertionsort(copyInts(data[i]), 0, n)
				}
			},
		)
	}
}

func BenchmarkQuicksort(b0 *testing.B) {
	data := [5]*ints{}
	var n int // Largest index

	for i := range data {
		data[i] = reversedInts(int(math.Pow10(i)))
		n = data[i].Len() - 1

		b0.Run(
			fmt.Sprintf("QuickSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					quicksort(copyInts(data[i]), 0, n)
				}
			},
		)
	}
}

func BenchmarkSearch(b0 *testing.B) {
	var (
		data   = [5]*ints{}// Careful... this is powers of ten
		length int
	)
	for i := range data {
		length = int(math.Pow10(i))
		data[i] = sortedInts(length)
		b0.Run(
			fmt.Sprintf("Search for 0th element on size %d", length),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					Search(data[i], 0)
				}
			},
		)
	}
}

func BenchmarkInsertionsort2(b0 *testing.B) {
	var (
		data   = [16]*ints{}
		length int
	)
	for i := range data {
		length = i + 1
		data[i] = reversedInts(length)
		b0.Run(
			fmt.Sprintf("Insertionsort on size %d", length), func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					insertionsort(copyInts(data[i]), 0, length-1)
				}
			},
		)
	}
}

func BenchmarkQuicksort2(b0 *testing.B) {
	var (
		data   = [16]*ints{}
		length int
	)
	for i := range data {
		length = i + 1
		data[i] = reversedInts(length)
		b0.Run(
			fmt.Sprintf("Quicksort on size %d", length),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					quicksort(copyInts(data[i]), 0, length-1)
				}
			},
		)
	}
}
