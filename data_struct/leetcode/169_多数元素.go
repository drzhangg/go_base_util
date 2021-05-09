package leetcode

import "fmt"

func majorityElement(nums []int) int {
	var hasMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, has := hasMap[nums[i]]; has {
			hasMap[nums[i]]++
		} else {
			hasMap[nums[i]] = 1
		}
	}

	var result,count int
	for k,v := range hasMap{
		if count < v{
			count = v
			result = k
		}
		fmt.Println(k,v)
	}

	return result
}
