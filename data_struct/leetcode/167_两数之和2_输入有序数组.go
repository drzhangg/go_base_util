package leetcode

func twoSum2(numbers []int, target int) []int {
	//low := 0
	//high := len(numbers) -1
	//for low < high{
	//
	//	if numbers[low] + numbers[high] == target {
	//		return []int{low,high}
	//	}else if numbers[low] + numbers[high] > target{
	//		high = high - 1
	//	}else {
	//		low = low +1
	//	}
	//}
	//return []int{}

	for i := 0; i < len(numbers); i++ {
		left ,right := 0,len(numbers) - 1
		for left < right{
			mid := left + (right - left) / 2
			if numbers[mid] == target - numbers[i]{
				return []int{i + 1,mid + 1}
			}else if numbers[mid] > target - numbers[i] {
				right = mid -1
			}else {
				left = mid + 1
			}
		}
	}
	return []int{-1,-1}
}
