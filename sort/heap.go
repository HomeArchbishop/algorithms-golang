package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

// We will *not* change the length and capacity of the slice,
// so we simply use []T instead of *[]T
func heapify[T constraints.Ordered](arr []T, i int, end int) {
	l := i*2 + 1
	if l > end {
		return
	}
	r := i*2 + 2
	maxChdIdx := l
	if r <= end && arr[r] > arr[l] {
		maxChdIdx = r
	}
	if arr[maxChdIdx] > arr[i] {
		arr[maxChdIdx], arr[i] = arr[i], arr[maxChdIdx]
		heapify(arr, maxChdIdx, end)
	}
}

func HeapSort[T constraints.Ordered](arr []T) []T {
	for i := len(arr)/2 - 1; i > -1; i-- {
		heapify(arr, i, len(arr)-1)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i-1)
	}
	return arr
}
