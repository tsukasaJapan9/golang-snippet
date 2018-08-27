package main

import "fmt"

type data struct {
	foo string
	bar string
}

func main() {
	d := data {
		foo: "dataFoo",
		bar: "dataBar",
	}

	fmt.Println(d)
}