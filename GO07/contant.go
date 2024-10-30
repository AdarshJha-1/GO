package main

import "fmt"

const x = 10     // untyped const
const y int = 11 // typed const

// Grouped constant
const (
	myName = "Adarsh Jha"
	myAge  = 19
)

func main() {
	fmt.Println(x)
	fmt.Println(myName, myAge)

	/*
		x = 100
		# command-line-arguments
		GO07/contant.go:16:2: cannot assign to x (neither addressable nor a map index expression)
	*/

	const yo int = 100
	const yoo float64 = 100.00

	// var a, b int = 10, 100
	/*
		const c = a + b
		fmt.Println(c)
		# command-line-arguments
		GO07/contant.go:27:12: a + b (value of type int) is not constant
	*/
}
