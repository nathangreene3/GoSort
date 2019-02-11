package main

import (
	"fmt"
	"math/rand"
	"sort"
	srt "sort"
	"testing"
)

func randomBts(n int) []byte {
	bts := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		bts = append(bts, byte(rand.Intn(10)))
	}
	return bts
}

func TestKeySortToIncrementPositions(t *testing.T) {
	data := []struct {
		test         []byte
		keyResult    []int
		incPosResult []int
		size         int
	}{
		// {randomBts(10), nil, nil, 10},
		// {randomBts(100), nil, nil, 100},
		// {randomBts(1000), nil, nil, 1000},
		// {randomBts(10000), nil, nil, 10000},
		{[]byte("1230"), nil, nil, 4},
	}
	for i := range data {
		data[i].keyResult = keySort(data[i].test, nil, 0, data[i].size-1)
		data[i].incPosResult = incrementPositions(data[i].test)

		if len(data[i].keyResult) != data[i].size {
			t.Fatalf("keySort returned length %d instead of %d\n", len(data[i].keyResult), data[i].size)
		}
		if len(data[i].incPosResult) != data[i].size {
			t.Fatalf("incrementPositions returned length %d instead of %d\n", len(data[i].incPosResult), data[i].size)
		}

		for j := range data[i].keyResult {
			if data[i].keyResult[j] != data[i].incPosResult[j] {
				t.Fatalf("\nkeySort result:           %v\nincrementPosition result: %v\n", data[i].keyResult, data[i].incPosResult)
			}
		}
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
