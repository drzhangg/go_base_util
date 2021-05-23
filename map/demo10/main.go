package main

import (
	"fmt"
	"sort"
)

func main() {
	r := fib(7)
	fmt.Println(r)

	arr := []int{4,5,1,-4,-9,1,5}
	sort.Ints(arr)
	fmt.Println(arr)
}

func fib(n int) []int {
	var arr []int
	if n < 1 {
		return nil
	}

	arr = append(arr, 1,1)

	if n == 1 || n == 2 {
		//arr = append(arr, 1)
		return arr
	}

	prev, curr := 1, 1
	for i := 3; i < n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
		arr = append(arr, curr)
	}
	return arr
}
