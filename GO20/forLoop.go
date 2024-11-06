package main

import "fmt"

func main() {
	// a := 10

	// for {
	// 	fmt.Println(a)
	// 	a--
	// 	if a == 0 {
	// 		break
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	name := "Adarsh"
	for i, c := range name {
		fmt.Println(i, c, string(c))
	}

	rollNoName := map[int]string{
		1: "Luffy",
		2: "Zoro",
		3: "Sanji",
		4: "Nami",
	}

	for k, v := range rollNoName {
		fmt.Println(k, v)
	}
}
