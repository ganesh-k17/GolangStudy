package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
	fmt.Println("end")
}

func leak(wg *sync.WaitGroup) {
	ch := make(chan string)
	go func() {
		value := <-ch
		fmt.Println("received value", value)
		wg.Done()
	}()

	fmt.Println("Exiting leak method")
	wg.Done()
}
