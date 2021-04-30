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

	sentence := "thequickbrownfoxjumpsoverthelazydog"

	m := 0
	for _,v := range sentence{
		fmt.Println(v)
		m |= 1 << (v - 'a')
	}
	fmt.Println(m)
	fmt.Println(1 << 26 -1)
	fmt.Println(1 << 26)
}
