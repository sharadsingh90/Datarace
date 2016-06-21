package main

import "sync"

//// setup
const (
	N    = 1000
	Conc = 1000
)

var Lock sync.Mutex

func adder(acc *int, done chan bool) {
	Lock.Lock()
	for i := 0; i < N; i++ {
		*acc++
	}
	Lock.Unlock()
	done <- true
}

func main() {
	acc := 0
	done := make(chan bool)
	for i := 0; i < Conc; i++ {
		go adder(&acc, done)
	}

	//	// Wait for all goroutines to complete
	for i := 0; i < Conc; i++ {
		<-done
	}

	println(acc) // Should == Conc x N
}
