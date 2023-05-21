package main

import (
	"fmt"
	"os"
)

func main() {
	// your code goes here
	file, err := os.Open("temp.txt")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		b := make([]byte, 50)
		for {
			n, err := file.Read(b)
			if err != nil {
				fmt.Println(err.Error())
				break
			} else {
				fmt.Println(b[0:n])         // it will print bytes
				fmt.Println(string(b[0:n])) // it will print the data
			}
		}
	}
}
