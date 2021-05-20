package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		fmt.Println("goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("goroutine 2")
		wg.Done()
	}()

	wg.Wait()

}

func testContext(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
}
