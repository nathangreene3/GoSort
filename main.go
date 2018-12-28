package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	// sort.Sort(nil)
	// A := people{
	// 	&person{"", "a"},
	// 	&person{"", "b"},
	// 	&person{"", "c"},
	// 	&person{"", "d"},
	// 	&person{"", "e"},
	// 	&person{"", "f"},
	// 	&person{"", "g"},
	// 	&person{"", "h"},
	// 	&person{"", "i"},
	// 	&person{"", "j"},
	// }
	// A.randomize()
	// fmt.Println(A.String())
	// quickSortable(A, 0, A.length()-1)
	// fmt.Printf("%v\n", A)

	// A := randomIntSlice(10)
	// fmt.Println(A)
	// quickSortable(A, 0, A.length()-1)
	// fmt.Println(A)

	A := newSlice(10)
	var B []int
	var e, max float64
	for {
		A = nextSlice(A)
		e = avgIndexError(A)
		if max < e {
			max = e
			B = copySlice(A)
		}
		if isSorted(A) {
			break
		}
	}
	fmt.Println(B, max)
}
