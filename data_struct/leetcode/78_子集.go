package leetcode

func subsets(nums []int) [][]int {
	ans := make([][]int,1)
	ans[0] = []int{}

	for _, x:= range nums{
		for _,arr := range ans{
			a := make([]int,len(arr))
			copy(a,arr)
			ans = append(ans, append(a,x))
		}
	}
	return ans
}
