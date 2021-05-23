package leetcode

import "fmt"

func sumOfUnique(nums []int) int {
	n := rrr(nums)
	fmt.Println(n)
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

func rrr(arr []int) []int {
	var res []int
	var am = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		am[arr[i]]++
	}
	for i := 0; i < len(arr); i++ {
		if am[arr[i]] == 1{
			res = append(res, arr[i])
		}
	}
	return res
}
