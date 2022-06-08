package main

import (
	"fmt"
)

func main() {

	var a = 5.0
	b := 4.0

	println(a, b)

	a++
	b--

	println(&a, &b)

	var scanData int

	fmt.Scan(&scanData)

	println("scanData: ", scanData)

}
