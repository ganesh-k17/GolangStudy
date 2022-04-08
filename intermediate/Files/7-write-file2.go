package main

import (
	"fmt"
	"os"
)

func main() {

	fileName := "test2.txt"
	var f *os.File

	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err.Error())
	}

	words := []string{"sky", "falcon", "rock", "hawk"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("end")
}
