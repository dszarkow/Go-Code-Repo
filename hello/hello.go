// File: hello.go
//
// A simple Go program based on go-shell-go
//
// D. Szarkowicz
// January, 2023

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	var err error
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("err was nil")
	}
}
