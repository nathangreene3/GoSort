package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	bts := []byte("fdsa")
	n := len(bts)
	perm := newPermutation(n)
	fmt.Println(bts, perm)

	perm = keySort(bts, nil, 0, n-1)
	fmt.Println(bts, perm)

	perm = keySort(bts, perm, 0, n-1)
	fmt.Println(bts, perm)
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
