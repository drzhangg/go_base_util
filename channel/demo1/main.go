package main

import (
	"fmt"
	"time"
)

type Format int

const (
	B Format = (iota + 1) * 2
	A
)

func hello() {
	fmt.Println("this is a goroutine")
}

func main() {
	fmt.Println(B)
	fmt.Println(A)
	go hello()

	time.Sleep(time.Second)

	fmt.Println("main function end")
	fmt.Println("12345")
}
