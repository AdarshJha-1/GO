package main

import "fmt"

// func func1() {
// 	defer func() {
// 		fmt.Println("func1")
// 	}()
// 	func2()
// }
// func func2() {
// 	defer func() {
// 		fmt.Println("func2")
// 	}()
// 	panic("Error in func2")
// }

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong, this will be return value of recover func")
}

func main() {
	// defer func() {
	// 	fmt.Println("main func")
	// }()
	// func1()

	fmt.Println("Start of the program")
	panicExample()
	fmt.Println("Exited program")

}
