package main

import "fmt"

func main() {
	var b string
	a := "hello"
	a = "world"

	if a == "world" {
		fmt.Println("a is world")
	} else {
		fmt.Println("a is not world")
	}

	if b == "" {
		fmt.Println("Primitive zero values are always not nil")
	}
}
