package main

import "fmt"

func main() {
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list2)
	fmt.Println(list2)

	test()
}

func InsertSort(list []int) {
	// 从左往右依次比较，
	n := len(list)
	for i := 1; i <= n-1; i++ {
		deal := list[i] //待排序的数字
		j := i - 1      //待排序左边的数字

		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = deal
		}
	}
}

func test() {
	i := 10
	j := i-1
	for ;j >=0;j--{
		fmt.Println(j)
	}
	fmt.Println("end:",j)
}
