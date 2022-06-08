package main

import (
	"fmt"
	"time"
)

func main() {

	// work("programmer")
	// work("designer")
	// work("producer")
	// work("markerter")

	go work("programmer")
	go work("designer")
	go work("producer")
	go work("markerter")

	wait := 0
	fmt.Scanln(&wait)

}

func work(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "working", i, "hours")
		time.Sleep(time.Second) // 1초 멈춤
	}
}
