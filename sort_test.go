package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/nathangreene3/sort/ints"
)

const (
	// maxIter is the maximum value to test linear changes in size
	maxIter = 16

	// maxPow is the maximum power of two to test changes in magnitude of the size
	maxPow = 8
)

func TestFPQuicksort(t *testing.T) {
	var (
		A = ints.Reversed(256)
		B = A.FPQuicksort()
	)

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

func TestMIsort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		misort(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

func TestQuicksort(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		quicksortIter(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

func TestQuicksortIter(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		quicksortIter(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatalf("\ni = %d\nexpected sorted\nreceived: %v\n", i, data)
		}
	}
}

func TestQuicksortTail(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
		)

		quicksortTail(&data, 0, n-1)
		if !IsSorted(&data) {
			t.Fatal(data)
		}
	}
}

func TestReverse(t *testing.T) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			exp  = *data.Copy()
		)

		Reverse(&data)
		for i := 0; i < n; i++ {
			if data[i] != exp[i] {
				t.Fatalf("\nexpected %v\nreceived %v\n", exp, data)
			}
		}
	}
}

func TestSearch(t *testing.T) {
	{
		var (
			n    = 1 << maxPow
			data = ints.Random(n)
		)

		// Test each value against it's known index i
		Sort(&data) // A = [0 1 2 ... n-1]
		for i, v := range data {
			index, found := Search(&data, v)
			if i != index || !found {
				t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", i, true, index, found)
			}

			// Compare to Go's search
			if expected := sort.Search(n, func(j int) bool { return v <= data[j] }); expected != index {
				t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, true, index, found)
			}
		}

		// Test value that precedes all of data
		index, found := Search(&data, -1)
		if index != 0 || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", 0, false, index, found)
		}

		// Compare to Go's search
		if expected := sort.Search(n, func(j int) bool { return -1 <= data[j] }); expected != index || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}

		// Test value that follows all of data
		index, found = Search(&data, n)
		if index != n || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", 5, false, index, found)
		}

		// Compare to Go's search
		if expected := sort.Search(n, func(j int) bool { return n <= data[j] }); expected != index || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}
	}

	{
		var (
			n    = 1 << maxPow
			data = make(ints.Ints, 0, n)
		)

		// Test on data that is not a permutation of [0 1 2 ... n-1]
		for i := 0; i < n; i++ {
			data = append(data, rand.Intn(10)) // Each item is on range [0,10)
		}

		// Test value of -1 that precedes all of data
		var (
			expected     int
			target       = -1
			index, found = Search(&data, target)
		)

		if index != expected || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}

		// Compare to Go's search
		if expected := sort.Search(n, func(j int) bool { return target <= data[j] }); expected != index || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}

		// Test value of 10 that follows all of data
		target = 10
		expected = data.Len()
		index, found = Search(&data, target)
		if index != expected || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}

		// Compare to Go's search
		if expected := sort.Search(n, func(j int) bool { return target <= data[j] }); expected != index || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}
	}

	{
		// Test searching for value not in sorted data
		// Test value of 3 that should have index 2
		var (
			data         = ints.Ints{1, 2, 4, 5}
			expected     = 2
			target       = 3
			index, found = Search(&data, target)
		)

		if index != expected || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}

		// Compare to Go's search
		if expected := sort.Search(len(data), func(j int) bool { return target <= data[j] }); expected != index || found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, false, index, found)
		}
	}

	{
		// Test searching for value not in sorted data
		// Test value of 2 that should have index 1
		var (
			data         = ints.Ints{1, 2, 2, 2, 2, 2, 2, 2, 2, 2}
			expected     = 1
			target       = data[expected]
			index, found = Search(&data, target)
		)

		if index != expected || !found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, true, index, found)
		}

		// Compare to Go's search
		if expected := sort.Search(len(data), func(j int) bool { return target <= data[j] }); expected != index || !found {
			t.Fatalf("\nexpected (%d,%t)\nreceived (%d,%t)", expected, true, index, found)
		}
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

func BenchmarkGosort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
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

func BenchmarkHeapsort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
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

func BenchmarkInsertionsort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					insertionsort(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkMIsort(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					misort(&cpy, 0, n-1)
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

func BenchmarkQuicksortIter(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					quicksortIter(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkQuicksortTail(b0 *testing.B) {
	for i := 0; i < maxPow; i++ {
		var (
			n    = 1 << uint(i)
			data = ints.Reversed(n)
			cpy  = ints.New(n, n)
		)

		b0.Run(
			fmt.Sprintf("size %d", n),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data)
					quicksortTail(&cpy, 0, n-1)
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

func BenchmarkQuicksortIter2(b0 *testing.B) {
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
					quicksortIter(&cpy, 0, n-1)
				}
			},
		)
	}
}

func BenchmarkQuicksortTail2(b0 *testing.B) {
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
					quicksortTail(&cpy, 0, n-1)
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

// ------------------------------------------------------------------
// Benchmark ...
// ------------------------------------------------------------------
