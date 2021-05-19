package main

import "fmt"

type Level int

const (
	a Level = iota
	b
	c
	d
)

func (o Level) String() string {
	switch o {
	case a:
		return "A"
	case b:
		return "B"
	case c:
		return "C"
	case d:
		return "D"
	default:
		return "..."
	}
}

func main() {
	i := a
	fmt.Printf("%d",i)
}
