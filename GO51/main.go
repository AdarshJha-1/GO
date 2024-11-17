package main

import (
	"fmt"
	"sync"
)

var (
	mutex1 sync.Mutex
	mutex2 sync.Mutex
)

// Deadlock
func resource1(name string) {
	mutex1.Lock()
	fmt.Println(name, "acquired resource 1")
	resource2(name)
	mutex1.Unlock()
}
func resource2(name string) {
	mutex1.Lock()
	fmt.Println(name, "acquired resource 2")
	resource1(name)
	mutex1.Unlock()
}

func main() {
	/*
		var counter int

		var lock sync.Mutex
		for range 10 {
			go func() {
				lock.Lock()
				counter++
				lock.Unlock()
			}()
		}

		lock.Lock()
		fmt.Println("Expected counter:", 10)
		lock.Unlock()
		fmt.Println("Actual counter:", counter)
	*/

	go resource1("Goroutine 1")
	go resource2("Goroutine 2")

	fmt.Println("Waiting...")
	select {}
}
