package main

import "fmt"

func test(i []int, x int) () {
	i = append(i, x)
	fmt.Println("dfsdfa", i)
}

func main() {
	var i = make([]int, 0, 10)
	test(i, 15)
	fmt.Println(i)
}
