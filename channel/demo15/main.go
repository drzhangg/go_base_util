package main

import (
	"fmt"
	"time"
)

func main() {
	worker(nil)
	time.Sleep(5*time.Second)
}

func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")

		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case <-time.Timer{}.C:
				fmt.Println("working")
			}
		}
	}()
}
