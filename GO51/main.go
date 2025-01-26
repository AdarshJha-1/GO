package main

import (
	"fmt"
	"sync"
)

var (
	mutex1 sync.Mutex
	mutex2 sync.Mutex
)

func resource1(name string) {
	mutex1.Lock()
	fmt.Println(name, "acquired resource 1")
	resource2(name)
	mutex1.Unlock()
}

func resource2(name string) {
	mutex2.Lock()
	fmt.Println(name, "acquired resource 2")
	resource1(name)
	mutex2.Unlock()
}

func main() {
	// var counter int
	// var lock sync.Mutex

	// for range 1000 {
	// 	go func() {
	// 		lock.Lock()
	// 		counter++
	// 		lock.Unlock()
	// 	}()
	// }

	// // time.Sleep(10 * time.Second)

	// lock.Lock()
	// fmt.Println("Expected counter:", 1000)
	// fmt.Println("Actual counter:", counter)
	// lock.Unlock()

	go resource1("Goroutine 1")
	go resource1("Goroutine 2")

	fmt.Println("waiting...")
	select {}
}
