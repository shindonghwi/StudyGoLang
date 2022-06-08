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

	var scanData1 int
	var scanData2 int
	var scanData3 int

	fmt.Scan(&scanData1, &scanData2, &scanData3)

	println("scanData: ", scanData1, scanData2, scanData3)

}
