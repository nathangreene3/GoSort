package sort

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/nathangreene3/sort/ints"
)

func TestSort(t *testing.T) {
	data := [5]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(10, float64(i))))
		Sort(&data[i])
		if !IsSorted(&data[i]) {
			t.Fatalf("TestSort: %v\n", data[i])
		}
	}
}

func TestStable(t *testing.T) {
	data := [5]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(10, float64(i))))
		Stable(&data[i])
		if !IsSorted(&data[i]) {
			t.Fatalf("TestSort: %v\n", data[i])
		}
	}
}

func TestShellsort(t *testing.T) {
	data := [5]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(10, float64(i))))
		shellsort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatalf("TestShellsort: %v\n", data[i])
		}
	}
}

func TestHeapsort(t *testing.T) {
	data := [5]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(10, float64(i))))
		heapsort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatalf("TestHeapsort: %v\n", data[i])
		}
	}
}

func TestSearch(t *testing.T) {
	var (
		n     = 5
		data  = ints.Sorted(n)
		found bool
		index int
	)
	for i := range data {
		if index, found = Search(&data, i); i != index || !found {
			t.Fatalf("expected %d and %t, received %d and %t\n", i, true, index, found)
		}
	}

	if index, found = Search(&data, n); index != n || found {
		t.Fatalf("expected %d and %t, received %d and %t\n", 5, false, index, found)
	}

	if index, found = Search(&data, -1); index != 0 || found {
		t.Fatalf("expected %d and %t, received %d and %t\n", 0, false, index, found)
	}
}

func BenchmarkSearch(b0 *testing.B) {
	var (
		data = [5]ints.Ints{} // Careful... this is powers of ten
		n    int
	)
	for i := range data {
		n = int(math.Pow10(i))
		data[i] = ints.Sorted(n)
		b0.Run(
			fmt.Sprintf("Search for 0th element on size 10^%d", i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					Search(&data[i], 0)
				}
			},
		)
	}
}

func BenchmarkInsertionsort(b0 *testing.B) {
	var (
		data = [5]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow10(i))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("InsertionSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					insertionsort(&cpy, 0, n)
				}
			},
		)
	}
}

func BenchmarkQuicksort(b0 *testing.B) {
	var (
		data = [5]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow10(i))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("QuickSort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					quicksort(&cpy, 0, n)
				}
			},
		)
	}
}

func BenchmarkInsertionsort2(b0 *testing.B) {
	var (
		data = [16]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = i + 1
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		b0.Run(
			fmt.Sprintf("Insertionsort on size %d", n), func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					insertionsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkShellsort2(b0 *testing.B) {
	var (
		data = [16]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = i + 1
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		b0.Run(
			fmt.Sprintf("Shellsort on size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					shellsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkHeapsort2(b0 *testing.B) {
	var (
		data = [16]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = i + 1
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		b0.Run(
			fmt.Sprintf("Heapsort on size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					heapsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkQuicksort2(b0 *testing.B) {
	var (
		data = [16]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = i + 1
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		b0.Run(
			fmt.Sprintf("Quicksort on size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					quicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkGosort(b0 *testing.B) {
	var (
		data = [16]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = i + 1
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		b0.Run(
			fmt.Sprintf("Gosort on size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					sort.Ints(cpy)
				}
			},
		)
	}
}
