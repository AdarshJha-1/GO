package main

import "fmt"

// any == interface{}

func printValue(value interface{}) {
	fmt.Println(value)
}

func main() {
	mixedSlice := []interface{}{
		343, "Yooo", true, 434.433,
	}
	printValue(mixedSlice)

	// var emptyInterface interface{}
	var emptyInterface any
	emptyInterface = "Hello World"

	if str, ok := emptyInterface.(string); ok {
		fmt.Println("The underlying value is a strings:", str)
	} else {
		fmt.Println("Not true")
	}

}
