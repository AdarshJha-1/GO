package main

import "fmt"

func main() {

	// var a []int // nil slice, no underling array
	// fmt.Println(a == nil)

	// a = []int{} // empty slice, declare underling array
	// fmt.Println(a == nil)

	// a := []int{1, 2, 3}
	// fmt.Println(a)

	// a = []int{3, 6: 11, 100}
	// fmt.Println(a)

	// a := make([]int, 5)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	// a := make([]string, 5, 10)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	// a := make([]int, 5)
	// a := make([]int, 0)
	// // a = append(a, 122)
	// a = append(a, 122, 1222, 122222, 1222222)
	// fmt.Println(a)

	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	a = append(a, b...)
	fmt.Println(a)

}
