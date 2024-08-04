package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func merge[T constraints.Ordered](arr []T, start1 int, end1 int, start2 int, end2 int) {
	arr1, arr2 := make([]T, end1-start1), make([]T, end2-start2)
	copy(arr1, arr[start1:end1])
	copy(arr2, arr[start2:end2])
	i, j := 0, 0
	for i < end1-start1 && j < end2-start2 {
		if arr1[i] < arr2[j] {
			arr[i+j+start1] = arr1[i]
			i++
			continue
		}
		if arr1[i] > arr2[j] {
			arr[i+j+start1] = arr2[j]
			j++
			continue
		}
		if arr1[i] == arr2[j] {
			arr[i+j+start1] = arr1[i]
			i++
			arr[i+j+start1] = arr2[j]
			j++
		}
	}
	for i < end1-start1 {
		arr[i+j+start1] = arr1[i]
		i++
	}
	for j < end2-start2 {
		arr[i+j+start1] = arr2[j]
		j++
	}
}

func MergeSort[T constraints.Ordered](arr []T) []T {
	step := 1
	for step < len(arr) {
		for i := 0; i+step < len(arr); i += step * 2 {
			start1 := i
			end1 := i + step
			start2 := end1
			end2 := min(i+step*2, len(arr))
			merge(arr, start1, end1, start2, end2)
		}
		step *= 2
	}
	return arr
}
