package main

import "fmt"

var processing bool

func flagTest() {
	processing = true

	fmt.Println(processing)

	defer func() {
		processing = false
	}()
}

func main() {
	flagTest()
	fmt.Println(processing)
}

