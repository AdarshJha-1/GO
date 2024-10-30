package main

import "fmt"

func main() {

	// uint8 (only +ve values)

	var smallPosValue uint8
	smallPosValue = 10
	fmt.Println(smallPosValue)
	/*
		smallPosValue = -10
		# command-line-arguments
		./integer.go:11:18: cannot use -10 (untyped int constant) as uint8 value in assignment (overflows)
	*/

	// int8 (+ve and -ve values)
	var smallPosNegValueInt int8
	smallPosNegValueInt = -100
	fmt.Println(smallPosNegValueInt)
	smallPosNegValueInt = 100
	fmt.Println(smallPosNegValueInt)

	/*
		all other datatype of "Integer" are same but has larger bit to store larger values
		- uint16, uint32, uint64
		- int16, int32, int64
	*/

	var myInt int = 203003
	fmt.Println(myInt)

	// type casting because int and int8 or other int data types are not same
	myInt = int(smallPosValue)
	myInt = int(smallPosNegValueInt)

	var myByte byte = 'A'
	/*
		byte is mainly use to represent raw data and also to distinguish b/w uint8 and byte since byte is an alias for uint8
	*/
	fmt.Println(myByte)

	// rune is alias for int32
	var myRune rune = 'B'
	fmt.Println(myRune)

	// can also store emoji
	var emojiRune rune = '😊'
	fmt.Println(emojiRune)
	fmt.Println(string(emojiRune))
}
