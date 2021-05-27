package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]*Student

func main() {
	list = make(map[string]*Student)

	student := Student{
		Name: "jerry",
	}

	list["student"] = &student
	list["student"].Name = "zhang"

	//s := list["student"]
	//s.Name = "boob"
	//fmt.Println(s)
	fmt.Println(list["student"])

	//list["student"] = s
	//fmt.Println(list)
}
