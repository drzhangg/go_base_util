package main

import "fmt"

func main() {
	arr := []int{1,5,2,6,9,7}
	fmt.Println(InsertSort(arr))
}

func InsertSort(arr []int) []int {
	/*
		从第一个开始和第二个进行比较，如果第一个比第二个大，则第一个和第二个换位置
		后面同理，每次将后一个与前一个比较，如果前一个比后一个大，就将后一个和前一个互换位置
	*/

	if len(arr) < 2 || arr == nil {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		k := i - 1
		//如果当前的数比前一位数小，互换位置
		for arr[k] > temp && k >= 0 {
			k--
		}

		for j := i; j > k+1; j-- {
			arr[j] = arr[j-1]
		}

		arr[k+1] = temp
	}
	return arr
}
