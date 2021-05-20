package main

import (
	"fmt"
	"sort"
)

func main() {

	var m = map[string]interface{}{
		"hello":2,
		"morning":"ad",
		"keke":1,
		"jame":13,
		"age":25,
	}

	var keys []string
	for k := range m{
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println(keys)

	for _,v := range keys {
		fmt.Println(m[v])
	}
}
