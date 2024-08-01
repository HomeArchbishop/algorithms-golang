package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func CombSort[T constraints.Ordered](arr []T) []T {
	gap := len(arr)
	shrink := 0.8
	swapped := true
	for swapped || gap > 1 {
		if gap > 1 {
			gap = int(float64(gap) * shrink)
		}
		swapped = false
		for i := 0; i < len(arr)-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
