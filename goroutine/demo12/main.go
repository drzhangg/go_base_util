package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	//senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):

				}
			}
		}()
	}

	//receiver
	go func() {
		for v := range dataCh {
			if v == Max-1 {
				fmt.Println("send stop signal to senders")
				fmt.Println("end value:",v)
				close(stopCh)
				return
			}

			fmt.Println(v)
		}
	}()

	select {
	case <-time.After(5 * time.Second):

	}
}
