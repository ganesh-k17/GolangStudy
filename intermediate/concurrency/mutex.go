/*

A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time. This is done to prevent race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex

* Lock
* Unlock

Any code present between a call to Lock and Unlock will be executed by only one Goroutine.

mutex.Lock()

x = x + 1 // this statement be executed
          // by only one Goroutine
          // at any point of time

mutex.Unlock()

If one Goroutine already has the lock and if a new Goroutine is trying to get the lock, then the
new Goroutine will be stopped until the mutex is unlocked

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("end")
}
