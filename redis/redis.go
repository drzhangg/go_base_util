package main

import (
	"context"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.31.191:6379",
	})

	client.Set("name", "jerry", 0)
}
