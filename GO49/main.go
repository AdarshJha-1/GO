package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, task stopping")
			return
		default:
			fmt.Println("Simulating work..")
			time.Sleep(1 * time.Second)
		}
	}
}

func ProcessOrder(ctx context.Context, orderID int) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("Error: userID not found in context")
		return
	}
	fmt.Printf("Processing order %d for user %d\n", orderID, userID)
}

func GetUserAndProcessOrder(orderID int) {
	ctx := context.WithValue(context.Background(), "userID", 69)
	ProcessOrder(ctx, orderID)
}

func main() {
	/*
		// Context Go

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // cancellation happens when main exits

		go longRunningTask(ctx)

		// simulate some main thread work for 2 seconds
		time.Sleep(2 * time.Second)

		// cancel the context after some time
		fmt.Println("Cancelling context")
		cancel()

		// wait for the goroutine to finish
		time.Sleep(1 * time.Second)
	*/

	GetUserAndProcessOrder(6969)
}
