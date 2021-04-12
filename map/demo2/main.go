package main

import "fmt"

func main() {
	var ms []map[string]interface{}
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["key_"+fmt.Sprint(i)] = i
		ms = append(ms, m)
	}
	fmt.Println(ms)
}
