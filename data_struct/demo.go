package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8,9,10}
	s := SplitArray(a, 3)
	fmt.Println(s)
}

func SplitArray4(array []int, part int) (ret [][]int) {
	c := make([][]int, part)
	//x := len(array) / part
	var x int
	if len(array) % part == 0 {
		x = len(array) / part
	}else {
		x = len(array) / part + 1
	}
	b := x
	for i := 0; i < part-1; i++ {
		c[i] = append(c[i], array[b-x:b]...)
		b += x
	}
	c[part-1] = append(c[part-1], array[b-x:]...)
	return c
}

func SplitArray3(array []int, part int) (ret [][]int) {
	c := make([][]int, part)
	x := len(array) / part

	for i := 0; i < part-1; i++ {
		c[i] = append(c[i], array[:x]...)
		array = array[x:]
	}
	c[part-1] = append(c[part-1], array[:]...)
	return c
}

func SplitArray(array []int, part int) (ret [][]int) {
	alen := len(array)
	//zlen := alen / part
	var zlen int
	if alen % part == 0 {
		zlen = alen / part
	}else {
		zlen = alen / part + 1
	}
	fmt.Println("zzzz:", zlen)

	ret = make([][]int, part)
	for i := 0; i < part; i++ {
		if i == part-1 {
			ret[i] = append(ret[i], array...)
		} else {
			ret[i] = append(ret[i], array[:zlen]...)
			array = array[zlen:]
		}
	}
	return
}

func SplitArray1(array []int, part int) (ret [][]int) {
	alen := len(array)
	zlen := alen / part
	index, count := 0, 0
	ret = make([][]int, part)
	for i := 0; i < len(array); i++ {
		if count == zlen {
			index++
			if index == part {
				index = part - 1
			}
			count = 0
		}
		ret[index] = append(ret[index], array[i])
		count++
	}
	return
}

func SplitArray2(array []int, part int) (ret [][]int) {
	length := len(array)
	l := length / part
	fmt.Println("l:", l)
	/*
		ret[0] = [1,2,3] array[0:3] sum+=3 = 3
		ret[1] = [4,5,6] array[3:6]	sum+=3 = 6
		ret[2] = [7,8,9,10]	array[6:]
	*/

	ret = make([][]int, l)
	start, end := 0, part
	for i := 0; i < l; i++ {
		ret[i] = make([]int, 0)
		if i == l-1 {
			ret[i] = append(ret[i], array[start:]...)
		} else {
			ret[i] = append(ret[i], array[start:end]...)
			start = end
			end += 3
		}
	}
	return ret
}
