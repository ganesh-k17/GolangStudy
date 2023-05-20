package main

import (
	"fmt"
	"os"
)

func main() {
	path := "1-filepath.go"

	simpleRead(path) // using string(data)

	readAsChunks(path) // using opening the file and read as chunks.  it is helpful to read from large file

}

/*
here we are just convert the data to string and print it.
It is not always feasible if the file is big and not enough to hold the value
we could not use it.  We have to read the bytes as chunks and print it.
*/
func simpleRead(path string) {

	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reading File")
		fmt.Println(string(data))
	}
}

func readAsChunks(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		b := make([]byte, 4)
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
