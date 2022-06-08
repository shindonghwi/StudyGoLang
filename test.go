package main

import "fmt"

func main() {

	// array
	var a [5]int
	var b [5]string
	c := [5]int{1, 2, 3, 4, 5}
	var d [2][3]int // 2차원

	a[0] = 100
	a[4] = 200

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// slice
	var aSlice []int
	var bSlice []string

	s := make([]int, 3)
	fmt.Println(aSlice, bSlice, s)

}
