package heap

import (
	"fmt"
	"math"
	"testing"

	"github.com/nathangreene3/sort"
	"github.com/nathangreene3/sort/ints"
)

func heapsort(A ints.Ints) ints.Ints {
	Heapify(&A)
	n := (&A).Len()
	B := ints.New(0, n)
	for i := 0; i < n; i++ {
		(&B).Push(Pop(&A))
	}

	return B
}

func TestHeap(t *testing.T) {
	var (
		data       = ints.Random(10)
		sortedData = heapsort(data)
	)
	if !sort.IsSorted(&sortedData) {
		t.Fatalf("expected sorted, received %v", sortedData)
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
		n--
		b0.Run(
			fmt.Sprintf("Heapsort on size 10^%d", i+1),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					copy(cpy, data[i])
					heapsort(cpy)
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
					heapsort(cpy)
				}
			},
		)
	}
}
