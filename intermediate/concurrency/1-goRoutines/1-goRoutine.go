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
	/*
		Here we are adding a delay before the main goroutine ends so that the two started goroutines
		have time to finish.
		The main goroutine must be running for any other goroutines to run. If the main goroutine terminates,
		then the program will exit and no other goroutine will run.
	*/
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("end")
	/*
		As we can see from the output, the printing of both goroutines are interleaved, which reflects that
		the goroutines are running concurrently by the Go runtime.
	*/
}
