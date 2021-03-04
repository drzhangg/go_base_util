package main

import "fmt"

type innerS struct {
	inner1 int
	inner2 int
}

type outerS struct {
	b int
	innerS
}

func main() {
	out := outerS{}
	out.b = 1
	out.innerS.inner1 = 3
	out.inner2 = 2
	fmt.Println(out)
}
