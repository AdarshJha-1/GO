package main

import "fmt"

func main() {
	age := 1023

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are teenager")
	} else {
		fmt.Println("You are still a brat")
	}

	if even := isEven(age); even {
		fmt.Println("Age is even")
	} else {
		fmt.Println("Age is not even")
	}
}

func isEven(n int) bool {
	return n&1 == 0
}
