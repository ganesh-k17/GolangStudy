/*

In the below program we have been closing file in each error if it got failed to do the job.
And we are closing the file after completing all the operations.  So we are repeating the code again and again.


package main

import (
	"fmt"
	"io"
	"os"
)

var path = "test.txt"

var myFile *os.File

func Write() bool {

	len, err := myFile.WriteString("Welcome to GeeksforGeeks.")

	println("length of the file content %s", len)

	if err != nil {
		println("Error while writing file")
		println(err.Error())
		myFile.Close()
	}
	return true
}

func main() {
	myFile, err := os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = myFile.WriteString("Welcome to GeeksforGeeks.")

	if err != nil {
		fmt.Println(err.Error())
		myFile.Close()
	}

	myFile.Seek(0, io.SeekStart) // moving to start of file

	var b []byte = make([]byte, 1024)
	len, err := myFile.Read(b)

	if err != nil {
		fmt.Println(err.Error())
		myFile.Close()
	} else {
		fmt.Printf("Read %d bytes: %s", len, b)
	}

	// Write()

	myFile.Close()

	fmt.Println("done")
}


To avoid the duplication we can use defer, So that the defer will call when the function is going to exit in either way due to error or completion.
Below program we have used defer and it will call at the end.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

var path = "test.txt"

var myFile *os.File

func close() {
	myFile.Close()
}

func main() {
	myFile, err := os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = myFile.WriteString("Welcome to GeeksforGeeks.")

	if err != nil {
		fmt.Println(err.Error())
	}

	myFile.Seek(0, io.SeekStart) // moving to start of file

	var b []byte = make([]byte, 1024)
	len, err := myFile.Read(b)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Read %d bytes: %s", len, b)
	}

	defer close()

	fmt.Println("done")
}
