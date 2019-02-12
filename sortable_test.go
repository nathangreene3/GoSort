package main

import (
	"fmt"
	"sort"
	srt "sort"
	"testing"

	"github.com/cheynewallace/tabby"
)

func TestIndexMap(t *testing.T) {
	data := []struct {
		test   []byte
		result []int
		size   int
	}{
		{randomBts(0), nil, 0},
		{randomBts(1), nil, 0},
		{randomBts(10), nil, 0},
		{randomBts(100), nil, 0},
		{randomBts(1000), nil, 0},
		{randomBts(10000), nil, 0},
	}

	table := tabby.New()
	table.AddHeader("INDEX", "INPUT", "RESULT")
	var failed bool

	for i := range data {
		data[i].size = len(data[i].test)
		data[i].result = indexMap(data[i].test, 0, data[i].size-1)

		if len(data[i].result) != data[i].size {
			table.AddLine(i, data[i].test, data[i].result)
			failed = true
			continue
		}

		for j := 0; j < data[i].size-1; j++ {
			if data[i].test[data[i].result[j+1]] < data[i].test[data[i].result[j]] {
				table.AddLine(i, data[i].test, data[i].result)
				failed = true
			}
		}
	}

	if failed {
		table.Print()
		t.Fatal("\n")
	}
}

func TestIndexMapToIncrementPositions(t *testing.T) {
	data := []struct {
		test           []byte
		indexMapResult []int
		incPosResult   []int
		size           int
	}{
		{randomBts(0), nil, nil, 0},
		{randomBts(1), nil, nil, 0},
		{randomBts(2), nil, nil, 0},
		{randomBts(3), nil, nil, 0},
		{randomBts(4), nil, nil, 0},
		{randomBts(5), nil, nil, 0},
		{randomBts(6), nil, nil, 0},
		{randomBts(7), nil, nil, 0},
		{randomBts(8), nil, nil, 0},
		{randomBts(9), nil, nil, 0},
	}

	table := tabby.New()
	table.AddHeader("FUNCTION", "INPUT", "RESULT")
	var failed bool

	for i := range data {
		data[i].size = len(data[i].test)
		data[i].indexMapResult = indexMap(data[i].test, 0, data[i].size-1)
		data[i].incPosResult = incrementPositions(data[i].test)

		if len(data[i].indexMapResult) != data[i].size || len(data[i].incPosResult) != data[i].size {
			table.AddLine("GetIncrementPositions", data[i].test, data[i].incPosResult)
			table.AddLine("indexMap", data[i].test, data[i].indexMapResult)
			failed = true
			continue
		}

		for j := range data[i].indexMapResult {
			if data[i].indexMapResult[j] != data[i].incPosResult[j] {
				table.AddLine("GetIncrementPositions", data[i].test, data[i].incPosResult)
				table.AddLine("indexMap", data[i].test, data[i].indexMapResult)
				failed = true
			}
		}
	}
	if failed {
		table.Print()
		// t.Fatal("\n")
	}
}

func BenchmarkKeySortable(b0 *testing.B) {
	data := []struct {
		test []byte
		size int
	}{
		{randomBts(10), 10},
		{randomBts(100), 100},
		{randomBts(1000), 1000},
		{randomBts(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("KeySort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					keySort(data[i].test, nil, 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkIndexMap(b0 *testing.B) {
	data := []struct {
		test []byte
		size int
	}{
		{randomBts(10), 10},
		{randomBts(100), 100},
		{randomBts(1000), 1000},
		{randomBts(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("indexMap on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					indexMap(data[i].test, 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkIncrementPositions(b0 *testing.B) {
	data := []struct {
		test []byte
		size int
	}{
		{randomBts(10), 10},
		{randomBts(100), 100},
		{randomBts(1000), 1000},
		{randomBts(10000), 10000},
	}
	for i := range data {
		b0.Run(
			fmt.Sprintf("IncrementPositions on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					incrementPositions(data[i].test)
				}
			},
		)
	}
}

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
