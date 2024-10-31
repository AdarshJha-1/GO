package main

import "fmt"

func main() {
	/*
		10 + 5
		10, 5 are operands
		+ is operator
	*/

	// Unary Operators
	// +, - , !, ^, ++, --, *, &
	num1 := 19
	fmt.Println(num1)
	num1++
	fmt.Println(num1)
	num1--
	fmt.Println(num1)

	/*
		// Binary Operators

		// Arithmetic Operators
		// +, -, *, /, % use with two operands
		var a, b int = 10, 4
		var c float64 = 69.69

		sum := a + b
		fmt.Println(sum)

		firstName := "Adarsh"
		lastName := "Jha"
		// fullName := firstName + " " + lastName
		fullName := fmt.Sprintf("%s %s", firstName, lastName)
		fmt.Println(fullName)

		sub := a - b
		fmt.Println(sub)

		mul := a * b
		fmt.Println(mul)

		div := a / b
		fmt.Println(div)

		remainder := a % b // remainder operator can only be used with int or uint
		fmt.Println(remainder)

		newSum := a + int(c) // type conversion
		fmt.Println(newSum)

		z := 10
		// z = z + 5
		z += 5 // shorthand ope. +=, -= , *=, /=
		fmt.Println(z)

		z -= 4
		fmt.Println(z)


			// zeroDiv := a / 0
			// fmt.Println(zeroDiv)
			// # command-line-arguments
			// GO09/operators.go:57:17: invalid operation: division by zero
	*/

	// Bitwise Logical & Shift Operators
	// &, |, ^, &^, <<, >>

	// Comparison Operators
	// ==, !=, <, >, <=, >=

	// Logical Operators
	// &&, ||, !

	/*
		Other Operators
		& : address operator
		* : de-referencing operator
		<- : receive value from channel

	*/

	/*
		Operator Precedence
		P E M D A S
		Parentheses, Exponents, Multiplication, Addition, Subtraction
	*/

}
