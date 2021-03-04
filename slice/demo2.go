package main

import "fmt"

type F interface {
	f()
	read() string
}

type S1 struct {
	data string
}

func (s S1) f() {
	s.data = "b"
}

func (s S1) read() string {
	return s.data
}

type S2 struct {
	data string
}

func (s *S2) f()  {
	s.data = "b"
}

func (s *S2)read() string {
	return s.data
}

var f1 F = S1{"a"}
var f2 F = &S2{"a"}

func main() {
	f1.f()
	fmt.Println(f1.read())
	f2.f()
	fmt.Println(f2.read())
}

func bb() []int {
	m := make([]int, 10)
	return m
}
