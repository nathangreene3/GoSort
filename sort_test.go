package sort

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/nathangreene3/sort/ints"
)

func TestInsertionsort(t *testing.T) {
	data := [8]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(2, float64(i))))
		insertionsort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatal(data[i])
		}
	}
}

func TestShellsort(t *testing.T) {
	data := [8]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(2, float64(i))))
		shellsort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatal(data[i])
		}
	}
}

func TestHeapsort(t *testing.T) {
	data := [8]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(2, float64(i))))
		heapsort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatal(data[i])
		}
	}
}

func TestQuicksort(t *testing.T) {
	data := [8]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(2, float64(i))))
		quicksort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatal(data[i])
		}
	}
}

func TestIterativeQuicksort(t *testing.T) {
	data := [8]ints.Ints{}
	for i := range data {
		data[i] = ints.Reversed(int(math.Pow(2, float64(i))))
		iterativeQuicksort(&data[i], 0, len(data[i])-1)
		if !IsSorted(&data[i]) {
			t.Fatal(data[i])
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
			t.Fatalf("expected (%d,%t), received (%d,%t)", i, true, index, found)
		}
	}

	if index, found = Search(&data, n); index != n || found {
		t.Fatalf("expected (%d,%t), received (%d,%t)n", 5, false, index, found)
	}

	if index, found = Search(&data, -1); index != 0 || found {
		t.Fatalf("expected (%d,%t), received (%d,%t)", 0, false, index, found)
	}
}

func TestFPQuicksort(t *testing.T) {
	n := 10
	A := ints.Reversed(n)
	B := A.FPQuicksort()
	if !IsSorted(&B) {
		t.Fatalf("expected sorted, received %v", B)
	}
}

func BenchmarkSearch(b0 *testing.B) {
	var (
		data = [8]ints.Ints{} // Careful... this is powers of ten
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Sorted(n)
		b0.Run(
			fmt.Sprintf("0th element, size 2^%d", i),
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
		data = [8]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
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
		data = [8]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					quicksort(&cpy, 0, n)
				}
			},
		)
	}
}

func BenchmarkIterativeQuicksort(b0 *testing.B) {
	var (
		data = [8]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					iterativeQuicksort(&cpy, 0, n)
				}
			},
		)
	}
}

func BenchmarkHeapsort(b0 *testing.B) {
	var (
		data = [8]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					heapsort(&cpy, 0, n)
				}
			},
		)
	}
}

func BenchmarkGosort(b0 *testing.B) {
	var (
		data = [8]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					sort.Ints(cpy)
				}
			},
		)
	}
}

func BenchmarkFPQuicksort(b0 *testing.B) {
	var (
		data = [8]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow(2, float64(i)))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		n--
		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					cpy.FPQuicksort()
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
			fmt.Sprintf("size %d", n), func(b1 *testing.B) {
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
			fmt.Sprintf("size %d", n),
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
			fmt.Sprintf("size %d", n),
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
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					quicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkIterativeQuicksort2(b0 *testing.B) {
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
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					iterativeQuicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkGosort2(b0 *testing.B) {
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
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					sort.Ints(cpy)
				}
			},
		)
	}
}

func BenchmarkFPQuicksort2(b0 *testing.B) {
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
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					cpy.FPQuicksort()
				}
			},
		)
	}
}
