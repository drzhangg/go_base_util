package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 1, 1, 3}
	//var count int
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i] == nums[j] && i < j {
	//			fmt.Println(i, j)
	//			count++
	//		}
	//	}
	//}
	//fmt.Println(count)

	//result := []int{1,2,4,6}
	//s := result[len(result)-1]
	//fmt.Println(s)

	a := 234
	var  b int
	c:=1
	for a > 0 {
		m := a % 10
		a/=10
		b+=m
		c*=m
		fmt.Println(m)
	}
	fmt.Println(b,c)

	fmt.Println(12/2)
}
