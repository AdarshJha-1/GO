package main

import "fmt"

// Embedding != Inheritance
// Go is object oriented language but does not have the concept of inheritance
// It used composition and promotion to promote code reuse

type Engine struct {
	HorsePower int
}

type GPS struct {
	Model string
}

func (e *Engine) Start() {
	fmt.Println("Engine started!!")
}

type Car struct {
	Model  string
	Engine // Embedding another type in structure
	GPS    // Embedding another type in structure
}

func (c *Car) Drive() {
	fmt.Printf("Driving my %s...\n", c.Model)
}

func main() {
	myCar := Car{
		Model: "Porsche",
		Engine: Engine{
			HorsePower: 572,
		},
		GPS: GPS{
			Model: "Germin",
		},
	}

	// fmt.Printf("%+v\n", myCar)
	// fmt.Println(myCar)
	fmt.Println(myCar.Model)
	fmt.Println(myCar.HorsePower)
	fmt.Println(myCar.GPS.Model)

	myCar.Start()
	myCar.Drive()

}
