package main

import (
	"fmt"
	"math"
)

func main() {
	//var s []string
	//for i := 0; i < 9999; i++ {
	//	go func() {
	//		s = append(s, "脑子进煎鱼了")
	//	}()
	//}
	//
	//fmt.Printf("进了%d只煎鱼", len(s))

	//65280

	//-1 >> ? = 65280

	//0001
	//1110
	//1111
	//num := 65280
	//result := 0
	//for num >0 {
	//	num = num >> 1
	//	result++
	//}
	//fmt.Println(result)
	//
	//fmt.Println(-1 << 16)
	//
	//now := time.Now()
	//r := now.String() > "2020-05-01"
	//fmt.Println(r)

	m := make(map[float64]interface{})
	m[math.NaN()] = 1
	fmt.Println(m)
	fmt.Println(m[math.NaN()])
}
