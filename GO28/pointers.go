package main

import "fmt"

func main() {
	// var numPtr *int
	// fmt.Println(numPtr)

	// num := 19
	// numPtr = &num
	// fmt.Println(&num)
	// fmt.Println(numPtr)
	// fmt.Println(*numPtr)

	a := 10
	fmt.Println(a)
	// increment(a)
	incrementWithPtr(&a)
	fmt.Println(a)
}

func increment(a int) {
	a++
}
func incrementWithPtr(a *int) {
	*a++
}
