package main

import "fmt"

func main() {
	// basic form
	// var arr [5]int
	// fmt.Println(arr)

	// declaration with elements
	// var arr = [5]int{1, 2, 3, 4, 5}
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)

	// sparse array declaration
	// arr := [5]int{22, 3: 20, 100}
	// fmt.Println(arr)

	// implicit len. dec.
	// arr := [...]int{2, 3, 54, 33, 3}
	// fmt.Println(arr)

	// fmt.Println(len(arr))

	// // accessing array ele
	// fmt.Println(arr[0])

	arr := [2][2]int{
		{1, 2}, {3, 4},
	}
	fmt.Println(arr)
}
