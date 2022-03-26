/*
If a select statement does not contain any case statement, then that select statement waits forever.
Syntax:

select{}
Example:

*/

// Go program to illustrate the
// concept of select statement

package main

func main() {

	// Select statement
	// without any case
	select {}

}

/*
output -

PS E:\Golang\GolangStudy\intermediate\concurrency\2-select-statement> go run 2-select-without-case.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
main.main()
        E:/Golang/GolangStudy/intermediate/concurrency/2-select-statement/2-select-without-case.go:19 +0x27
exit status 2
PS E:\Golang\GolangStudy\intermediate\concurrency\2-select-statement>

*/
