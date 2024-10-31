package main

import (
	"fmt"
	"strconv"
)

func main() {

	// right shift by i is equal to number divided by 2^i
	var a uint8 = 20
	// b := a >> 2
	// fmt.Println(b)

	// // lest shift by i is equal to number multiply by 2^i
	// var x uint8 = 10
	// y := x << 1
	// fmt.Println(y)

	// // 1 << i will result in the a single bit using set in the ith position while rest of the bits are zero
	// var num uint8 = 1
	// fmt.Println(strconv.FormatUint(uint64(num), 2))
	// num = num << 3
	// fmt.Println(strconv.FormatUint(uint64(num), 2))

	// bit mask
	a = 1

	// a = a << 3 // bit-mask where 3rd bit is set to 1
	a <<= 3 // bit-mask where 3rd bit is set to 1
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a = 5
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	a |= (1 << 1)
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	a &^= (1 << 1)
	fmt.Println(strconv.FormatUint(uint64(a), 2))

}
