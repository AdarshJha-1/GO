package main

import "fmt"

func main() {
	var myString string
	fmt.Println("Zero value: ", myString) // ""
	myString = "Hello"
	fmt.Println("Value: ", myString)

	var myInt int
	/*
		ERROR, if not use
		 go run primitive.go
		# command-line-arguments
		./primitive.go:11:6: declared and not used: myInt
	*/
	fmt.Println("Zero value: ", myInt) // 0
	myInt = 100
	fmt.Println("Value: ", myInt)

	var myBool bool
	fmt.Println("Zero value: ", myBool) // false
	myBool = true
	/*
		statically typed can't change
		myBool = 10
		# command-line-arguments
		DataTypes/primitive.go:25:11: cannot use 10 (untyped int constant) as bool value in assignment
	*/
	fmt.Println("Value: ", myInt)
}
