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
	data := &ints{1, 2, 3, 4, 5}
	var j int // Index returnd from search

	for i := range *data {
		if j = Search(i+1, data); i != j {
			t.Fatalf("TestSearch: expected %d, received %d\n", i, j)
		}
	}
}

func BenchmarkInsertionsort(b0 *testing.B) {
	data := [5]*ints{}
	var n int // Largest index

	for i := range data {
		data[i] = reversedInts(int(math.Pow10(i)))
		n = data[i].Length() - 1

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
		n = data[i].Length() - 1

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

// func BenchmarkSearch(b0 *testing.B) {
// 	data := [5]*ints{}
// 	var n int // Largest index

// 	for i := range data {
// 		data[i] = sortedInts(int(math.Pow10(i)))
// 		n = data[i].Length() - 1

// 		b0.Run(
// 			fmt.Sprintf("QuickSortable on size 10^%d", i+1),
// 			func(b1 *testing.B) {
// 				for j := 0; j < b1.N; j++ {
// 					Search(data[i])
// 				}
// 			},
// 		)
// 	}
// }

/*
func BenchmarkInsertionsort2(b0 *testing.B) {
	data := [16]*ints{}
	var n int
	for i := range data {
		data[i] = reversedInts(2 * i)
		n = data[i].length() - 1
		b0.Run(
			fmt.Sprintf("Insertionsort on size %d", 2*i), func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					insertionsort(copyInts(data[i]), 0, n)
				}
			},
		)
	}
}

func BenchmarkQuicksort2(b0 *testing.B) {
	data := [16]*ints{}
	var n int
	for i := range data {
		data[i] = reversedInts(2 * i)
		n = data[i].length() - 1
		b0.Run(
			fmt.Sprintf("Quicksort on size %d", 2*i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					quicksort(copyInts(data[i]), 0, n)
				}
			},
		)
	}
}
*/
