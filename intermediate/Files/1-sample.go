package main

import (
	"fmt"
	"io"
	"os"
)

var myFile *os.File

func main() {

	defer myFile.Close()

	myFile, err := os.OpenFile("test.txt", os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = myFile.WriteString("Hello go!!!")

	myFile.Seek(0, io.SeekStart) // moving to start of file

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var b []byte = make([]byte, 1024)
	len, err := myFile.Read(b)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Read %d bytes: %s", len, b)
}
