/*
https://www.geeksforgeeks.org/multiple-goroutines

A Goroutine is a function or method which executes independently and simultaneously in connection with any other
Goroutines present in your program. Or in other words, every concurrently executing activity in Go language is
known as a Goroutines.

In Go language, you are allowed to create multiple goroutines in a single program. You can
create a goroutine simply by using go keyword as a prefixing to the function or method call as shown in the below syntax:

func name(){

// statements
}

// using go keyword as the
// prefix of your function call
go name()

*/

package main

import (
	"fmt"
	"time"
)

// For goroutine 1
func Aname() {

	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

// For goroutine 2
func Aid() {

	arr2 := [4]int{300, 301, 302, 303}

	for t2 := 0; t2 <= 3; t2++ {

		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

// Main function
func main() {

	fmt.Println("!...Main Go-routine Start...!")

	// calling Goroutine 1
	go Aname()

	// calling Goroutine 2
	go Aid()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}

/*
Output:

!...Main Go-routine Start...!
Rohit
Suman
Aman
300
Ria
301
302
303

!...Main Go-routine End...!

Creation: In the above example, we have two goroutines other than main goroutine, i.e, Aname, and Aid. Here,
Aname prints the name of the authors and Aid prints the id of the authors.

Working: Here, we have two goroutines, i.e, Aname, and Aid and one main goroutine. When we run this program
first the main goroutine strat and print “!...Main Go-routine Start...!“, here the main goroutine is like a
parent and other goroutines are its children so first main goroutine run after that those other goroutines start and
if the main goroutine terminates, then other goroutines are also terminated. So, after the main goroutine,
Aname and Aid goroutines start their working concurrently. The Aname goroutine starts it working from 150ms and Aid
start its working from 500ms

*/
