package main

import "fmt"

func main() {
	fullName := func(s1, s2 string) string {
		return fmt.Sprintf("%s %s", s1, s2)
	}
	welcomeString := sayHello("ADARSH", "JHA", fullName)
	fmt.Println(welcomeString)
}

func sayHello(first, last string, fn func(string, string) string) string {
	fullName := fn(first, last)
	return fmt.Sprintf("Welcome %s", fullName)
}
