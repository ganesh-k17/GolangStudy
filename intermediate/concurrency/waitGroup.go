/*
To wait for multiple goroutines to finish, we can use a wait group.

When multiple go routines are running main would not wait for go routines to complete.
As soon as go routines start and when main complete its statements and terminated routines
also get terminated.

So to get the main wait, we should have a sleep statement till go routines completes.
In other way we can use waitgroup to count go routines completion.

How it works,

WaitGroup exports 3 methods.

1	Add(int) - It increases WaitGroup counter by given integer value.
2	Done()	 - It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
3	Wait()   - It Blocks the execution until it’s internal counter becomes 0.

*/

package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	fmt.Print("\nI am first runner")
	defer wg.Done() // This decreases counter by 1

}

func runner2(wg *sync.WaitGroup) {
	fmt.Print("\nI am second runner")
	defer wg.Done()
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2``
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	// Launching both the runners
	execute()
}
