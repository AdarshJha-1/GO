package main

import "fmt"

func main() {
	// var nameAge map[string]int // nil map
	// fmt.Println(nameAge == nil)
	// fmt.Println(nameAge["luffy"]) // zero value of int = 0

	/*
		nameAge["zoro"] = 21
		panic: assignment to entry in nil map

		goroutine 1 [running]:
		main.main()
			/home/adarsh/Codes/Go-Code/go-learning/GO16/maps.go:10 +0xbf
		exit status 2
	*/

	// nameAge := map[string]int{} // emoty map
	// fmt.Println(nameAge)
	// fmt.Println(nameAge == nil)
	// fmt.Println(nameAge["zoro"]) // zero value of int

	// nameAge := make(map[string]int) // emoty map
	// fmt.Println(nameAge)
	// fmt.Println(nameAge == nil)
	// fmt.Println(nameAge)

	// nameAge := map[string]int{}	// emoty map

	nameAge := map[string]int{ // map with pre-initialized values
		"usopp": 19,
		"sanji": 21,
		"nami":  20,
	}

	nameAge["luffy"] = 19
	nameAge["zoro"] = 21
	nameAge["robin"] = 30
	// fmt.Println(len(nameAge))
	// fmt.Println(nameAge)
	fmt.Println(nameAge["zoro"])
	fmt.Println(nameAge["brook"]) // brook is not in map, prints zero value for age int

	fmt.Println(nameAge["usopp"])
	nameAge["usopp"] = 200 // map can have only unique key so it will set existing key's value to new value
	fmt.Println(nameAge["usopp"])

	nameGrade := map[string]int{
		"luffy": 10,
		"zoro":  11,
		"nami":  95,
	}
	// fmt.Println(nameGrade["luffy"])

	gradeLuffy, ok := nameGrade["luffy"] // luffy present in map got value and true
	fmt.Println(gradeLuffy, ok)

	gradeChopper, ok := nameGrade["chopper"] // chopper is not present in map got zero value for corresponding data type (int here) and false
	fmt.Println(gradeChopper, ok)
}
