package main

import (
	"fmt"
	"math/rand"
)

func main() {
	g := getIntChan()
	for v := range g {
		fmt.Println(v)
	}

	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index

	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elm := <-intChannels[2]:
		fmt.Println("The third candidate case is selected. the elm is ",elm)
	default:
		fmt.Println("sssss")

	}
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
