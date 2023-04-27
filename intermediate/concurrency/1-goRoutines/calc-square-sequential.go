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
	for i := 0; i < 50; i++ {
		calculateSquare(i)
	}
	end := time.Since(start)
	fmt.Println("end in ", end)
}
