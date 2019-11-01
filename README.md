# Sort and Heap

## Sort

```GoLang
import github.com/nathangreene3/sort
```

The sort package contains functionality to sort and search data types that implement `sort.Sortable`.

### Sortable

```go
Compare(i, j int) int               // Return -1 if the ith item is less than the jth item, 0 if equal, and 1 if greater.
CompareTo(i int, x interface{}) int // Return -1 if ith item is less than x, 0 if equal, and 1 if greater.
Len() int                           // Return the number of items.
Swap(i, j int)                      // Interchange the ith item with the jth item.
```

### Functions

```go
IsSorted(A Sortable) bool             // Determines if sortable data is sorted.
Reverse(A Sortable)                   // Reverse-sorts sortable data.
Search(x interface{}, A Sortable) int // Determines the index in sortable data an item would be inserted into. Does not guarentee item is in the sortable data set.
Sort(A Sortable) Sortable             // Sorts sortable data using quicksort on large ranges and insertionsort on small ranges.
Stable(A Sortable) Sortable           // Sorts sortable data using insertionsort.
```

## Heap

```go
import github.com/nathangreene3/sort/heap
```

The heap sub-package contains functionality to implement a minimum heap as a priority queue for data types that implement `heap.Heapable`, which is an extension of `sort.Sortable`.

### Heapable

```go
sort.Sortable      // Implement the sort.Interface requirements.
Pop() interface{}   // Pop an item being the least-valued from the heap.
Push(x interface{}) // Push an item onto the heap.
```

### Functions

```go
Heapify(h Heapable)              // Initialize h as the heap.
Pop(h Heapable) interface{}      // Pop an item as the least-valued from the heap.
Push(h Interface, x interface{}) // Push an item onto the heap.
```
