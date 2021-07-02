package main

import (
	"fmt"
)

func main() {
	s := make([]string,200)
	for i := 0; i < 200; i++ {
		s = append(s, "test"+ fmt.Sprintf("%d",i))
	}
	fmt.Println(s)

	//now := time.Now()

}
