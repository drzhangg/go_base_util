package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	goCnt := 3
	for i := 0; i < goCnt; i++ {
		go busi(ch)
	}

	taskCnt := math.MaxInt64
	for i := 0; i < taskCnt; i++ {
		sendTask(i,ch)
	}

	wg.Wait()
}

func busi(ch chan int) {
	for t := range ch{
		fmt.Println("go task = ", t, ", goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int,ch chan int)  {
	wg.Add(1)
	ch <- task
}
