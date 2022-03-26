/*

https://www.geeksforgeeks.org/unidirectional-channel-in-golang

As we know that a channel is a medium of communication between concurrently running goroutines so that they can send and receive data to each other. By default a channel is bidirectional but you can create a unidirectional channel also. A channel that can only receive data or a channel that can only send data is the unidirectional channel. The unidirectional channel can also create with the help of make() function as shown below:

// Only to receive data
c1:= make(<- chan bool)

// Only to send data
c2:= make(chan<-bool

*/

package main

import "fmt"

// Main function
func main() {

	// Only for receiving
	mychanl1 := make(<-chan string)

	// Only for sending
	mychanl2 := make(chan<- string)

	// Display the types of channels
	fmt.Printf("%T", mychanl1)
	fmt.Printf("\n%T", mychanl2)
}

/*

Output:

<-chan string
chan<- string

*/
