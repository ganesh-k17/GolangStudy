package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	_, err := os.Stat("test.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist")
	} else {
		fmt.Println("file exists")
	}

	fmt.Println("end")
}
