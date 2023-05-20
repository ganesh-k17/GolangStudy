package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir", "dir2", "dir3", "text.txt") // we dont need to use contenate with character '/' to form a file path.

	//We can just use filepath to get it generated.
	fmt.Println(path)

	fmt.Println(filepath.Dir(path)) // gives the directory path of the file

	fmt.Println(filepath.Base(path))

	fmt.Println(filepath.IsAbs(path))

	fmt.Println(filepath.IsAbs("d:/dir/file"))

	fmt.Println(filepath.Ext(path))
}

/*
output:

dir\dir2\dir3\text.txt
dir\dir2\dir3
*/
