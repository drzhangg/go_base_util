package main

import "fmt"

func main() {
	arr := []int{1,5,2,6,9,7}
	fmt.Println(BubbleSort(arr))
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
