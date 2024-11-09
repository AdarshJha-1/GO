package main

import "fmt"

func main() {
	// message := "Hello"
	// sayHello := func(name string) string {
	// 	return fmt.Sprintf("%s, %s", message, name)
	// }

	// fmt.Println(sayHello("Adarsh"))

	// func() {
	// 	fmt.Println("Hellooooooooo")
	// }()

	// func(s1 string) {
	// 	fmt.Println(s1)
	// }(message)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sliceOperation(numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(result)
}

func sliceOperation(s []int, op func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = op(v)
	}
	return result
}
