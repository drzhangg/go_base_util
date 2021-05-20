package main

import "fmt"

func main() {
	//fmt.Println(increase(2))
	//for i := 0; i < 5; i++ {
	//	defer fmt.Println(i)
	//}

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func increase(d int) (ret int) {
	defer func() {
		ret++
	}()
	return d
}

func f1() (result int) {
	defer func() {
		result++
	}()
	result = 0
	return result
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
