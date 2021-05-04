package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		var count int
		for j := 1; j < len(nums); j++ {
			if nums[i] > nums[j]{
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
