package main

import "fmt"

func main() {
	var myString string
	fmt.Println("Zero value:", myString) // ""
	myString = "Hello World"
	fmt.Println(myString)

	myString = "hello\nworld"
	fmt.Println(myString)

	myString = `Yokoso
Watashi no Soul society`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "Roronoa"
	lastName = "Zoro"

	var fullName string
	// fullName = firstName + " " + lastName
	// fmt.Println(fullName)

	fmt.Printf("%s %s\n", firstName, lastName)
	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}
