package main

import (
	"fmt"
)

func testChan(chan1 chan int)  {
	fmt.Println("zu se zhu le 0")
	chan1 <- 12
	fmt.Println("zu se zhu le 1")

}

func main() {

	chan1 := make(chan int)
	go testChan(chan1)

	select {
	//case <- chan1:
	//default:


	}
}