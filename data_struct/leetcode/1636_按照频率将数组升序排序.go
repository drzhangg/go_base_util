package leetcode

import (
	"sort"
)

/*
输入：nums = [1,1,2,2,2,3]
输出：[3,1,1,2,2,2]
解释：'3' 频率为 1，'1' 频率为 2，'2' 频率为 3 。
 */
func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if m[nums[i]] == m[nums[j]] {
			return nums[i] > nums[j]
		}
		return m[nums[i]] < m[nums[j]]
	})
	return nums
}
