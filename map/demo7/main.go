package main

import "fmt"

func main() {
	var s1 = "123456789"
	var s2 = s1

	//fmt.Println(len(s2))
	//fmt.Println(1<<9)
	var a1 byte = 1 << len(s1) / 128
	var a2 byte = 1 << len(s2) / 128

	fmt.Println(a1,a2)

	fmt.Println(1^2)
}


