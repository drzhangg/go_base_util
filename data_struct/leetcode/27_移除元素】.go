package leetcode

/*
双指针解法，
开始的时候快慢指针都指向数组的头部
当遍历的指针对应的数值不等于target值时，将慢指针指向新的下标位置
等于是跳过了target目标值对应的下标
 */
func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val{
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
