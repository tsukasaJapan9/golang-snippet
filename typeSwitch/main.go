package main

import "fmt"

func main() {
	var x interface{}
	x = 10

	switch x.(type) {
	case int:
		v := x.(int)
		fmt.Printf("type=%T, value=%d\n", x, x)
		// error
		//x += 10

		fmt.Printf("type=%T, value=%d\n", v, v)
		v += 10
		fmt.Printf("type=%T, value=%d\n", v, v)

	default:
		fmt.Println("don't know")
	}
}
