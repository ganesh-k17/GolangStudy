/*

* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

Declaring a pointer:

var pointer_name *Data_Type

eg: var s *string

// normal variable declaration
var a = 45

// Initialization of pointer s with
// memory address of variable a
var s *int = &a

*/

package main

import "fmt"

func main() {
	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x

	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Value of p = ", *p)

	fmt.Println("end")
}
