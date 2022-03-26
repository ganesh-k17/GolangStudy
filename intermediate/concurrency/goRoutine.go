/*

A goroutine is a lightweight execution thread in the Go programming language and a function that
executes concurrently with the rest of the program.

Goroutines are incredibly cheap when compared to traditional threads as the overhead of creating
a goroutine is very low. Therefore, they are widely used in Go for concurrent programming.

To invoke a function as a goroutine, use the go keyword.

*/

package main

import (
	"fmt"
	"time"
)

func foo(s string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, ": ", i)
	}
}

func main() {
	go foo("cycle 1")
	go foo("cycle 2")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("end")
}
