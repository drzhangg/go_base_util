package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var a chan int
	fmt.Println(a)
	if a == nil {
		a = make(chan int)
		fmt.Printf("%T\n", a)
		fmt.Println(a)
	}

	var p chan Person
	fmt.Println(p)
	if p == nil {
		p = make(chan Person)
		fmt.Printf("%T\n", p)
		fmt.Println(p)
	}
}
