package main

import "fmt"

func main() {
	fact := 1 // var fact = 1
	number := 5
	for i := 1; i <= number; i++ {
		fact = fact * i
	}
	fmt.Print(fact)
}
