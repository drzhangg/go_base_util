package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func do() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("我是一行文字，我可能收不到值：%v", v)
		default:

		}
	}
}

func main() {
	file,err := os.Create("./cpu.prof")
	if err != nil {
		fmt.Printf("创建采集文件失败, err:%v\n", err)
		return
	}

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// 执行一段有问题的代码
	for i := 0; i < 4; i++ {
		go do()
	}
	time.Sleep(10 * time.Second)
}
