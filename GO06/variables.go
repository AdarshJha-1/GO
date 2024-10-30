package main

import "fmt"

var myInt int = 10 // package level variable

// Group similar variables
// can be moved/create inside a functions
var (
	myName string = "Adarsh"
	myAge  int    = 19
)

func main() {

	fmt.Println(myInt)
	fmt.Println(myName, myAge)

	// Implicit type assignment
	var num0 int // zero value 0
	num0 = 120
	fmt.Println(num0)

	var num1 int = 34
	fmt.Println(num1)

	// Explicit type assignment
	var num2 = 1020
	fmt.Println(num2)

	// multiple variables declaration
	var num3, num4, num5 int = 12, 14, 16
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)

	// Short hand variable declaration
	day := 30
	month := 10
	year := 2024

	fmt.Printf("%d/%d/%d\n", day, month, year)

	day = 1
	month = 11

	fmt.Printf("%d/%d/%d\n", day, month, year)

}
