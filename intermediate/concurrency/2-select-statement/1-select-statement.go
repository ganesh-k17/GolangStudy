/*
In Go language, the select statement is just like switch statement, but in the select statement,
case statement refers to communication, i.e. sent or receive operation on the channel.

Syntax:

select{

case SendOrReceive1: // Statement
case SendOrReceive2: // Statement
case SendOrReceive3: // Statement
.......
default: // Statement

*/

// Select statement waits until the communication(send or receive operation) is prepared for
// some cases to begin. for example,

package main

import (
	"fmt"
	"time"
)

func portal1(channel1 chan string) {
	time.Sleep(3 * time.Second)
	channel1 <- "welcome to channel1"
}

func portal2(channel2 chan string) {
	time.Sleep(3 * time.Second)
	channel2 <- "welcome to channel2"
}

func main() {
	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling func 1 and func 2 as go routines.
	go portal1(R1)
	go portal2(R2)

	select {

	// case 1 for portal 1
	case opt1 := <-R1:
		fmt.Println(opt1)
	// case 1 for portal 2
	case opt2 := <-R2:
		fmt.Println(opt2)

	}

	fmt.Println("end")
}

/*
Explanation: In the above program, portal 1 sleep for 3 seconds and portal 2 sleep for 9 seconds after
their sleep time over they will ready to proceed. Now, select statement waits till their sleep time,
when the portal 2 wakes up, it selects case 2 and prints “Welcome to channel 1”. If the portal 1 wakes up
before portal 2 then the output is “welcome to channel 2”.
*/

/*
If a select statement does not contain any case statement, then that select statement waits forever.
Syntax:

select{}
Example:

// Go program to illustrate the
// concept of select statement

package main

func main() {

    // Select statement
   // without any case
   select{ }

}

Output:

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
main.main()
    /home/runner/main.go:9 +0x20
exit status 2
*/
