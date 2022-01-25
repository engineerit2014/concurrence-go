package main

import (
	"fmt"
	"sync"
)

// build: go build --race main.go

var balance = 100

func Deposit(amount int, wg *sync.WaitGroup, mux *sync.Mutex) {
	defer wg.Done() //decrements the WaitGroup counter by one

	// Lock writing
	mux.Lock()
	b := balance
	balance = b + amount
	// Unlock writing
	mux.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {
	// Problem Race condition: There is no race condition but now I can't get the balance whenever I want because it may be blocked.
	// For this purpose we use sync.Mutex to lock & unlock when it is writing
	// Problem Solution: see the example3.go

	var wg sync.WaitGroup
	var mux sync.Mutex

	for i := 0; i < 5; i++ {
		// increment the WaitGroup counter of go routines
		wg.Add(1)
		go Deposit(100, &wg, &mux)
	}

	// blocks until the WaitGroup counter is zero.
	wg.Wait()

	fmt.Println(Balance())
}
