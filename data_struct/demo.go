package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := SplitArray(a, 7)
	fmt.Println(s)
}

func SplitArray(array []int, part int) (ret [][]int)  {
	alen := len(array)
	zlen := alen / part

	ret = make([][]int,part)
	for i := 0; i < part; i++ {
		if i ==  part - 1{
			ret[i] = append(ret[i], array...)
		}else {
			ret[i] = append(ret[i], array[:zlen]...)
			array = array[zlen:]
		}
	}
	return
}

func SplitArray1(array []int, part int) (ret [][]int)  {
	alen := len(array)
	zlen := alen / part
	index ,count :=0,0
	ret = make([][]int,part)
	for i := 0; i < len(array); i++ {
		if count == zlen{
			index++
			if index== part{
				index = part-1
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
	fmt.Println("l:",l)
	/*
		ret[0] = [1,2,3] array[0:3] sum+=3 = 3
		ret[1] = [4,5,6] array[3:6]	sum+=3 = 6
		ret[2] = [7,8,9,10]	array[6:]
	*/

	ret = make([][]int,l)
	start, end := 0, part
	for i := 0; i < l; i++ {
		ret[i] = make([]int,0)
		if i == l-1 {
			ret[i] = append(ret[i], array[start:]...)
		}else {
			ret[i] = append(ret[i], array[start:end]...)
			start = end
			end += 3
		}
	}
	return ret
}
