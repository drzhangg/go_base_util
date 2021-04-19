package 数组

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}

		// 输入: [1,3,5,6], 7
		//输出: 4
		if nums[i] > target {
			return i
		} else if nums[len(nums)-1] < target {
			return len(nums)
		}
	}
	return 0
}
