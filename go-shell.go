// File: go-shell.go
//
// A starter Go program that can serve as the basis for new examples.
//
// D. Szarkowicz
// December, 2022

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

func checkError(error err) {
	if err != nil {
		panic(err)
	}
}
