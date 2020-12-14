package main

import (
	"log"
	"os/exec"
)

func main() {


	//i := 0
	//c := cron.New()
	//spec := "*/5 * * * * ?"
	//c.AddFunc(spec, func() {
	//	exec.Command("/bin/sh","./upload.sh")
	//
	//	i++
	//	log.Println("cron running:", i)
	//})
	//c.Start()
	//
	//select{}


	command := "echo hello"
	cmd := exec.Command("/bin/sh", "-c", command)
	bytes,err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	resp := string(bytes)
	log.Println(resp)
}
