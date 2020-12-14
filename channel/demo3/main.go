package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go i:",i)
		}
	}()

	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println("执行：",i)
	}
}
