/*

Read condition occur when an variable or object is accessed (read or write) concurrently.
The value might be inconsistence since the while reading the value, the variable or object
might be altered by other routine or thread.

To avoid race condition we can use mutex lock to restrain the object or variable from
accessing by other go routine till it is unlocked.  It has been discussed in mutex.go file.

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0

	const num = 100
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			temp := counter
			temp++
			runtime.Gosched()
			counter = temp
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}

/*

Note:

runtime.GoShed(): in go normally once a thread take a control it would continue to run untill it finish.  runtime.GoShed allow it to perform other go routines.

*/

/*

Output:
PS E:\Golang\GolangStudy\intermediate\concurrency> go run raceCondition.go
count: 71
PS E:\Golang\GolangStudy\intermediate\concurrency> go run raceCondition.go
count: 59
PS E:\Golang\GolangStudy\intermediate\concurrency> go run raceCondition.go
count: 50

*/
