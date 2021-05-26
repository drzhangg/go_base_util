package main

import "fmt"

func main() {

	fmt.Println(f())
	fmt.Println(f1())

	m := map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}

	for k,v := range m{
		delete(m,"two")
		m["four"] = 4
		fmt.Printf("k:%v, v:%v\n",k,v)
	}
}

func DeferFunc1(i int) (t int) {

	fmt.Println("t = ", t)

	return 2
}

func f() (r int){
	t := 5
	defer func() {
		t = t + 5
	}()
	fmt.Println("r:",r)
	return t
}

func f1() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
