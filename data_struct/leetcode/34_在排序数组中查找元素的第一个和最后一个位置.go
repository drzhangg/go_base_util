package leetcode

func searchRange(nums []int, target int) []int {
	left,right := 0,len(nums)
	for left < right{
		mid := left + (right - left)/2
		if nums[mid] == target{
			l,r := mid,mid
			//找左边
			for l > 0 && nums[l] == nums[l-1]{
				l--
			}
			for r+1 < len(nums)&& nums[r] == nums[r + 1]{
				r++
			}
			return []int{l,r}
		}else if nums[mid] > target {
			right = mid
		}else {
			left = mid +1
		}
	}
	return []int{-1,-1}

	//官方解法
	//left := sort.SearchInts(nums,target)
	//if left == len(nums) || nums[left] != target{
	//	return []int{-1,-1}
	//}
	//
	//right := sort.SearchInts(nums,target+1) - 1
	//return []int{left,right}
}
