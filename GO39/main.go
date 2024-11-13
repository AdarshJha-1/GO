package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	luffy := color.RedString("Monkey D. Luffy")
	naruto := color.YellowString("Uzumaki Naruto")
	fmt.Println(luffy, naruto)
}
