package main

import "fmt"

func main() {
	arr := []int{1,5,2,6,9,7}
	fmt.Println(selectSort(arr))
}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		//min := i
		min := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
			}
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}

	}
	return arr
}
