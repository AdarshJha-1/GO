package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// b := a[2:7] // 2 to 6 index, 7 won't be included
	// b := a[:8]
	// b := a[:]
	// b := a[1:]

	// b := a[:6:10]
	// // fmt.Println(a)
	// // fmt.Println(b)

	// // a[4] = 40

	// // fmt.Println(a)
	// // fmt.Println(b)

	// fmt.Println(b)

	b := make([]int, 5)
	copy(b, a)
	fmt.Println(b)

}
