package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func ShellSort[T constraints.Ordered](arr []T) []T {
	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			for j := i; j >= step && arr[j-step] > arr[j]; j -= step {
				arr[j-step], arr[j] = arr[j], arr[j-step]
			}
		}
	}
	return arr
}
