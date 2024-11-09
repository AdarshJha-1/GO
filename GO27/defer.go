package main

import (
	"fmt"
)

func main() {
	message := "Hello,"

	greeting := func(name string) {
		fmt.Println(message, name)
	}

	defer greeting("Luffy")
	defer greeting("Zoro")
	// os.Exit(1)
	message = "Hii"
}
