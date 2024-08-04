package sort

import "github.com/homearchbishop/algorithms-golang/constraints"

func PatienceSort[T constraints.Ordered](arr []T) []T {
	piles := [][]T{}

	for i := 0; i < len(arr); i++ {
		l, r := 0, len(piles)
		for l < r {
			mid := (l + r) / 2
			if piles[mid][len(piles[mid])-1] >= arr[i] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l == len(piles) {
			piles = append(piles, []T{arr[i]})
		} else {
			piles[l] = append(piles[l], arr[i])
		}
	}

	for idx := 0; idx < len(arr); idx++ {
		minPileIdx := 0
		for j := 1; j < len(piles); j++ {
			if piles[j][len(piles[j])-1] < piles[minPileIdx][len(piles[minPileIdx])-1] {
				minPileIdx = j
			}
		}
		arr[idx] = piles[minPileIdx][len(piles[minPileIdx])-1]
		piles[minPileIdx] = piles[minPileIdx][:len(piles[minPileIdx])-1]
		if len(piles[minPileIdx]) == 0 {
			piles = append(piles[0:minPileIdx], piles[minPileIdx+1:]...)
		}
	}
	return arr
}
