package sort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/nathangreene3/sort/ints"
)

const (
	maxPow  = 8
	maxIter = 16
)

func TestFPQuicksort(t *testing.T) {
	n := 10
	A := ints.Reversed(n)
	B := A.FPQuicksort()
	if !IsSorted(&B) {
		t.Fatalf("expected sorted, received %v", B)
	}
}

func TestHeapsort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		heapsort(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

func TestInsertionsort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		insertionsort(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

func TestIterativeQuicksort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		iterQuicksort(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatalf("\ni = %d\nexpected sorted\nreceived: %v\n", i, data)
		}
	}
}

func TestQuicksort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		quicksort(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

func TestSearch(t *testing.T) {
	var (
		n    = 1 << maxPow
		data = ints.Sorted(n)
	)

	for i := range data {
		if index, found := Search(&data, i); i != index || !found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", i, true, index, found)
		}
	}

	if index, found := Search(&data, n); index != n || found {
		t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", 5, false, index, found)
	}

	if index, found := Search(&data, -1); index != 0 || found {
		t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", 0, false, index, found)
	}
}

func TestShellsort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		shellsort(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

// ------------------------------------------------------------------
// Benchmark order-increasing sizes 2^k, for k in [0,8)
// ------------------------------------------------------------------

func BenchmarkFPQuicksort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					cpy.FPQuicksort()
				}
			},
		)
	}
}

func BenchmarkGosort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					sort.Ints(cpy)
				}
			},
		)
	}
}

func BenchmarkHeapsort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					heapsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkInsertionsort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					insertionsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkIterativeQuicksort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					iterQuicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkQuicksort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					quicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkShellsort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size 2^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					shellsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkSearch(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		data := ints.Sorted(1 << uint(i))
		b0.Run(
			fmt.Sprintf("0th element, size 2^%d", i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					Search(&data, 0)
				}
			},
		)
	}
}

// ------------------------------------------------------------------
// Benchmark linear-increasing sizes n in [1,16]
// ------------------------------------------------------------------

func BenchmarkFPQuicksort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					cpy.FPQuicksort()
				}
			},
		)
	}
}

func BenchmarkGosort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					sort.Ints(cpy)
				}
			},
		)
	}
}

func BenchmarkHeapsort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					heapsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkInsertionsort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n), func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					insertionsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkIterativeQuicksort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					iterQuicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkQuicksort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					quicksort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkShellsort2(b0 *testing.B) {
	for i := 0; i < maxIter; i++ {
		var (
			n    = i + 1
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					shellsort(&cpy, 0, n-1)
				}
			},
		)
	}
}
