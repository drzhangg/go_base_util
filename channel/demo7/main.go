package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
