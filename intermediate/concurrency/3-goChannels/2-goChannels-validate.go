/*
The receiver goroutine can check the channel is open or close with the help of the given syntax:

ele, ok:= <- Mychannel
Here, if the value of ok is true which means the channel is open so, read operations can be performed. And if the value of is false which means the channel is closed so, read operations are not going to perform.

Example:
*/

// Go program to illustrate how
// to close a channel using for
// range loop and close function

package main

import "fmt"

// Function
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

// Main function
func main() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}

/*

Here the channel got closed before the receiving end.  In the validation we had validated whether the channel is closed.const

Output:

PS E:\Golang\GolangStudy\intermediate\concurrency\3-goChannels> go run 2-goChannels-validate.go
Channel Open  GeeksforGeeks true
Channel Open  GeeksforGeeks true
Channel Open  GeeksforGeeks true
Channel Open  GeeksforGeeks true
Channel Close  false

*/
