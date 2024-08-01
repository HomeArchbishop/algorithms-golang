package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func CocktailSort[T constraints.Ordered](arr []T) []T {
	swapped := true
	start := 0
	end := len(arr) - 1
	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		end--
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		start++
	}
	return arr
}
