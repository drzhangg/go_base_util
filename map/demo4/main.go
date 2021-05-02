package main

func main() {
	//nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	//s := []int{1, 2, 3, 4}
	//fmt.Println(s[0:0])

	var max int
	for i := 0; i < len(index); i++ {
		if max < index[i]{
			max = index[i]
		}
	}
}
