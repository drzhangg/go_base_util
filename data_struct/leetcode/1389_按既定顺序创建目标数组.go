package leetcode

func createTargetArray(nums []int, index []int) []int {
	l := len(nums)
	target := make([]int,l)

	for i:=0;i<l;i++{
		if index[i] == i{
			target[i] = nums[i]
		}else if index[i] < i{
			for j := i-1;j>=index[i];j--{
				target[j+1]=target[j]  //后移
			}
			target[index[i]]=nums[i]
		}
	}
	return target
}

