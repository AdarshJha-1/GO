package main

import "fmt"

// Package level variable
var a int = 10

func main() {
	fmt.Println(a)
	weAre()

	{
		// block scoped variable a
		a := 15 // Shadowing
		fmt.Println(a)
	}
	fmt.Println(a)

}
