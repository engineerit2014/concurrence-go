package main

import (
	"fmt"
	"sync"
)

// build: go build --race main.go

var balance = 100

func Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done() //decrements the WaitGroup counter by one
	b := balance
	balance = b + amount
}

func Balance() int {
	b := balance
	return b
}

func main() {
	// Let's use go routines to simulate N deposits at same time
	// Problem Race condition: there is a race condition with the balance variable because many are trying to deposit at the same time which will generate information loss.
	// Solution: example2.go

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		// increment the WaitGroup counter of go routines
		wg.Add(1)
		go Deposit(100, &wg)
	}

	// blocks until the WaitGroup counter is zero.
	wg.Wait()

	fmt.Println(Balance())
}
