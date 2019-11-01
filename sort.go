package sort

import "math"

// Sortable defines the contract for sorting items.
type Sortable interface {
	Compare(i, j int) int
	CompareTo(i int, x interface{}) int
	Len() int
	Swap(i, j int)
}

// Reverse sort data.
func Reverse(A Sortable) {
	n := A.Len() - 1
	quicksortIter(A, 0, n)
	for i := 0; i < n; i++ {
		A.Swap(i, n)
		n--
	}
}

// Sort data.
func Sort(A Sortable) {
	quicksortIter(A, 0, A.Len()-1)
}

// Stable sort data.
func Stable(A Sortable) {
	insertionsort(A, 0, A.Len()-1)
}

// IsSorted determines if data is sorted.
func IsSorted(A Sortable) bool {
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
func Search(A Sortable, x interface{}) (int, bool) {
	var a int
	for c := A.Len() - 1; a <= c; {
		b := int(uint(a+c) >> 1) // (a+c)/2
		r := A.CompareTo(b, x)
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
func insertionsort(A Sortable, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i - 1; 0 <= j && 0 < A.Compare(j, j+1); j-- {
			A.Swap(j, j+1)
		}
	}
}

// quicksort data on the range [a,b].
func quicksort(A Sortable, a, b int) {
	if a < b {
		medianOfThree(A, a, b)
		p := pivot(A, a, b)
		quicksort(A, a, p-1)
		quicksort(A, p+1, b)
	}
}

// quicksortIter on the range [a,b].
func quicksortIter(A Sortable, a, b int) {
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
				n -= 2
				a, b = stack[n], stack[n+1]
				stack = stack[:n]
			}
		}
	}
}

// quicksortTail on the range [a,b].
func quicksortTail(A Sortable, a, b int) {
	// Source: https://www.geeksforgeeks.org/quicksort-tail-call-optimization-reducing-worst-case-space-log-n/

	for a < b {
		medianOfThree(A, a, b)
		p := pivot(A, a, b)

		// One recursive call on the smaller partition
		if p-a < b-p {
			quicksortTail(A, a, p-1)
			a = p + 1
		} else {
			quicksortTail(A, p+1, b)
			b = p - 1
		}
	}
}

// medianOfThree places the median of A[a], A[(a+b)/2], and A[b] at index a, the smallest at index (a+b)/2, and the largest at index b.
func medianOfThree(A Sortable, a, b int) {
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

// pivot returns the index on the range [a,b] such that all values smaller than A
// [p] are on the range [a,p-1] and all the values larger than A[p] are on the
// range [p+1,b]. The pivot value is A[0], which will have index p when finished.
func pivot(A Sortable, a, b int) int {
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
func heapsort(A Sortable, a, b int) {
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
func siftDown(A Sortable, a, b int) {
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

// shellsort data on the range [a,b]. This is the general case for insertionsort,
// but it is not stable.
func shellsort(A Sortable, a, b int) {
	// Why do we do this? TODO: Document why.

	// Generate more gaps if needed
	gaps := []int{1, 8, 23, 77, 281} // OEIS: A036526, Sedgewick: {1, 8, 23, 77, 281, ...}
	for i := 4.0; ; i++ {
		g := int(math.Pow(4, i) + 3*math.Pow(2, i-1) + 1)
		if b < g+a {
			break
		}

		gaps = append(gaps, g)
	}

	for n := len(gaps) - 1; 0 <= n; n-- {
		g := gaps[n]
		for i := a + g; i <= b; i++ {
			for j := i - g; 0 <= j && 0 < A.Compare(j, j+g); j -= g {
				A.Swap(j, j+g)
			}
		}
	}
}
