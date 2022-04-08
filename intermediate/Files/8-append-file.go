package main

import (
	"log"
	"os"
)

func main() {

	fileName := "test2.txt"

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // parameter for appending file os.O_APPEND

	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString("cloud\n"); err != nil {

		log.Fatal(err)
	}
}
