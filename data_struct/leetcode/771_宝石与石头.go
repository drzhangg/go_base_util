package leetcode

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	var count int
	var jm = make(map[int32]int)
	for _, v := range jewels {
		jm[v] = 0
	}

	for _,v := range stones{
		_, ok := jm[v]
		if ok {
			fmt.Println()
			count++
		}
	}
	return count
}
