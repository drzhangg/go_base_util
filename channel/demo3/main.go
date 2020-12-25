package main

import (
	"fmt"
)

func main() {
	var a = map[string]int{}
	a["age"] = 24
	fmt.Println(a)
	fmt.Println(a["age"])
}
