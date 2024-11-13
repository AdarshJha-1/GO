package main

import (
	"errors"
	"fmt"
)

func firstFunc() error {
	return fmt.Errorf("original error: something went wrong in firstFunc")
}

func secondFunc() error {
	firstErr := firstFunc()
	if firstErr != nil {
		return fmt.Errorf("failed in first func: %w", firstErr)
		// secondErr := errors.New("failed in second func")
		// return errors.Join(firstErr, secondErr)
	}
	return nil
}

func main() {
	err := secondFunc()
	fmt.Println("original error: ", err)
	innerErr := errors.Unwrap(err)
	fmt.Println(innerErr)
}
