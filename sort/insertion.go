package sort

import (
	"github.com/homearchbishop/algorithms-golang/constraints"
)

func InsertionSort[T constraints.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		curVal := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > curVal; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = curVal
	}
	return arr
}
