package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// Create first goroutine here
	go printNumbers(&wg)
	// Create second goroutine here
	go printLetters(&wg)
	// Wait for both goroutines to finish here
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
	values := [5]string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}

	wg.Done()
}
