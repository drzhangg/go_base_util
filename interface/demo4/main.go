package main

import "fmt"

type Person interface {
	job()
	growUp()
}

func whatJob(p Person) {
	p.job()
}

func growUp(p Person) {
	p.growUp()
}

type Student struct {
	age int
}

func (p Student) job() {
	fmt.Println("i am a student")
	return
}

func (p *Student) growUp() {
	p.age+=1
	return
}

type Programmer struct {
	age int
}

func (p Programmer) job() {
	fmt.Println("i am a programmer")
	return
}

func (p Programmer) growUp() {
	p.age+=10
	return
}

func main() {
	qrao := Student{18}
	whatJob(&qrao)

	growUp(&qrao)
	fmt.Println(qrao)

	stefno := Programmer{100}
	whatJob(stefno)

	growUp(stefno)
	fmt.Println(stefno)
}
