package main

import "fmt"

func main() {
	command := "ab"
	var newStrMap = make(map[uint8]int)
	for i := 0; i < len(command); i++ {
		newStrMap[command[i]] = 0
	}
	fmt.Println(newStrMap)

	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	var r int
	for i := 0; i < len(words); i++ {
		var count int
		for j := 0; j < len(words[i]); j++ {
			if _, ok := newStrMap[words[i][j]]; ok {
				count++
			} else {
				count = 0
				break
			}
		}
		if count == len(words[i]){
			r ++
		}
	}
	fmt.Println(r)
}
