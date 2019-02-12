package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	// bts := []byte("1230")
	// n := len(bts)

	// fmt.Println(indexMap(bts, 0, n-1))
	// fmt.Println(incrementPositions(bts))
	// fmt.Println(getIncrementPositions(string(bts)))
	fmt.Println(perm(10))
	fmt.Println(randomPermBts(10))
}

// keySort sorts a byte slice and returns a permutation indicating the index
// changes for each character in the given byte slice. Runs in O(n) at best and
// O(n^2) at worst.
//
// For example, [b c a d e] would be sorted as [a b c d e] and [2 0 1 3 4] would
// be returned. The a moved from index 2 to index 0, b moved from index 0 to 1,
// the c moved from index 1 to index 2, and the d and e did not move.
func keySort(bts []byte, indexMap []int, a, b int) []int {
	if indexMap == nil {
		indexMap = make([]int, 0, len(bts)) // A permutation, that is, an ordering of [0, 1, 2, ..., n-1]
		for i := range bts {
			indexMap = append(indexMap, i)
		}

		// insertionsort bts and update indexMap at the same time
		// TODO: use quicksort on larger data
		for i := a + 1; i <= b; i++ {
			for j := i - 1; a <= j && bts[j+1] < bts[j]; j-- {
				bts[j], bts[j+1] = bts[j+1], bts[j]
				indexMap[j], indexMap[j+1] = indexMap[j+1], indexMap[j]
			}
		}

		return indexMap
	}

	for i := range bts {
		bts[i], bts[indexMap[i]] = bts[indexMap[i]], bts[i]
		indexMap[i], indexMap[indexMap[i]] = indexMap[indexMap[i]], indexMap[i]
	}

	return indexMap
}

// indexMap returns a permutation indicating the sort order of the indices of the
// given byte slice over the range [a,b], where 0 <= a <= b < len(bts).
func indexMap(bts []byte, a, b int) []int {
	m := make([]int, 0, len(bts)) // An ordering (permutation) of [0, 1, 2, ..., n-1]
	for i := range bts {
		m = append(m, i)
	}

	// Insertionsort on m, not bts
	var j int    // Indexer
	var mi int   // Temporary storage; mi is always a copy of m[i]
	var bte byte // Temporary storage; bte is always a copy of bts[m[i]]

	for i := a + 1; i <= b; i++ {
		mi = m[i]
		bte = bts[mi]

		for j = i - 1; a <= j && bte < bts[m[j]]; j-- {
			// m[j+1] = m[j]
		}

		// m[j+1] = mi
		copy(m[j+2:i+1], m[j+1:i])
		m[j+1] = mi
	}

	return m
}

func newPermutation(n int) []int {
	perm := make([]int, 0, n)
	for i := 0; i < n; i++ {
		perm = append(perm, i)
	}
	return perm
}

func incrementPositions(incrementOrder []byte) []int {
	var pos []int
	for i := 0; i <= len(incrementOrder); i++ {
		for j, k := range incrementOrder {
			if int(k) == i+48 {
				pos = append(pos, j)
				break
			}
		}
	}
	return pos
}

func getIncrementPositions(incrementOrder string) string {
	var indexOf int
	sb := strings.Builder{}
	sb.Grow(8)

	for i := 1; i < 9; i++ {
		indexOf = strings.Index(incrementOrder, strconv.Itoa(i))
		if 0 <= indexOf {
			sb.WriteString(strconv.Itoa(indexOf))
		}
	}

	return sb.String()
}

func randomBts(n int) []byte {
	bts := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		bts = append(bts, byte(rand.Intn(10)))
	}
	return bts
}

// TODO: make permutation type with full functionality.  See GoTSP usage.
func perm(n int) []int {
	if n < 1 {
		return []int{}
	}

	p := make([]int, n) // First value is always zero
	var j int

	for i := 1; i < n; i++ {
		j = rand.Intn(i + 1)
		p[i] = p[j]
		p[j] = i
	}

	return p
}

// randomPermBts returns a random permutation of a given length. Panics if length
// is greater than 255. Returns an empty byte slice if n is less than one.
func randomPermBts(n int) []byte {
	if 255 < n {
		panic("maximum random permutation length is 255")
	}

	if n < 1 {
		return []byte{}
	}

	bts := make([]byte, n)
	var j int

	for i := 1; i < n; i++ {
		j = rand.Intn(i + 1)
		bts[i] = bts[j]
		bts[j] = byte(i)
	}

	return bts
}
