package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	var cursorcopy uint64 = 0
	var results = make([]string,0)

	for {
		res := client.Scan(cursorcopy,"",1000)
		result,cursor,err := res.Result()
		if err != nil {
			panic(err)
		}
		if cursor == 0{
			break
		}
		fmt.Println("cursor:",cursor)
		cursor = cursor
		results = append(results, result...)
	}


	//fmt.Println("res:",results)
	//fmt.Println("res.len:", len(results))
	//fmt.Println("cursor:",cursorcopy)

	_ = http.ListenAndServe(":6060",nil)


}
