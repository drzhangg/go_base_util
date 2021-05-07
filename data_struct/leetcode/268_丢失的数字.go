package leetcode

func missingNumber(nums []int) int {
	var result int
	for i, k := range nums {
		result ^= k ^ i
	}
	return result ^ len(nums)
}

