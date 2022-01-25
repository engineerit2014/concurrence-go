package main

import (
	"fmt"
	"sync"
)

// build: go build --race main.go

var balance = 100

func Deposit(amount int, wg *sync.WaitGroup, mux *sync.RWMutex) {
	defer wg.Done() //decrements the WaitGroup counter by one

	// Lock writing
	mux.Lock()
	b := balance
	balance = b + amount
	// Unlock writing
	mux.Unlock()
}

func Balance(mux *sync.RWMutex) int {
	mux.RLock()
	b := balance
	mux.RUnlock()
	return b
}

func main() {
	// Solution: The problem of having only 1 deposit at a time to avoid account imbalances is solved, and it is allowed
	// to have N people consulting the balance at the same time.
	// For this purpose we use sync.RWMutex

	var wg sync.WaitGroup
	var mux sync.RWMutex

	for i := 0; i < 5; i++ {
		// increment the WaitGroup counter of go routines
		wg.Add(1)
		go Deposit(i*100, &wg, &mux)
	}

	wg.Add(1)
	go Deposit(100, &wg, &mux)

	wg.Add(1)
	go Deposit(100, &wg, &mux)

	// blocks until the WaitGroup counter is zero.
	wg.Wait()

	fmt.Println(Balance(&mux))
}
