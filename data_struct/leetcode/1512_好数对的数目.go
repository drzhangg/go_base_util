package leetcode

import "fmt"

func numIdenticalPairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[i] == nums[j] && i < j {
				fmt.Println(i, j)
				count++
			}
		}
	}
	return count
}
