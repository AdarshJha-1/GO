package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("It's not Monday or Tuesday")
	}

	word := "Test"
	switch wordLen := len(word); wordLen {
	case 3:
		fmt.Println("Word len in 3")
	case 4:
		fmt.Println("Word len is 4")
	default:
		fmt.Println("Word len is not 3 or 4 but", wordLen)
	}
}
