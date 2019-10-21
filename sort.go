package sort

import (
	"math"
)

// Interface defines the contract for sorting items.
type Interface interface {
	Compare(i, j int) int
	CompareAt(i int, x interface{}) int
	Len() int
	Swap(i, j int)
}

// Sort data.
func Sort(A Interface) {
	iterativeQuicksort(A, 0, A.Len()-1)
}

// Stable sort data.
func Stable(A Interface) {
	insertionsort(A, 0, A.Len()-1)
}

// IsSorted determines if data is sorted.
func IsSorted(A Interface) bool {
	n := A.Len() - 1
	for i := 0; i < n; i++ {
		if 0 < A.Compare(i, i+1) {
			return false
		}
	}

	return true
}

// Search returns the index an item belongs in a list of items and
// whether or not it was found.
func Search(A Interface, x interface{}) (int, bool) {
	var a int
	for c := A.Len() - 1; a <= c; {
		b := int(uint(a+c) >> 1) // (a+c)/2
		r := A.CompareAt(b, x)
		switch {
		case r < 0:
			a = b + 1
		case 0 < r:
			c = b - 1
		default:
			return b, true
		}
	}

	return a, false
}

// insertionsort data on the range [a,b].
func insertionsort(A Interface, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.Compare(j, j+1); j-- {
			A.Swap(j, j+1)
		}
	}
}

// quicksort data on the range [a,b].
func quicksort(A Interface, a, b int) {
	if a < b {
		if a < b+9 {
			insertionsort(A, a, b)
			// TODO
			// } else if a < b+1<<4 {
			// 	heapsort(A, a, b)
		} else {
			medianOfThree(A, a, b)
			p := pivot(A, a, b)
			quicksort(A, a, p-1)
			quicksort(A, p+1, b)
		}
	}
}

// iterativeQuicksort on the range [a,b].
func iterativeQuicksort(A Interface, a, b int) {
	if a < b {
		stack := append(make([]int, 0, b-a+1), a, b)
		for n := 2; 0 < n; { // n is stack len
			if a < b {
				medianOfThree(A, a, b)
				if p := pivot(A, a, b); b-p < p-a {
					stack = append(stack, a, p-1)
					a = p + 1
				} else {
					stack = append(stack, p+1, b)
					b = p - 1
				}

				n += 2
			} else {
				n--
				b = stack[n]
				stack = stack[:n]

				n--
				a = stack[n]
				stack = stack[:n]
			}
		}
	}
}

func medianOfThree(A Interface, a, b int) {
	c := int(uint(a+b) >> 1) // (a+b)/2
	if A.Compare(a, b) < 0 {
		A.Swap(a, b)
	}

	if A.Compare(c, a) < 0 {
		A.Swap(a, c)
	}

	if A.Compare(b, c) < 0 {
		A.Swap(b, c)
	}
}

func pivot(A Interface, a, b int) int {
	p := a
	for i := a + 1; i <= b; i++ {
		if A.Compare(i, a) < 0 {
			p++
			A.Swap(i, p)
		}
	}

	A.Swap(a, p)
	return p
}

// heapsort data on the range [a,b].
func heapsort(A Interface, a, b int) {
	// Heapify A into a max heap
	for i := b >> 1; 0 <= i; i-- {
		siftDown(A, i, b)
	}

	// Pop b-a times
	for a < b {
		A.Swap(a, b)
		b--
		siftDown(A, a, b)
	}
}

// siftDown corrects the heap on the index range [a,b].
func siftDown(A Interface, a, b int) {
	// j is left child, k is right child
	if j := int(uint(a)<<1) + 1; j <= b {
		k := j + 1
		if k <= b && A.Compare(j, k) < 0 {
			j = k
		}

		if A.Compare(a, j) < 0 {
			A.Swap(a, j)
			siftDown(A, j, b)
		}
	}
}

// shellsort data on the range [a,b]. This is the general case for insertionsort, but it is not stable.
func shellsort(A Interface, a, b int) {
	// Generate more gaps if needed
	var (
		gaps = []int{1, 8, 23, 77, 281} // OEIS: A036526, Sedgewick: {1, 8, 23, 77, 281, ...}
		g    int
	)
	for i := 4.0; ; i++ {
		g = int(math.Pow(4, i) + 3*math.Pow(2, i-1) + 1)
		if b < g+a {
			break
		}

		gaps = append(gaps, g)
	}

	for n := len(gaps) - 1; 0 <= n; n-- {
		g = gaps[n]
		for i := a + g; i <= b; i++ {
			for j := i - g; 0 <= j && 0 < A.Compare(j, j+g); j -= g {
				A.Swap(j, j+g)
			}
		}
	}
}
