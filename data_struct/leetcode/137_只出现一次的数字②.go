package leetcode

func singleNumber2(nums []int) int {
	number,res := 0,0
	for i := 0; i < 64; i++ {
		number = 0
		for _, v := range nums{
			number += (v >> i) &1
		}
		res |= (number) %3 << i
	}
	return res
}
