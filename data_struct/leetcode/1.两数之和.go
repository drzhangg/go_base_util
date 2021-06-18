package leetcode

func twoSum(nums []int, target int) []int {
	var result []int
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+ nums[j] == target{
	//			result = append(result, i,j)
	//		}
	//	}
	//}

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok && m[target-nums[i]] != i{
			result = append(result,m[target-nums[i]],i)
			break
		}
	}
	return result
}
