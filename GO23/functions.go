package main

import "fmt"

func main() {
	hello()
	helloToSomeone("Zoro")
	myFullName := fullName("Adarsh", "Jha")
	helloToSomeone(myFullName)
	myFullName2, nameLen := fullNameWithLength("Adarsh", "Jha")
	fmt.Println(myFullName2, nameLen)
}

func hello() {
	fmt.Println("Hello!!")
}

func helloToSomeone(name string) {
	fmt.Println("Hello!!", name)
}

func fullName(firstName, lastName string) string {
	return fmt.Sprintf("%s %s\n", firstName, lastName)
}

func fullNameWithLength(firstName, lastName string) (string, int) {
	fn := fmt.Sprintf("%s %s", firstName, lastName)
	return fn, len(fn)
}
