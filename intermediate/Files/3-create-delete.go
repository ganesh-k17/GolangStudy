package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var myFile *os.File

	fileInfo, err := os.Stat("test1.txt") //  getting the fileInfo to check if exists

	if errors.Is(err, os.ErrNotExist) {
		myFile, err = os.Create("test1.txt") // to create the file
		myFile.Close()
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	fileInfo, err = os.Stat("test1.txt") // getting the fileInfo to check the file size

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The file size is %d bytes\n", fileInfo.Size())    // Size of the file
		fmt.Println("The last modified time is ", fileInfo.ModTime()) // last modifie time
	}

	err = os.Remove("test1.txt") // deleting file

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("end")
}
