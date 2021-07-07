package main

import "fmt"

func main() {
	x := 200
	var any interface{} = x
	fmt.Println(any)

	g := Gopher{"go"}
	var c coder = g
	fmt.Println(c)
}

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("i am coding %s language\n",p.language)
}

func (p Gopher) debug() {
	fmt.Printf("i am debugging %s language\n",p.language)
}
