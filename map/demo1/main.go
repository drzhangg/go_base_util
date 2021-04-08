package main

import "fmt"

func main() {
	ageMap := make(map[string]int)
	ageMap["qcrao"] = 18

	for name, age := range ageMap {
		fmt.Println(name, age)
	}
}
