package main

import "fmt"

// type Calculator struct {
// }

// func (c *Calculator) Add(a, b int) int {
// 	return a + b
// }

// func AddFunction(a, b int) int {
// 	return a + b
// }

// func ArithmeticOperation(f func(int, int) int, a, b int) int {
// 	return f(a, b)
// }

type Person struct {
	Name string
	Age  int
}

func (p Person) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) SetLength(newLength int) {
	r.Length = newLength
}

func main() {
	// c := Calculator{}
	// fmt.Println(AddFunction(10, 18))
	// fmt.Println(ArithmeticOperation(AddFunction, 10, 18))
	// fmt.Println(ArithmeticOperation(c.Add, 10, 18))

	f1 := Person.GetDetails
	p := Person{Name: "Nami", Age: 20}
	fmt.Println(f1(p))

	f2 := (*Rectangle).SetLength // SetLength is a pointer receiver thats why we need to type cast it

	r := Rectangle{
		Length: 10,
		Width:  20,
	}

	fmt.Println(r)
	f2(&r, 20)
	fmt.Println(r)
}
