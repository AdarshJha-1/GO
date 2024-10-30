package main

import "fmt"

func main() {

	var smallFloat float32
	fmt.Println("Zero value:", smallFloat) // 0
	smallFloat = 69.696969
	fmt.Println(smallFloat)

	var bigFloat float64
	bigFloat = 69.69696969696969
	fmt.Println(bigFloat)
}
