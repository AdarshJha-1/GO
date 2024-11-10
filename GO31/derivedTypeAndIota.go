package main

import "fmt"

type Age int

type Human struct {
	Name string
	Age  int
}

func (h Human) PrintDetails() {
	fmt.Println(h.Name, h.Age)
}
func PrintAge(age Age) {
	fmt.Println(age)
}

type Student Human

type Size int

const (
	ExtraSmall Size = iota
	Small
	Medium
	Large
	ExtraLarge
)

func main() {
	var young Age = 19
	var old Age = 66

	// var age int = 100
	// PrintAge(age) // type not compatible
	var age Age = 100
	PrintAge(age)

	fmt.Println(young + old)

	s1 := Student{Name: "Luffy", Age: 19}
	fmt.Println(s1)
	// s1.PrintDetails() // can't use

	fmt.Println(ExtraSmall, Small, Medium, Large, ExtraLarge)
}
