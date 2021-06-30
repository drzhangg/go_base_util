package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	wg := sync.WaitGroup{}

	r := []int{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(j int) {
			client.Set("key"+fmt.Sprintf("%d", j), j, 0)
			r = append(r, j)
			wg.Done()
		}(i)
	}
	fmt.Println("len:",r)
	fmt.Println("main ok")
	wg.Wait()
}
