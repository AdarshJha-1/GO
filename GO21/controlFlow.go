package main

import "fmt"

func main() {
	a := 100

	for { // works same as do while loop
		fmt.Println(a)
		a--
		if a == 90 {
			break // breaks loop when condition is met
		}
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
