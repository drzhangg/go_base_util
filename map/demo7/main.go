package main

import "fmt"

func main() {
	var s1 = "123456789"
	var s2 = s1

	//fmt.Println(len(s2))
	//fmt.Println(1<<9)
	var a1 byte = 1 << len(s1) / 128
	var a2 byte = 1 << len(s2) / 128

	fmt.Println(a1, a2)

	fmt.Println(1 ^ 2)

	arr := []int{4, 4, 9, 2}
	var m = make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	fmt.Println(m)

	fmt.Println(2<<4)

}
