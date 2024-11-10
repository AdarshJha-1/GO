package main

import "fmt"

type Person struct { // Concrete type with data and methods included
	Name string
	Age  int
}

// Methods are associated with type in go
func (p Person) GetDetails() string { // Value receiver, it does not allow to make changes in values
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p *Person) ChangeName(name string) { // Pointer receiver, it allow to make change in values
	p.Name = name
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var p1 Person = Person{
		Name: "Money D. Luffy",
		Age:  19,
	}

	fmt.Println(p1.GetDetails())
	p1.ChangeName("Roronoa Zoro")
	fmt.Println(p1.GetDetails())
}
