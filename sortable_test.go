package main

import (
	"fmt"
	srt "sort"
	"testing"
)

func BenchmarkBubbleSortable(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{reversedIntSlice(10), 10},
		{reversedIntSlice(100), 100},
		{reversedIntSlice(1000), 1000},
		{reversedIntSlice(10000), 10000},
	}
	var c intSlice
	for i := range data {
		b0.Run(
			fmt.Sprintf("BubbleSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copyIntSlice(data[i].test)
					b1.StartTimer()
					bubbleSortable(c)
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
		{reversedIntSlice(10), 10},
		{reversedIntSlice(100), 100},
		{reversedIntSlice(1000), 1000},
		{reversedIntSlice(10000), 10000},
	}
	var c intSlice
	for i := range data {
		b0.Run(
			fmt.Sprintf("InsertionSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copyIntSlice(data[i].test)
					b1.StartTimer()
					insertionSortable(c)
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
		{reversedIntSlice(10), 10},
		{reversedIntSlice(100), 100},
		{reversedIntSlice(1000), 1000},
		{reversedIntSlice(10000), 10000},
	}
	var c intSlice
	for i := range data {
		b0.Run(
			fmt.Sprintf("QuickSortable on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copyIntSlice(data[i].test)
					b1.StartTimer()
					quickSortable(c, 0, data[i].size-1)
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
		{reversedSlice(10), 10},
		{reversedSlice(100), 100},
		{reversedSlice(1000), 1000},
		{reversedSlice(10000), 10000},
	}
	var c []int
	for i := range data {
		b0.Run(
			fmt.Sprintf("BubbleSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copySlice(data[i].test)
					b1.StartTimer()
					bubbleSort(c)
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
		{reversedSlice(10), 10},
		{reversedSlice(100), 100},
		{reversedSlice(1000), 1000},
		{reversedSlice(10000), 10000},
	}
	var c []int
	for i := range data {
		b0.Run(
			fmt.Sprintf("InsertionSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copyIntSlice(data[i].test)
					b1.StartTimer()
					insertionSort(c)
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
		{reversedSlice(10), 10},
		{reversedSlice(100), 100},
		{reversedSlice(1000), 1000},
		{reversedSlice(10000), 10000},
	}
	var c []int
	for i := range data {
		b0.Run(
			fmt.Sprintf("QuickSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copyIntSlice(data[i].test)
					b1.StartTimer()
					quickSort(c, 0, data[i].size-1)
				}
			},
		)
	}
}

func BenchmarkGosSort(b0 *testing.B) {
	data := []struct {
		test intSlice
		size int
	}{
		{reversedSlice(10), 10},
		{reversedSlice(100), 100},
		{reversedSlice(1000), 1000},
		{reversedSlice(10000), 10000},
	}
	var c []int
	for i := range data {
		b0.Run(
			fmt.Sprintf("GosSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					b1.StopTimer()
					c = copyIntSlice(data[i].test)
					b1.StartTimer()
					srt.Ints(c)
				}
			},
		)
	}
}
