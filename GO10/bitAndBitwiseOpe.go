package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 202
	// var b uint8 = 141

	// OR
	// c := a | b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// AND
	// c := a & b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// XOR
	// c := a ^ b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// AND NOT
	// c := a &^ b
	// fmt.Println(c)
	// fmt.Println(strconv.FormatUint(uint64(a), 2))
	// fmt.Println(strconv.FormatUint(uint64(b), 2))
	// fmt.Println(strconv.FormatUint(uint64(c), 2))

	// INVERT
	c := ^a
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	// odd or even
	fmt.Println("ODD:", 11&1)
	fmt.Println("EVEN:", 10&1)
}
