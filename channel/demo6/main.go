package main

import "fmt"

func send(sendCh chan<- int) {
	sendCh <- 10
	defer close(sendCh)
}

func read(readCh <-chan int) int {
	return <-readCh
}

func main() {
	sendCh := make(chan int)
	go send(sendCh)
	//read1 := read(sendCh)
	//fmt.Println(read1)
	//read2 := read(sendCh)
	//fmt.Println(read2)

	v, ok := <-sendCh
	fmt.Println(v, ok)

	v1, ok1 := <-sendCh
	fmt.Println(v1, ok1)
}
