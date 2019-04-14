# Sort

```GoLang
import github.com/nathangreene3/sort
```

## Description

The sort package contains functionality to sort and search data types that implement the sortable interface.

## Sortable interface

* **Compare(i, j int) int:** Return -1 if the ith datum is less than the jth datum, 0 if equal, and 1 if greater.
* **CompareTo(x interface{}, i int) int:** Return -1 if x is less than the ith datum, 0 if equal, and 1 if greater.
* **Length() int:** Return the number of datums.
* **Swap(i, j int):** Interchange the ith datum with the jth datum.

## Functions

* **Sort(A Sortable) Sortable:** Sorts sortable data using quicksort on large ranges and insertionsort on small ranges.
* **Stable(A Sortable) Sortable:** Sorts sortable data using insertionsort.
* **IsSorted(A Sortable) bool:** Determines if sortable data is sorted.
* **Search(x interface{}, A Sortable) int:** Determines the index in sortable data an datum would be inserted into. Does not guarentee datum is in the sortable data set.
