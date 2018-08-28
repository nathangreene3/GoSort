package main

import (
	"fmt"
)

func main() {
	A := randomSlice(10)
	// B := copySlice(A)
	// C := copySlice(A)
	D := intSlice(copySlice(A))
	// fmt.Println("A before mergeSort:", A)
	// fmt.Println("A after mergeSort: ", mergeSort(A, 0, len(A)-1))
	// fmt.Println("B before quickSort:", B)
	// fmt.Println("B after quickSort: ", quickSort(B, 0, len(B)-1))
	// fmt.Println("C before insertionSort:", C)
	// fmt.Println("C after insertionSort: ", insertionSort(C))

	// n := 10
	// alphabet := "abcdefghijklmnopqrstuvwxyz"
	// ppl := make(people, n) // or people{}
	// for i := 0; i < n; i++ {
	// 	ppl[i] = person{
	// 		first: charAt(rand.Intn(len(alphabet)), alphabet),
	// 		last:  charAt(rand.Intn(len(alphabet)), alphabet),
	// 	}
	// }
	// fmt.Println(bubbleSortable(ppl))
	fmt.Println(bubbleSortable(D))
	q := newQueue(5)
	for i := 0; i < 5; i++ {
		q.enqueue(comparableInt(i))
	}
}
