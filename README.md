# Sort and Heap

## Sort

```GoLang
import github.com/nathangreene3/sort
```

The sort package contains functionality to sort and search data types that implement `sort.Interface`.

### Interface

* `Compare(i, j int) int:` Return -1 if the ith datum is less than the jth datum, 0 if equal, and 1 if greater.
* `CompareTo(i int, x interface{}) int:` Return -1 if ith datum is less than x, 0 if equal, and 1 if greater.
* `Len() int:` Return the number of datums.
* `Swap(i, j int):` Interchange the ith datum with the jth datum.

### Functions

* `Sort(A Interface) Sort:` Sorts sort data using quicksort on large ranges and insertionsort on small ranges.
* `Stable(A Sort) Sort:` Sorts sort data using insertionsort.
* `IsSorted(A Sort) bool:` Determines if sort data is sorted.
* `Search(x interface{}, A Sort) int:` Determines the index in sort data an datum would be inserted into. Does not guarentee datum is in the sort data set.

## Heap

```go
import github.com/nathangreene3/sort/heap
```

The heap sub-package contains functionality to implement a minimum heap as a priority queue for data types that implement `heap.Interface`, which is an extension of `sort.Interface`.

### Interface

* `sort.Interface:` Implement the `sort.Interface` requirements.
* `Push(x interface{}):` Push a datum onto the heap.
* `Pop() interface{}:` Pop a datum being the least-valued from the heap.

### Functions

* `Heapify(h Interface):` Initialize `h` as the heap.
* `Push(h Interface, x interface{}):` Push a datum onto the heap.
* `Pop(h Interface) interface{}:` Pop a datum as the least-valued from the heap.
