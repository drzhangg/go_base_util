package main

import "fmt"

func main() {
	defer function(1,function(3,0))
	defer function(2,function(4,0))
}

func function(i, v int) int {
	fmt.Println(i)
	return i
}
