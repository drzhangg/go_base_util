package main

import (
	"fmt"
	"time"
)

func goroutineA(a <-chan int) {
	fmt.Println("wait goroutine A")
	val := <-a
	fmt.Println("G1 received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	fmt.Println("wait goroutine B")
	val := <-b
	fmt.Println("G2 reveived data:", val)
	return
}

func main() {
	ch := make(chan int)
	go goroutineA(ch)
	go goroutineB(ch)
	ch <- 3
	time.Sleep(time.Second)
}
