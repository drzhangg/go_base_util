package main

import "fmt"

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age += 1
}

func main() {
	qcrao := Person{18}

	fmt.Println(qcrao.howOld())

	qcrao.growUp()

	fmt.Println(qcrao.howOld())

	// ------------------------

	stefno := &Person{100}
	fmt.Println(stefno.howOld())

	stefno.growUp()

	fmt.Println(stefno.howOld())
}
