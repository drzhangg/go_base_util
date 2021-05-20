package main

import (
	"fmt"
	"strconv"
)

func main() {

	f,_ := strconv.ParseFloat(fmt.Sprintf("%2.2f", (10000/0)*100),64)

	fmt.Println(f)
	fmt.Println(12312)
}

