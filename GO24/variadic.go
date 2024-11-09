package main

import "fmt"

func main() {
	// numsSum := sum(1, 2, 3, 4, 5, 6, 8, 9, 10)
	// fmt.Println(numsSum)

	nums := []int{123, 434, 54, 52, 23, 2}
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
