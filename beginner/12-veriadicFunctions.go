/*
https://gobyexample.com/variadic-functions

The function that called with the varying number of arguments is known as variadic function. Or in other words, a user is allowed to pass zero or more arguments in the variadic function. fmt.Printf is the example of the variadic function, it required one fixed argument at the starting after that it can accept any number of arguments.
Important Points:


In the declaration of the variadic function, the type of the last parameter is preceded by an ellipsis, i.e, (…). It indicates that the function can be called at any number of parameters of this type.
Syntax:

function function_name(para1, para2...type)type{
// code...
}
Inside the function …type behaves like a slice. For example, suppose we have a function signature, i.e, add( b…int)int, now the a parameter of type[]int.
You can also pass an existing slice in a variadic function. To do this, we pass a slice of the complete array to the function as shown in Example 2 below.
When you do not pass any argument in the variadic function, then the slice inside the function is nil.
The variadic functions are generally used for string formatting.
You can also pass multiple slice in the variadic function.
You can not use variadic parameter as a return value, but you can return it as a slice.

When we use a Variadic function:


Variadic function is used when you want to pass a slice in a function.
Variadic function is used when we don’t know the quantity of parameters.
When you use variadic function in your program, it increase the readability of your program.

*/

package main

import "fmt"

func main() {
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(values ...int) {
	fmt.Print(values, "")
	total := 0
	for _, num := range values {
		total += num
	}
	fmt.Println(total)
}
