package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Create wait group
	var wg sync.WaitGroup

	// Add five goroutine to wait group
	x := 5
	wg.Add(x)

	// Print the current date and time 5 times
	for i := 0; i < x; i++ {
		go fmt.Println(CurrentDateTime(&wg))
	}

	// Wait for all goroutines to finish
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

// CurrentDateTime returns the current date and time with wait group
func CurrentDateTime(wg *sync.WaitGroup) string {
	// Defer the wait group
	defer wg.Done()
	// Sleep a random milisecond to simulate a real time.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// Return the current date and time with goroutine index
	return time.Now().Format(time.RFC3339)
}

// Sum two numbers
func Sum(a, b int) int {
	return a + b
}
