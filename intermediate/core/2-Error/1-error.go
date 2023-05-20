package main

import (
	"errors"
	"fmt"
)

func process(i int) error {
	if i%2 == 0 {
		return errors.New("only odd numbers are allowed.")
	}
	return nil
}

func main() {
	err := errors.New("custom error occured") // creating a custom error
	fmt.Println(err)
	err = process(6)
	if err != nil {
		fmt.Println(err)
	}
}
