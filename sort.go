package sort

// Interface defines the contract for sorting items.
type Interface interface {
	Compare(i, j int) int
	CompareAt(i int, x interface{}) int
	Len() int
	Swap(i, j int)
}

// Sort data.
func Sort(A Interface) {
	quicksort(A, 0, A.Len()-1)
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
	var (
		r int
		a int
		b int
		c = A.Len() - 1
	)
	for a <= c {
		b = int(uint(a+c) >> 1) // (a+c)/2
		r = A.CompareAt(b, x)
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
		// if b-a < 9 {
		// 	// Insertionsort
		// 	for i := a + 1; i <= b; i++ {
		// 		for j := i - 1; 0 <= j && 0 < A.Compare(j, j+1); j-- {
		// 			A.Swap(j, j+1)
		// 		}
		// 	}
		// } else {
		// Median of three
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

		// Pivot
		p := a
		for i := a + 1; i <= b; i++ {
			if A.Compare(i, a) < 0 {
				p++
				A.Swap(i, p)
			}
		}
		A.Swap(a, p)

		quicksort(A, a, p-1)
		quicksort(A, p+1, b)
		// }
	}
}
