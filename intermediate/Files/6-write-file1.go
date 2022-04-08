package main

import (
	"fmt"
	"io/ioutil"
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

	b := []byte("hello")

	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("end")
}
