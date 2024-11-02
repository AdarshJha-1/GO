package main

import (
	"fmt"
)

func main() {
	name := "Adarsh Jha"
	fmt.Println(string(name[0]))
	fmt.Println(string(name[len(name)-1]))

	var b rune = 'z'
	fmt.Println(string(b))

	// byteName := []byte(name)
	// fmt.Println(byteName)
	// runeName := []rune(name)
	// fmt.Println(runeName)

	firstName := name[:6]
	fmt.Println(firstName)

	nameHindi := "आदर्श"
	fmt.Println(nameHindi)
	fmt.Println(string(nameHindi[0]))

	unicodeString := []rune(nameHindi)
	fmt.Println(string(unicodeString[0]))
}
