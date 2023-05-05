/*
Output should be

sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs


The main idea of this question is to test users ability to start and close jobs correctly.


As long as the job is closed successfully, the validation should pass for any given number of sent and received jobs.

*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				wg.Done()
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("sent job", j)
		jobs <- j
	}

	// your code goes here

	fmt.Println("sent all jobs")

	wg.Wait()

	close(jobs)

	if <-done {
		close(done)
	}

}
