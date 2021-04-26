package leetcode

func runningSum(nums []int) []int {

	//我的写法
	var result int
	var resultArr []int
	for i := 0; i < len(nums); i++ {
		result += nums[i]
		resultArr = append(resultArr, result)
	}
	return resultArr

	//leetcode双百解法
	/*
		for i := 1; i < len(nums); i++ {
			nums[i] = nums[i] + nums[i-1]
		}
		return nums
	*/
}
