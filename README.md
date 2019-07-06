# Sort and Heap

## Sort

```GoLang
import github.com/nathangreene3/sort
```

The sort package contains functionality to sort and search data types that implement `sort.Interface`.

### Interface

* `Compare(i, j int) int`: Return -1 if the ith item is less than the jth item, 0 if equal, and 1 if greater.
* `CompareTo(i int, x interface{}) int`: Return -1 if ith item is less than x, 0 if equal, and 1 if greater.
* `Len() int`: Return the number of items.
* `Swap(i, j int)`: Interchange the ith item with the jth item.

### Functions

* `IsSorted(A Sort) bool`: Determines if sort data is sorted.
* `Search(x interface{}, A Sort) int`: Determines the index in sort data an item would be inserted into. Does not guarentee item is in the sort data set.
* `Sort(A Interface) Sort`: Sorts sort data using quicksort on large ranges and insertionsort on small ranges.
* `Stable(A Sort) Sort`: Sorts sort data using insertionsort.

## Heap

```go
import github.com/nathangreene3/sort/heap
```

The heap sub-package contains functionality to implement a minimum heap as a priority queue for data types that implement `heap.Interface`, which is an extension of `sort.Interface`.

### Interface

* `sort.Interface`: Implement the `sort.Interface` requirements.
* `Pop() interface{}`: Pop an item being the least-valued from the heap.
* `Push(x interface{})`: Push an item onto the heap.

### Functions

* `Heapify(h Interface)`: Initialize `h` as the heap.
* `Pop(h Interface) interface{}`: Pop an item as the least-valued from the heap.
* `Push(h Interface, x interface{})`: Push an item onto the heap.
