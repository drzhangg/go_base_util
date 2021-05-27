package leetcode

func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1{
		if nums[0] == target{
			return 0
		}
		return -1
	}

	left,right := 0, len(nums) - 1
	for left <= right{
		mid := left + (right - left) /2
		if nums[mid] == target{
			return mid
		}

		if nums[0] <= nums[mid]{//说明左边是有序排列的
			if nums[0] <= target && target < nums[mid]{//边界条件
				right = mid - 1
			}else {
				left = mid + 1
			}
		}else {//说明右边是有序排列的
			if target > nums[mid] && target <= nums[len(nums)-1]{//边界条件
				left = mid + 1
			}else {
				right = mid - 1
			}
		}
	}
	return -1
}
