package main

import "fmt"

func main() {
	a := []int{1,3,5,7,8,9,10}
	r := bsearch(a,7,5)
	fmt.Println(r)
}

func bsearch(a []int, n, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) /2
		if a[mid] == value{
			return mid
		}else if a[mid] < value {
			low = mid +1
		}else {
			high = mid -1
		}
	}
	return -1
}
