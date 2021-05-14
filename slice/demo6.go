package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {
	//写入
	data := []string{"煎鱼", "咸鱼", "烤鱼", "蒸鱼"}

	for i := 0; i < 4; i++ {
		go func(j int) {
			fmt.Println("j:",j)
			m.Store(j, data[j])
		}(i)
	}
	time.Sleep(time.Second)

	v, ok := m.Load(0)
	fmt.Printf("Load: %v, %v\n", v, ok)


	m.Delete(1)

	v,ok = m.LoadOrStore(1,"洗浴")
	fmt.Printf("LoadOrStore: %v, %v\n",v,ok)

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Range: %v, %v\n",key,value)
		return true
	})
}
