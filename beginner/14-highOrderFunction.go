/*
Highorder functions means, passing functions as argument for another functions
*/

/* The below code snippets has been modified to high order functions in the uncommented code snippents below this */

// package main

// import "fmt"

// func sum(a int, b int) int {
// 	return a + b
// }

// func mul(a int, b int) int {
// 	return a * b
// }

// func main() {
// 	s := 2

// 	if s == 1 {
// 		println("sum is ", sum(1, 2))
// 	} else if s == 2 {
// 		println("multiplication is ", mul(1, 2))
// 	}

// 	fmt.Println("end")
// }

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}

func printResult(s int) func(int, int) int {
	optionMap := map[int]func(int, int) int{
		1: sum,
		2: mul,
	}
	return optionMap[s]
}

func main() {
	s := 2

	fmt.Println(printResult(s)(1, 2))

	fmt.Println("end")
}
