package main

import (
	"fmt"
	"math/rand"
	srt "sort"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	srt.Sort(nil)
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

	A := randomIntSlice(10)
	fmt.Println(A)
	quickSortable(A, 0, A.length()-1)
	fmt.Println(A)
}
