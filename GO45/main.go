package main

import (
	"fmt"
	"time"
)

// Time Go

func main() {
	// Current time
	ct := time.Now()
	fmt.Println(ct)

	// Calculate time 2 hour from now
	twoHours := 2 * time.Hour
	futureTime := time.Now().Add(twoHours)
	fmt.Println(futureTime)

	// Execution time
	start := time.Now()
	for i := 0; i <= 10000000; {
		i++
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
