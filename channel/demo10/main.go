package main

import (
	"fmt"
	_ "net/http/pprof"
)

func main() {
	//var ch = make(chan int ,1000)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//}()
	//
	//go func() {
	//	v,ok := <- ch
	//	if !ok {
	//		fmt.Println("closed")
	//		return
	//	}
	//
	//	fmt.Println("a:",v)
	//}()
	//
	//close(ch)
	//fmt.Println("main close")
	//time.Sleep(100 * time.Second)


	//var ss = []string{"i","love","golang"}
	//for k,v := range ss{
	//   ss = append(ss,"hello")
	//	fmt.Println(k,v)
	//}

	var ss = []string{"i","love","golang"}
	for i := 0; i < len(ss); i++ {
		ss = append(ss, "hello")
		fmt.Println(i,ss[i])
	}


	/*
	var ss = []string{"i","love","golang"}

	for i:=0;i<10;i++{
		go func (i int,s string){
			fmt.Println(i,s,k,v)
		}(k,v)
	}
	time.Sleep(100 * time.Second)
	 */

}
