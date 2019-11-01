package heap

import (
	"fmt"
	"testing"

	"github.com/nathangreene3/sort"
	"github.com/nathangreene3/sort/ints"
)

const (
	// maxIter is the maximum value to test linear changes in size
	maxIter = 16

	// maxPow is the maximum power of two to test changes in magnitude of the size
	maxPow = 8
)

func TestHeap(t *testing.T) {
	data := ints.Random(256)
	Sort(&data)
	if !sort.IsSorted(&data) {
		t.Fatalf("expected sorted, received %v", data)
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
					Sort(&cpy)
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
					Sort(&cpy)
				}
			},
		)
	}
}
