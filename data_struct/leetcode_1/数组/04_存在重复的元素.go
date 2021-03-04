package 数组

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	set := make(map[int]struct{})
	for _,v := range nums{
		if _, has := set[v];has{
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
