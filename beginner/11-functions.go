/*

Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code
which ultimately saves the excessive use of memory, acts as a time saver and more importantly, provides better readability of
the code. So basically, a function is a collection of statements that perform some specific task and return the result to the caller.
A function can also perform some specific task without returning anything.

func function_name(Parameter-list)(Return_type){
    // function body.....
}

*/

package main

import (
	"fmt"
)

func main() {
	sum := add(5, 6) // Calling functions

	fmt.Printf("Sum is  %d", sum)

	fmt.Println("")

	// testing call by value

	var a = 5
	var b = 8
	swap_value(a, b) // calling call by value fucntion

	fmt.Printf("a is %d and b is %d", a, b)
	fmt.Println()

	swap_ref(&a, &b) // calling call by refernce function

	fmt.Printf("a is %d and b is %d", a, b)

	fmt.Println()
	fmt.Println("end")
}

func add(a int, b int) int { // Defining function
	return a + b
}

/*
Call by value : In this parameter passing method, values of actual parameters are copied to function’s formal parameters and
the two types of parameters are stored in different memory locations. So any changes made inside functions are not reflected in
actual parameters of the caller.
*/

func swap_value(a int, b int) {
	var t = 0
	t = a
	a = b
	b = t
}

/*
 Call by reference: Both the actual and formal parameters refer to the same locations, so any changes made inside the function
 are actually reflected in actual parameters of the caller.
*/

func swap_ref(a *int, b *int) {
	var t = 0
	t = *a
	*a = *b
	*b = t
}
