package main

import (
	"io/ioutil"
	"log"
)

func main() {

	src := "test2.txt"
	dest := "test3.txt"

	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644) // will create new file if does not exists.

	if err != nil {
		log.Fatal(err)
	}
}
