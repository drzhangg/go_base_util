package leetcode

func minOperations(nums []int) int {
	/*
		1.先找到最大值，和最大值的下标
		2.如果下标大于最大值，值小于最大值，要加i-max[index]个数


	*/
	var result int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i - 1]{
			result += nums[i-1] -nums[i] + 1
			nums[i] = nums[i -1 ]+1
		}
	}
	return result

}
