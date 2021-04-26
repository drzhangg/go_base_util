package leetcode

func arraySign(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var x int
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result++
			break
		} else if nums[i] < 0 {
			x++
		}
	}

	if result > 0 {
		return 0
	}
	if x%2 != 0 {
		return -1
	}
	return 1
}
