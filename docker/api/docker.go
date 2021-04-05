package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.31.232:2375/images/search")
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
