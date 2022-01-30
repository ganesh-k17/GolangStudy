package main

import "fmt"

func main() {

	country := "India"

	/* if condition */
	if country == "India" {
		fmt.Print("India")
	}

	/* if else conditions */
	if country == "India" {
		fmt.Print("India")
	} else {
		fmt.Print("Other")
	}

	country = "USA"
	/* if else if */
	if country == "India" {
		fmt.Print("India")
	} else if country == "USA" {
		fmt.Print("USA")
	} else {
		fmt.Print("Other")
	}

}
