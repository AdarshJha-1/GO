package main

import "fmt"

func main() {

	var ptr *int
	var emptyInterface interface{}

	fmt.Println(ptr)
	fmt.Println(ptr == nil)

	fmt.Println(emptyInterface)
	fmt.Println(emptyInterface == nil)

	emptyInterface = ptr
	fmt.Println(emptyInterface)
	fmt.Println(emptyInterface == nil)
}
