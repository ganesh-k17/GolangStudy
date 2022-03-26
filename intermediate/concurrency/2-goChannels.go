/*

What if two goroutines want to communicate with each other? For that, Go has channels.

Go, channels are a means through which different goroutines communicate.

Think of them as pipes through which you can connect with different concurrent goroutines.
The communication is bidirectional by default, meaning that you can send and receive values
from the same channel.

Moreover, by default, channels send and receive until the other side is ready. This allows
goroutines to synchronize without explicit locks or condition variables.

Syntax: Make a channel: make(chan [value-type]), [value-type] is the data type of the values to send and receive, e.g., int

Send and receive values: channel <- and <- channel, where <- is the channel operator.

Close a channel: close(channel), After closing, no value will be sent to the channel.

Both sending and receiving are blocking operations by default.

its kind of asyn/wait methodology

*/

package main

import (
	"fmt"
	"time"
)

// Prints a greeting message using values received in
// the channel
func greet(c chan string) {

	name := <-c // receiving value from channel

	fmt.Println("Hello", name)

}

func main() {

	// Making a channel of value type string
	c := make(chan string)

	// Starting a concurrent goroutine
	go greet(c)

	time.Sleep(100 * time.Millisecond) // greet method will wait in the reciving side for the main to send the value.

	// Sending values to the channel c
	c <- "World"

	// Closing channel
	close(c)

	time.Sleep(1500 * time.Millisecond) // adding a time to greet complete

	fmt.Println("end")
}
