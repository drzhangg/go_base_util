package leetcode

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		fmt.Println(arr)
		fmt.Println(arr[i],arr[j])
		a, b := getCount(arr[i]), getCount(arr[j])
		fmt.Println(a,b)
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

func getCount(n int) int {
	var count int
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}
