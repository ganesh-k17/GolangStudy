package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done() // each done the wg counter decrease a value
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(10) // adding counter for 10 go routines as we are going to run the loop for 10
	for i := 0; i < 10; i++ {
		go calculateSquare(i, &wg)
	}
	wg.Wait() // ask main to wait to all the wait groups are done means all go routines are done.

	end := time.Since(start)

	fmt.Println("elapsed ", end) // here the time calculation is reliable as we are not using any sleep
	fmt.Println("end")
}
