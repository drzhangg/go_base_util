package main

import "fmt"

func main() {
	s := "abcdefg"
	fmt.Println(len(s))
	s1 := s[2:] + s[:2]
	fmt.Println(s1)
	fmt.Println(s)
	for _, v := range s {
		fmt.Printf("%s",v)
	}
}
