package services

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func ProcessOrder(orderID int, orderDetails string) {
	mu.Lock()
	fmt.Printf("Processing order #%d: %s\n", orderID, orderDetails)
	go func() {
		defer mu.Unlock()
		time.Sleep(2 * time.Second) // Simulate processing delay
		fmt.Printf("Order #%d processed successfully!\n", orderID)
	}()
}
