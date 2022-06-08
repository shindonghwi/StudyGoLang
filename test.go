package main

import (
	"fmt"
	"time"
)

func main() {

	doneChannel := make(chan string, 5)

	// work("programmer")
	// work("designer")
	// work("producer")
	// work("markerter")

	go work("programmer", doneChannel)
	go work("designer", doneChannel)
	go work("producer", doneChannel)
	go work("markerter", doneChannel)
	go work("donghwi", doneChannel)

	for wait := range doneChannel {
		fmt.Println(wait)
	}

}

func work(s string, done chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "working", i, "hours")
		time.Sleep(time.Second) // 1초 멈춤
	}

	done <- s
}
