/*
Go has a special statement called defer that schedules a function call to be run after
the function completes.

in the below example, The program prints First followed by Second after completing all the functionalities.
*/

package main

import (
	"fmt"
)

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func main() {
	defer second()
	first()

	fmt.Println("done")
}
