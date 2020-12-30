package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
	"time"
)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func main() {
	t := struct {
		time.Time
		N int `json:"n"`
	}{
		time.Date(2020, 12, 30, 0, 0, 0, 0, time.UTC),
		5,
	}

	type Demo struct {
		t time.Time
		N int
	}

	d := Demo{
		t: time.Date(2020, 12, 30, 0, 0, 0, 0, time.UTC),
		N: 10,
	}
	s,_ := json.Marshal(d)


	m, _ := json.Marshal(t)
	fmt.Printf("%s\n", m)
	fmt.Printf("%s\n", s)
}
