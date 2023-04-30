package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func(t int) {
			fmt.Println(t)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("end")
}

/*
$ go run spawn-goroutine.go
9
11
11
11
11
8
11
11
11
11
end

Here the expectation is we would get it printed 1 to 10, may be in not correct order.  But
here we are getting as output above but it is not proper. It is because of closure.

To resolve this we have to pass the variable (i) to the function so that the closure issue get resolved.
*/
