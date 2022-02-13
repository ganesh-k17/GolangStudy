package main

import "fmt"

func main() {
	switch number := 5; {
	case number < 2:
		fmt.Println("< 2")
	case number == 2:
		fmt.Println("== 2")
	case number > 2:
		fmt.Println("> 2")
	default:
		fmt.Println("other")
	}
}
