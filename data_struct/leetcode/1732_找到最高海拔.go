package leetcode

import (
	"sort"
)

func largestAltitude(gain []int) int {

	sum := 0
	arr := []int{0}
	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		arr = append(arr, sum)
	}
	sort.Ints(arr)
	return arr[len(arr)-1]
}
