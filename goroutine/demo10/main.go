package main

import (
	"fmt"
	"time"
)

func main() {
	taskCh := make(chan int,100)
	go worker(taskCh)
	
	//塞任务
	for i := 0; i < 100; i++ {
		taskCh <- i
	}

	select {
	case <-time.After(20 * time.Second):
		fmt.Println("time after")
	//default:
	//	fmt.Println("default")

	}
}

func worker(taskCh <-chan int) {
	const N = 5

	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
