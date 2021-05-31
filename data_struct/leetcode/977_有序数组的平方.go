package leetcode

func sortedSquares(nums []int) []int {
	//循环
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = nums[i] * nums[i]
	//}
	//sort.Ints(nums)
	//return nums

	//双指针
	n := len(nums)
	arr := make([]int,n)
	i,j := 0, n-1
	for pos := n-1; pos >=0 ; pos-- {
		if v,w := nums[i]*nums[i],nums[j]*nums[j];v > w{
			arr[pos] = v
			i++
		}else {
			arr[pos] = w
			j--
		}
	}
	return arr
}
