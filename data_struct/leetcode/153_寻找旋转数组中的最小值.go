package leetcode

func findMin1(nums []int) int {
	n := len(nums)

	left,right := 0,n -1
	for left < right{
		mid := left + (right - left) / 2
		if nums[mid] > nums[right]{
			left = mid + 1
		}else if nums[mid] < nums[right]{
			right = mid
		}else {
			return nums[mid]
		}
	}
	return -1
}
