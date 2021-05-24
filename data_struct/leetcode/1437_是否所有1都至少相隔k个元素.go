package leetcode

/*
输入：nums = [1,0,0,0,1,0,0,1], k = 2
输出：true
解释：每个 1 都至少相隔 2 个元素
 */
func kLengthApart(nums []int, k int) bool {
	target := 0
	flag := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 && flag {	//第2个以后的1
			if target >= i{
				return false
			}
			target = i + k
		}else if nums[i] == 1 && !flag{ //当遇到第一个1的时候，flag改成true
			target = i + k
			flag = true
		}
	}
	return true
}
