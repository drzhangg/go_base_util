package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	a := struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}{"jerry", "上海"}
	bytes, _ := json.Marshal(&a)
	fmt.Println(string(bytes))

	var u User
	json.Unmarshal(bytes, &u)
	fmt.Println(u)
}
