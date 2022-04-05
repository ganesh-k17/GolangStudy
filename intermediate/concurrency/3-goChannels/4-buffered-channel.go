/*

https://www.geeksforgeeks.org/buffered-channel-in-golang

Channels can be defined as a pipes using which Goroutines communicate. Similar to water flows from one end
to another in a pipe, data can be sent from one end and received from the another end using channels.

By default channels are unbuffered, which states that they will only accept sends (chan <-) if there is
a corresponding receive (<- chan) which are ready to receive the sent value.

Buffered channels allows to accept a limited number of values without a corresponding receiver for
 those values. It is possible to create a channel with a buffe. Buffered channel are blocked only when the
 buffer is full. Similarly receiving from a buffered channel are blocked only when the buffer will be empty.

Buffered channels can be created by passing an additional capacity parameter to the make( ) function which specifies the size of the buffer.

Syntax :

ch := make(chan type, capacity)           // chan defines channel type

Here , capacity in the above syntax should be greater than 0 for a channel to have a buffer.
The capacity for an unbuffered channel is 0 by default and hence it omit the capacity parameter.

*/
package main

import (
	"fmt"
)

func main() {

	// create a buffered channel
	// with a capacity of 2.
	ch := make(chan string, 2)
	ch <- "geeksforgeeks"
	ch <- "geeksforgeeks world"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("done")
}

/*

If we try to send another text to the channel then it will create deadlock as below.
But here we have to receive a value before sending another data to the channel to avoid deadlock

package main

import (
	"fmt"
)

func main() {

	// create a buffered channel
	// with a capacity of 2.
	ch := make(chan string, 2)
	ch <- "geeksforgeeks"
	ch <- "geeksforgeeks world"
	ch <- "hello"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("done")
}

Output

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        E:/Golang/GolangStudy/intermediate/concurrency/3-goChannels/4-buffered-channel.go:38 +0xa5
exit status 2
*/
