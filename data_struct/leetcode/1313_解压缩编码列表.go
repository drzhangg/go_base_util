package leetcode

func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i += 2 {
		for j := 1; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
