package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var f, err = os.Create("./defer.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Fprintln(f, "something")
}

func a(b int) (int, error) {
	if b >= 10 {
		return -1, errors.New("Don't Exists")
	}
	return b, nil
}
