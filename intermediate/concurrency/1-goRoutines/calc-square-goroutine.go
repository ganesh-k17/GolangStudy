package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	// all the 50 calculateSquare run in different 50 goroutines so it
	// would be completed early than a sequential program
	for i := 0; i < 50; i++ {
		go calculateSquare(i)
	}
	// Here the calculation of time is not sense as it would get start run once
	// main got executed as all in the loops are running in goroutines.
	// To make it reliable we have to use wait group
	end := time.Since(start)
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("elapsed ", end)
	fmt.Println("end")
}
