/*
The blocking of select statement means when there is no case statement is ready and the select statement
does not contain any default statement, then the select statement block until at least one case statement or
communication can proceed.

Example:
*/

package main

func main() {

	// creating channel
	mychannel := make(chan int)

	// channel is not ready
	// and no default case
	select {
	case <-mychannel:

	}
}

/*
Output:

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
main.main()
        E:/Golang/GolangStudy/intermediate/concurrency/2-select-statement/2-select-without-case.go:19 +0x27
exit status 2
PS E:\Golang\GolangStudy\intermediate\concurrency\2-select-statement> go run 3-select-nocase-ready.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        E:/Golang/GolangStudy/intermediate/concurrency/2-select-statement/3-select-nocase-ready.go:19 +0x54
exit status 2
PS E:\Golang\GolangStudy\intermediate\concurrency\2-select-statement>

*/
