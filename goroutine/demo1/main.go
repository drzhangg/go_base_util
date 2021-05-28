package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	task_cnt := math.MaxInt64

	ch := make(chan bool,3)

	for i := 0; i < task_cnt; i++ {
		ch <- true
		go busi(ch,i)
	}
}

func busi(ch chan bool, i int) {

	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
}
