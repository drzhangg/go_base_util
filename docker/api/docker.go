package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"io"
	"strings"
)

func main() {

	cli, err := client.NewClient("http://192.168.31.232:2375", "v1.41", nil, map[string]string{"User-Agent": "engine-api-cli-1.0"})
	//client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}

	imageName := "busybox:latest"
	events, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(events)

	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		Progress       string `json:"progress"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}

	var event *Event
	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		fmt.Printf("EVENT: %+v\n", event)
	}

	if event != nil {
		if strings.Contains(event.Status, fmt.Sprintf("Downloaded newer image for %s", imageName)) {
			// new
			fmt.Println("new")
		}

		if strings.Contains(event.Status, fmt.Sprintf("Image is up to date for %s", imageName)) {
			// up-to-date
			fmt.Println("up-to-date")
		}
	}

	//resp, err := http.Get("http://192.168.31.232:2375/images/search")
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//
	//fmt.Println(resp.Body)
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("err1:", err)
	//	return
	//}
	//fmt.Println(string(body))
}
