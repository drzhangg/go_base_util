package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3}
	fmt.Println("a.len:",len(a))
	fmt.Println("a.cap:",cap(a))
	fmt.Printf("%p\n",a)

	b := a
	b = append(b, 4)
	fmt.Println("b.len:",len(b))
	fmt.Println("b.cap:",cap(b))
	fmt.Printf("%p\n",b)
}
