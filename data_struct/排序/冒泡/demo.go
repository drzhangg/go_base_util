package main

import "fmt"

func main() {
	a:= []int{1,3,5,6,2}
	n := 5
	bubbleSort(a,n)
	fmt.Println(a)
}

// 冒泡排序，a 表示数组，n 表示数组大小
func bubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	/*
		第一轮
	*/

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] == a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

}
