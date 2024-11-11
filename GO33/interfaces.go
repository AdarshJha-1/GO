package main

import "fmt"

type Shape interface {
	Area() float64
}
type Name interface {
	ShapeName() string
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) ShapeName() string {
	return "Rectangle"
}

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

func main() {
	rectangle := Rectangle{
		Height: 22,
		Width:  40,
	}
	// circle := Circle{Radius: 10}

	// fmt.Println(rectangle.Area())
	// fmt.Println(circle.Area())

	// shapes := []Shape{rectangle, circle}

	// for _, shape := range shapes {
	// 	fmt.Printf("Area: %f\n", shape.Area())
	// }

	var s Shape
	s = rectangle
	fmt.Println(s.Area())
	var n Name
	n = rectangle
	fmt.Println(n.ShapeName())

	// any type value --->  interface{}
}
