package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", c.Code, c.Message)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// return 0, errors.New("cannot divide by zero")
		// return 0, fmt.Errorf("cannot divide by zero %.2f", b)
		return 0, CustomError{Code: 400, Message: "cannot divide by zero"}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
