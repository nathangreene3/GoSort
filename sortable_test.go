package main

import (
	"fmt"
	"sort"
	srt "sort"
	"testing"
)

func BenchmarkBubbleSortable(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedIntSlice(10), 10},
		{mostUnsortedIntSlice(100), 100},
		{mostUnsortedIntSlice(1000), 1000},
		{mostUnsortedIntSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("BubbleSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					bubbleSortable(copyIntSlice(data[i].test), 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkInsertionSortable(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedIntSlice(10), 10},
		{mostUnsortedIntSlice(100), 100},
		{mostUnsortedIntSlice(1000), 1000},
		{mostUnsortedIntSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("InsertionSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					insertionSortable(copyIntSlice(data[i].test), 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkQuickSortable(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedIntSlice(10), 10},
		{mostUnsortedIntSlice(100), 100},
		{mostUnsortedIntSlice(1000), 1000},
		{mostUnsortedIntSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("QuickSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					quickSortable(copyIntSlice(data[i].test), 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkGosSortable(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedIntSlice(10), 10},
		{mostUnsortedIntSlice(100), 100},
		{mostUnsortedIntSlice(1000), 1000},
		{mostUnsortedIntSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("GosSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					sort.Sort(data[i].test)
				}
			},
		)
	}
}

func BenchmarkBubbleSort(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedSlice(10), 10},
		{mostUnsortedSlice(100), 100},
		{mostUnsortedSlice(1000), 1000},
		{mostUnsortedSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("BubbleSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					bubbleSort(copyIntSlice(data[i].test))
				}
			},
		)
	}
}

func BenchmarkInsertionSort(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedSlice(10), 10},
		{mostUnsortedSlice(100), 100},
		{mostUnsortedSlice(1000), 1000},
		{mostUnsortedSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("InsertionSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					insertionSort(copyIntSlice(data[i].test))
				}
			},
		)
	}
}

func BenchmarkQuickSort(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedSlice(10), 10},
		{mostUnsortedSlice(100), 100},
		{mostUnsortedSlice(1000), 1000},
		{mostUnsortedSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("QuickSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					quickSort(copyIntSlice(data[i].test), 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkGosInts(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{mostUnsortedSlice(10), 10},
		{mostUnsortedSlice(100), 100},
		{mostUnsortedSlice(1000), 1000},
		{mostUnsortedSlice(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("GosSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					srt.Ints(copyIntSlice(data[i].test))
				}
			},
		)
	}
}
