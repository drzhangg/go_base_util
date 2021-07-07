package main

import "fmt"

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct {

}

func (g Go) sayHello() {
	fmt.Println("hi,go")
}

type Php struct {

}

func (p Php) sayHello() {
	fmt.Println("hi, php")
}

func main() {
	golang := Go{}
	php := Php{}

	sayHello(golang)
	sayHello(php)
}
