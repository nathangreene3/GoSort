package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	data := [5]*ints{}

	for i := range data {
		data[i] = reversedInts(int(math.Pow(10, float64(i))))
		sort(data[i])
		if !isSorted(data[i]) {
			t.Fatalf("TestSort: %v\n", *data[i])
		}
	}
}

func TestStable(t *testing.T) {
	data := [5]*ints{}

	for i := range data {
		data[i] = reversedInts(int(math.Pow(10, float64(i))))
		stable(data[i])
		if !isSorted(data[i]) {
			t.Fatalf("TestSort: %v\n", *data[i])
		}
	}
}

func TestSearch(t *testing.T) {

}

func BenchmarkInsertionsort(b0 *testing.B) {
	data := [5]*ints{}
	var n int

	for i := range data {
		data[i] = reversedInts(int(math.Pow(10, float64(i))))
		n = data[i].length() - 1
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
	var n int

	for i := range data {
		data[i] = reversedInts(int(math.Pow(10, float64(i))))
		n = data[i].length() - 1
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
