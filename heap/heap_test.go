package heap

import (
	"fmt"
	"math"
	"testing"

	"github.com/nathangreene3/sort"
	"github.com/nathangreene3/sort/ints"
)

func TestHeap(t *testing.T) {
	data := ints.Random(10)
	Sort(&data)
	if !sort.IsSorted(&data) {
		t.Fatalf("expected sorted, received %v", data)
	}
}

func BenchmarkHeapsort(b0 *testing.B) {
	var (
		data = [5]ints.Ints{}
		cpy  ints.Ints
		n    int
	)
	for i := range data {
		n = int(math.Pow10(i))
		data[i] = ints.Reversed(n)
		cpy = ints.New(n, n)
		b0.Run(
			fmt.Sprintf("Heapsort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					Sort(&cpy)
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
					Sort(&cpy)
				}
			},
		)
	}
}
