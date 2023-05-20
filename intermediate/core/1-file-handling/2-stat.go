/*
$ go doc os Stat
package os // import "os"

func Stat(name string) (FileInfo, error)
    Stat returns a FileInfo describing the named file. If there is an error, it
    will be of type *PathError.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("filepath.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fileInfo.Name())
		fmt.Println(fileInfo.Size())
		fmt.Println(fileInfo.Mode())
		fmt.Println(fileInfo.IsDir())
	}
}
