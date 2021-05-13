package main

import "fmt"

func main() {
	arr := []int{1,5,2,6,9,7}
	fmt.Println(selectSort(arr))
}

func selectSort(arr []int) []int {
	/*
	1.将第一个元素的设置成最小值
	2.依次往后进行比较，如果后面的值小于最小值，更新最小值下标
	3.继续往后进行对比
	 */
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	return arr
}
