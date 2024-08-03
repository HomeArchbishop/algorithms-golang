package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func SelectionSort[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	return arr
}
