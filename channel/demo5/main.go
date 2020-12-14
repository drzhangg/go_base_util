package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine before sleep")
	time.Sleep(4 * time.Second)
	done <- true

	fmt.Println("hello goroutine after sleep")
}

func main() {

	done := make(chan bool)
	go hello(done)
	fmt.Println(<-done)
	fmt.Println("main")
}
