package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:2375/version")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err1:", err)
		return
	}
	fmt.Println(string(body))
}
