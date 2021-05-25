package leetcode

func search(nums []int, target int) int {
	//左闭右闭
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] > target {
			right = mid - 1
		}else if nums[mid] < target{
			left = mid + 1
		}else {
			return mid
		}
	}
	return -1
}
