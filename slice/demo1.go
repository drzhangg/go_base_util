package main

import "fmt"

func main() {
	//sli1 := []int{1,2,3,4,5,6}
	//one := sli1[0]
	//copy(sli1,sli1[1:])
	//sli1 = sli1[:len(sli1)-1]
	//
	//fmt.Println(one)
	//fmt.Println(sli1)
	//fmt.Println(len(sli1), cap(sli1))


	sli2 := []int{1,2,3,4,5,6}
	one := sli2[0]
	//copy(sli2,sli2[1:])
	fmt.Println(len(sli2),cap(sli2))

	sli2 = sli2[1:]
	sli2 = append(sli2, 7)
	fmt.Println(one)
	fmt.Println(sli2)
	fmt.Println(len(sli2), cap(sli2))
}
