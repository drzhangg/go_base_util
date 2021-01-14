package main

import "fmt"

func main() {
	r := find([]int{1,3,4,5},4,5)
	fmt.Println(r)
}

func find(arr []int,n ,x int) int  {
	var pools = -1
	for i := 0; i < n; i++ {
		if arr[i] == x {
			pools = i
		}
	}
	return pools
}
