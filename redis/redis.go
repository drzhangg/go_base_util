package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	r := client.Get("name")
	result, err := r.Result()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("result:", result)

	//client.Set("name", "jerry", 0)
}
