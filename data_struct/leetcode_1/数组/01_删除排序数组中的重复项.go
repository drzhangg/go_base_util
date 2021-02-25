package 数组

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var j = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1]{
			continue
		}else {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
