package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func GnomeSort[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr)-1; {
		if i == -1 || arr[i] <= arr[i+1] {
			i++
		} else {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i--
		}
	}
	return arr
}
