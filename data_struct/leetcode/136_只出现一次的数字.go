package leetcode

func singleNumber(nums []int) int {
	var ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
