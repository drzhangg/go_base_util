package leetcode

/*
双指针解法，
开始的时候快慢指针都指向数组的头部

 */
func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val{
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow
}
