/*
In this question, we are working on a project that involves processing a large amount of data in
parallel using goroutines. We want to use channels to communicate between the goroutines and
ensure that all goroutines have finished processing before the program exits.


Complete the code, so that the program runs successfully without any errors.

Expected output:

sent job 1
received job 1
received job 2
sent job 2
received result 2
received result 4
received job 3
received result 6
sent job 3

Where to close the channels and where to wg.Done() is something important to proper working of the program

*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for j := range jobs {
			fmt.Println("received job", j)
			result := j * 2
			results <- result
		}
		close(results)
		wg.Done()
		// your code goes here
	}()

	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		close(jobs)
		wg.Done()
		// your code goes here
	}()

	go func() {
		for r := range results {
			fmt.Println("received result", r)
		}
		wg.Done()
		// your code goes here
	}()

	wg.Wait()
}
