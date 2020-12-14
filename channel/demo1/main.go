package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("this is a goroutine")
}

func main() {
	go hello()

	time.Sleep(time.Second)

	fmt.Println("main function end")
}
