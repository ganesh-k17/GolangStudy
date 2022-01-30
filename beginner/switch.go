package main

import "fmt"

func main() {

	var number = 21

	/* Switch with single conditions */

	switch number {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Others")
	}

	/* Switch with multiple conditions */

	switch number {
	case 0, 1, 2:
		fmt.Println("starting...")
	case 3, 4, 5:
		fmt.Println("In middle...")
	case 6, 7, 8:
		fmt.Println("in the ends...")
	default:
		fmt.Println("Others")
	}

}
