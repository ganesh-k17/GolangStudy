/*

What if two goroutines want to communicate with each other? For that, Go has channels.

Go, channels are a means through which different goroutines communicate.

Think of them as pipes through which you can connect with different concurrent goroutines.
The communication is bidirectional by default, meaning that you can send and receive values
from the same channel.

Moreover, by default, channels send and receive until the other side is ready. This allows
goroutines to synchronize without explicit locks or condition variables.

Syntax:

var Channel_name chan Type

or

Make a channel: make(chan [value-type]), [value-type] is the data type of the values to send and receive, e.g., int

Send and receive values: channel <- and <- channel, where <- is the channel operator.

Close a channel: close(channel), After closing, no value will be sent to the channel.

Both sending and receiving are blocking operations by default.

its kind of asyn/wait methodology

* Send operation:

The send operation is used to send data from one goroutine to another goroutine with the help of a channel. Values like
int, float64, and bool can safe and easy to send through a channel because they are copied so there is no risk of
accidental concurrent access of the same value. Similarly, strings are also safe to transfer because they are immutable.
But for sending pointers or reference like a slice, map, etc. through a channel are not safe because the value of pointers or
reference may change by sending goroutine or by the receiving goroutine at the same time and the result is unpredicted.
So, when you use pointers or references in the channel you must make sure that they can only access by the one goroutine at a time.

Mychannel <- element

The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- operator.

* Receive operation:

The receive operation is used to receive the data sent by the send operator.

element := <-Mychannel

The above statement indicates that the element receives data from the channel(Mychannel).

If the result of the received statement is not going to use is also a valid statement.
You can also write a receive statement as:

<-Mychannel

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

/*

Closing a Channel:

You can also close a channel with the help of close() function. This is an in-built function and sets a flag which
indicates that no more value will send to this channel.

Syntax:

close()

You can also close the channel using for range loop. Here, the receiver goroutine can check the channel is open or
close with the help of the given syntax:

ele, ok:= <- Mychannel

Here, if the value of ok is true which means the channel is open so, read operations can be performed.
And if the value of is false which means the channel is closed so, read operations are not going to perform.

*/
