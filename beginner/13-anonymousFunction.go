/*
* An anonymous function is a function that was declared without any named identifier to refer to it.
* Anonymous functions can accept inputs and return outputs, just as standard functions do.
* Anonymous functions can be passed as parameter as well as return types.
 */

package main

import "fmt"

/*
func main() {
	fmt.Println("hello")
	func() {
		fmt.Println("inside method")
	}()
}
*/

/*
func main() {
	fmt.Println("hello")

	value := func() {
		fmt.Println("inside method")
	}

	value()
}
*/

var area = func(l int, b int) int {
	return l * b
}

func main() {
	fmt.Println(area(5, 6))
	fmt.Println("end")
}
