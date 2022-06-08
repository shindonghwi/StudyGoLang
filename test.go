package main

import (
	"fmt"
	"time"
)

func main() {

	work("hello world")
}

func work(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "working", i, "hours")
		time.Sleep(time.Second) // 1초 멈춤
	}
}
