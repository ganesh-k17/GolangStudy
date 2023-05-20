/*
go doc errorf
doc: no symbol errorf in package E:\Golang\GolangStudy\intermediate\core\Error
exit status 1

Ganesh K@DESKTOP-UJ4106D MINGW64 /e/Golang/GolangStudy/intermediate/core/Error (main)
$ go doc fmt errorf
package fmt // import "fmt"

func Errorf(format string, a ...interface{}) error
    Errorf formats according to a format specifier and returns the string as a
    value that satisfies error.

    If the format specifier includes a %w verb with an error operand, the
    returned error will implement an Unwrap method returning the operand. It is
    invalid to include more than one %w verb or to supply it with an operand
    that does not implement the error interface. The %w verb is otherwise a
    synonym for %v.
*/

package main

import (
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("only odd numbers are allowed., got %d", i) // it also return error
	}
	return nil
}

func main() {
	err := process(6)
	if err != nil {
		fmt.Println(err)
	}
}
