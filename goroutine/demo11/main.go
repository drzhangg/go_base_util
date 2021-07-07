package main

import "time"

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello "
	time.Sleep(2 * time.Second)
	done <- true

}

func main() {
	go aGoroutine()
	<-done
	println(msg)
}
