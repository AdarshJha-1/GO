package math

// exported
const PI = 3.14

// exported
func Add(x, y int) int {
	return x + y
}

// unexported
func sub(x, y int) int {
	return x - y
}
