package main

import "fmt"

type student struct {
	Name string
	Age int
}

func main() {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for k,stu := range stus{
		//fmt.Println(*&stu)
		m[stu.Name] = &stus[k]
	}

	for k,v := range m{
		fmt.Println(k ,"=>", v)
	}
	//fmt.Println(m)
}
