package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func quickSort[T constraints.Ordered](arr *[]T, left int, right int, pivotIndex int) {
	if left >= right {
		return
	}
	newPivotIndex := left
	pivotValue := (*arr)[pivotIndex]
	(*arr)[pivotIndex], (*arr)[right] = (*arr)[right], (*arr)[pivotIndex]
	for i := left; i < right; i++ {
		if (*arr)[i] < pivotValue {
			(*arr)[i], (*arr)[newPivotIndex] = (*arr)[newPivotIndex], (*arr)[i]
			newPivotIndex++
		}
	}
	(*arr)[right], (*arr)[newPivotIndex] = (*arr)[newPivotIndex], (*arr)[right]
	quickSort(arr, left, newPivotIndex-1, (left+newPivotIndex-1)/2)
	quickSort(arr, newPivotIndex+1, right, (newPivotIndex+1+right)/2)
}

func QuickSort[T constraints.Ordered](arr []T) []T {
	quickSort(&arr, 0, len(arr)-1, len(arr)/2)
	return arr
}
