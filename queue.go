package main

// import "fmt"

// type queue struct {
// 	count    int      // Number of items in the queue
// 	capacity int      // Maximum number of items able to be held in the queue
// 	heap     sortable // Ordered list of sortable items
// }

// func newQueue(capacity int) (q queue) {
// 	if 0 < capacity {
// 		q.count = 0
// 		q.capacity = capacity
// 		q.heap = make(sortable, capacity+1)
// 	}
// 	return q
// }

// func (q queue) enqueue(item sortable) {
// 	if q.count < q.capacity {
// 		q.count++
// 		q.heap[q.count] = item
// 		q.siftUp(q.count)
// 	}
// }

// func (q queue) dequeue() (item sortable) {
// 	if 0 < q.count {
// 		item = q.heap[1]
// 		q.heap[1] = q.heap[q.count]
// 		q.count--
// 		q.siftDown(1)
// 	}
// 	return item
// }

// func (q queue) swap(i, j int) {
// 	q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
// }

// func (q queue) siftUp(i int) {
// 	parent := i
// 	for 1 < i && i <= q.count {
// 		parent /= 2
// 		if q.heap[parent].less(q.heap[i]) {
// 			q.swap(parent, i)
// 			i = parent
// 		} else {
// 			break
// 		}
// 	}
// }

// func (q queue) siftDown(i int) {
// 	left := i
// 	right := i
// 	for 0 < i && i < q.count/2 {
// 		left = 2 * i
// 		right = 2*i + 1
// 		if q.heap[left].less(q.heap[right]) {
// 			if q.heap[i].less(q.heap[right]) {
// 				q.swap(i, right)
// 				i = right
// 			} else {
// 				break
// 			}
// 		} else if q.heap[i].less(q.heap[left]) {
// 			q.swap(i, left)
// 			i = left
// 		} else {
// 			break
// 		}
// 	}
// }

// func (q queue) printHeapln() {
// 	for i := 0; i < q.count; i++ {
// 		fmt.Print(q.heap[i])
// 	}
// 	fmt.Println()
// }
