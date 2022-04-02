/*
The default statement in the select statement is used to protect select statement from blocking.
This statement executes when there is no case statement is ready to proceed.

Example:
*/

package main

import "fmt"

func main() {

	// creating channel
	mychannel := make(chan int)

	select {

	case <-mychannel:

	default:
		fmt.Println("Not found")

	}
}

/* Output:

Not found

*/
