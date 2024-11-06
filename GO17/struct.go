package main

import (
	"fmt"
)

type student struct {
	name    string
	age     int
	courses []string
}

func main() {
	var s1 student
	fmt.Printf("%+v\n", s1)

	s1 = student{
		"luffy", 19, []string{"math", "science"},
	}

	fmt.Printf("%+v\n", s1)

	s2 := student{
		name:    "zoro",
		age:     21,
		courses: []string{"hindi :0", "geography"},
	}
	fmt.Printf("%+v\n", s2)

	fmt.Println(s2.name)
	s2.courses = append(s2.courses, "english")
	fmt.Println(s2.courses)

	// Anonymous struct
	guardian := struct {
		name  string
		email string
	}{
		name:  "dragon",
		email: "onDDesk@gmail.com",
	}
	fmt.Println(guardian)
}
