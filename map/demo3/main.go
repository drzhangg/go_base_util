package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				fmt.Println(i, j)
				count++
			}
		}
	}
	fmt.Println(count)
}